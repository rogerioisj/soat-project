package domain

type ProductType string

const (
	Snack         ProductType = "snack"
	Drink         ProductType = "drink"
	Dessert       ProductType = "dessert"
	Accompaniment ProductType = "accompaniment"
)

type Item struct {
	id          string
	name        string
	description string
	price       int64
	productType ProductType
	quantity    int64
}

func NewItem(id, name, description string, price int64, productType ProductType) (*Item, *DomainError) {
	i := &Item{
		id:          id,
		name:        name,
		description: description,
		price:       price,
		productType: productType,
	}

	if err := i.Validate(); err != nil {
		return nil, err
	}

	return i, nil
}

func (i *Item) GetID() string {
	return i.id
}

func (i *Item) GetName() string {
	return i.name
}

func (i *Item) GetDescription() string {
	return i.description
}

func (i *Item) GetPrice() int64 {
	return i.price
}

func (i *Item) GetProductType() ProductType {
	return i.productType
}

func (i *Item) GetQuantity() int64 {
	return i.quantity
}

func (i *Item) SetID(id string) {
	i.id = id
}

func (i *Item) SetName(name string) {
	i.name = name
}

func (i *Item) SetDescription(description string) {
	i.description = description
}

func (i *Item) SetPrice(price int64) {
	i.price = price
}

func (i *Item) SetProductType(productType ProductType) {
	i.productType = productType
}

func (i *Item) SetQuantity(quantity int64) {
	i.quantity = quantity
}

func (i *Item) ValidateName() *DomainError {
	if len(i.name) < 3 || len(i.name) > 100 {
		return NewDomainError(InvalidNameRange, "Name must be between 3 and 100 characters")
	}
	return nil
}

func (i *Item) ValidateDescription() *DomainError {
	if len(i.description) < 3 || len(i.description) > 500 {
		return NewDomainError(InvalidDescriptionRange, "Description must be between 3 and 500 characters")
	}
	return nil
}

func (i *Item) ValidatePrice() *DomainError {
	if i.price < 0 {
		return NewDomainError(InvalidPriceRange, "Price must be a positive value")
	}
	return nil
}

func (i *Item) ValidateProductType() *DomainError {
	switch i.productType {
	case Snack, Drink, Dessert, Accompaniment:
		return nil
	default:
		return NewDomainError(InvalidProductType, "Invalid product type")
	}
}

func (i *Item) Validate() *DomainError {
	if err := i.ValidateName(); err != nil {
		return err
	}

	if err := i.ValidateDescription(); err != nil {
		return err
	}

	if err := i.ValidatePrice(); err != nil {
		return err
	}

	if err := i.ValidateProductType(); err != nil {
		return err
	}

	return nil
}
