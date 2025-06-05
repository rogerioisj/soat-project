package domain

type OrderStatus string

const (
	Building  OrderStatus = "building"
	Cancelled OrderStatus = "cancelled"
	Received  OrderStatus = "received"
	Preparing OrderStatus = "preparing"
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

func (o *Order) SetStatus(status OrderStatus) {
	o.status = status
}

func (o *Order) UpgradeStage() *DomainError {
	switch o.status {
	case Building:
		o.SetStatus(Received)
		return nil
	case Received:
		o.SetStatus(Preparing)
		return nil
	case Preparing:
		o.SetStatus(Ready)
		return nil
	case Ready:
		o.SetStatus(Done)
		return nil
	default:
		return NewDomainError(InvalidOrderStatus, "Cannot upgrade stage from current status")
	}
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
