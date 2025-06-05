package domain

type OrderStatus string

const (
	Building  OrderStatus = "building"
	Cancelled OrderStatus = "cancelled"
	Received  OrderStatus = "received"
	Preparing OrderStatus = "confirmed"
	Ready     OrderStatus = "ready"
	Done      OrderStatus = "done"
)

type ItemOrderElement struct {
	ItemID   string
	Quantity int
}

type Order struct {
	id        string
	user      User
	items     []Item
	ItemOrder []ItemOrderElement
	status    OrderStatus
	price     int64
}

func NewOrder(user User) (*Order, *DomainError) {
	o := &Order{
		id:     "",
		user:   user,
		items:  []Item{},
		status: Building,
		price:  0,
	}

	return o, nil
}

func NewOrderWithoutUser() (*Order, *DomainError) {
	o := &Order{
		id:     "",
		user:   User{},
		items:  []Item{},
		status: Building,
		price:  0,
	}

	return o, nil
}

func (o *Order) GetID() string {
	return o.id
}

func (o *Order) GetUser() *User {
	return &o.user
}

func (o *Order) GetItems() *[]Item {
	return &o.items
}

func (o *Order) AddItemOrder(itemOrder []ItemOrderElement) {
	o.ItemOrder = itemOrder
}

func (o *Order) GetPrice() int64 {
	return o.price
}

func (o *Order) GetStatus() OrderStatus {
	return o.status
}

func (o *Order) SetStatus(status OrderStatus) *DomainError {
	if status != Received && status != Preparing && status != Ready && status != Done && status != Cancelled && status != Building {
		return NewDomainError(InvalidOrderStatus, "Invalid order status")
	}

	if o.status == Done {
		return NewDomainError(InvalidOrderStatus, "Cannot change status of a completed order")
	}

	if o.status == Cancelled {
		return NewDomainError(InvalidOrderStatus, "Cannot change status of a cancelled order")
	}

	if status == Building {
		return NewDomainError(InvalidOrderStatus, "Order cannot be set to building directly")
	}

	if status == Cancelled && o.status != Building {
		return NewDomainError(InvalidOrderStatus, "Order must be in building status to be cancelled")
	}

	if status == Received && o.status != Building {
		return NewDomainError(InvalidOrderStatus, "Order must be in building status to be received")
	}

	if status == Received && o.price == 0 {
		return NewDomainError(InvalidOrderStatus, "Order must have items before it can be received")
	}

	if status == Preparing && o.status != Received {
		return NewDomainError(InvalidOrderStatus, "Order must be received before it can be prepared")
	}

	if status == Ready && o.status != Preparing {
		return NewDomainError(InvalidOrderStatus, "Order must be prepared before it can be ready")
	}

	if status == Done && o.status != Ready {
		return NewDomainError(InvalidOrderStatus, "Order must be ready before it can be done")
	}

	o.status = status

	return nil
}

func (o *Order) AddItem(item Item) *DomainError {
	if o.status != Building {
		return NewDomainError(InvalidOrderStatus, "Cannot add items to an order that is not in building status")
	}

	o.items = append(o.items, item)
	o.price += item.GetPrice()

	return nil
}

func (o *Order) RemoveItem(itemID string) *DomainError {
	for i, item := range o.items {
		if item.GetID() == itemID {
			o.items = append(o.items[:i], o.items[i+1:]...)
			o.price -= item.GetPrice()
			return nil
		}
	}
	return NewDomainError(ProductNotFoundInOrder, "Item not found in order")
}

func (o *Order) SetId(id string) {
	o.id = id
}

func (o *Order) SetUser(user User) {
	o.user = user
}
