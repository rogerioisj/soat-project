package domain

import "testing"

var order *Order

func TestOrderDomain(t *testing.T) {
	t.Log("Order Domain Test Suite")
	t.Run("As a developer, I want to create an order domain model with validation rules", CreateOrder)
	t.Run("As a developer, I want to add items to an order domain model and check the value", AddItensToOrder)
	t.Run("As a developer, I want to update the status order to Received", SetStatusAsReceived)
	//t.Run("As a developer, I want to update the status order to Preparing", SetStatusAsPreparing)
}

func CreateOrder(t *testing.T) {
	u, err := NewUser("1", "Fulano da Silva", "fulano@silvas.com", "12345678910")

	if err != nil {
		t.Error("No error expected to create user, got:", err)
		return
	}

	_, err = NewOrder("1", u)

	if err != nil {
		t.Error("No error expected to create order, got:", err)
		return
	}
}

func AddItensToOrder(t *testing.T) {
	u, err := NewUser("1", "Fulano da Silva", "fulano@silvas.com", "12345678910")

	if err != nil {
		t.Error("No error expected to create user, got:", err)
		return
	}

	o, err := NewOrder("1", u)

	if err != nil {
		t.Error("No error expected to create order, got:", err)
		return
	}

	i, err := NewItem("1", "Coxinha", "Delicious chicken coxinha", 5000, Snack)
	if err != nil {
		t.Error("No error expected to create item, got:", err)
		return
	}

	i2, err := NewItem("2", "Refrigerante", "Refreshing soda", 3000, Drink)

	if err != nil {
		t.Error("No error expected to create item, got:", err)
		return
	}

	i3, err := NewItem("3", "Bolo de Chocolate", "Delicious chocolate cake", 7000, Dessert)

	if err != nil {
		t.Error("No error expected to create item, got:", err)
		return
	}

	i4, err := NewItem("4", "Batata Frita", "Crispy french fries", 4000, Accompaniment)

	if err != nil {
		t.Error("No error expected to create item, got:", err)
		return
	}

	err = o.AddItem(i)

	if err != nil {
		t.Error("No error expected to add item to order, got:", err)
		return
	}

	err = o.AddItem(i2)
	if err != nil {
		t.Error("No error expected to add item to order, got:", err)
		return
	}

	err = o.AddItem(i3)
	if err != nil {
		t.Error("No error expected to add item to order, got:", err)
		return
	}

	err = o.AddItem(i4)
	if err != nil {
		t.Error("No error expected to add item to order, got:", err)
		return
	}

	if len(o.GetItems()) != 4 {
		t.Error("Expected 4 items in order, got:", len(o.GetItems()))
		return
	}

	if o.GetPrice() != 19000 {
		t.Error("Expected total price to be 19000, got:", o.GetPrice())
		return
	}

}

func SetStatusAsReceived(t *testing.T) {
	u, err := NewUser("1", "Fulano da Silva", "fulano@silvas.com", "12345678910")

	if err != nil {
		t.Error("No error expected to create user, got:", err)
		return
	}

	o, err := NewOrder("1", u)

	if err != nil {
		t.Error("No error expected to create order, got:", err)
		return
	}

	i, err := NewItem("1", "Coxinha", "Delicious chicken coxinha", 5000, Snack)
	if err != nil {
		t.Error("No error expected to create item, got:", err)
		return
	}

	i2, err := NewItem("2", "Refrigerante", "Refreshing soda", 3000, Drink)

	if err != nil {
		t.Error("No error expected to create item, got:", err)
		return
	}

	i3, err := NewItem("3", "Bolo de Chocolate", "Delicious chocolate cake", 7000, Dessert)

	if err != nil {
		t.Error("No error expected to create item, got:", err)
		return
	}

	i4, err := NewItem("4", "Batata Frita", "Crispy french fries", 4000, Accompaniment)

	if err != nil {
		t.Error("No error expected to create item, got:", err)
		return
	}

	err = o.AddItem(i)

	if err != nil {
		t.Error("No error expected to add item to order, got:", err)
		return
	}

	err = o.AddItem(i2)
	if err != nil {
		t.Error("No error expected to add item to order, got:", err)
		return
	}

	err = o.AddItem(i3)
	if err != nil {
		t.Error("No error expected to add item to order, got:", err)
		return
	}

	err = o.AddItem(i4)
	if err != nil {
		t.Error("No error expected to add item to order, got:", err)
		return
	}

	err = o.SetStatus(Received)

	if err != nil {
		t.Error("No error expected to set order status to Received, got:", err)
		return
	}

}
