package domain

import "testing"

var O *Order

func TestOrderDomain(t *testing.T) {
	t.Log("Order Domain Test Suite")
	t.Run("As a developer, I want to create an order domain model with validation rules", CreateOrder)
	t.Run("As a developer, I want to fail to update the status order to Received without items", FailToSetStatusAsReceivedWithoutItens)
	t.Run("As a developer, I want to fail to update the status order to Received when it has no items with value", FailToSetStatusAsReceiveWithoutItensWithValue)
	t.Run("As a developer, I want to cancel an order domain model", CancelOrder)
	t.Run("As a developer, I want to fail to update a cancelled order", FailToUpdateACancelledOrder)
	t.Run("As a developer, I want to add items to an order domain model and check the value", AddItensToOrder)
	t.Run("As a developer, I want to fail to set order as Ready when is not in Preparing status", FailToSetOrderAsReady)
	t.Run("As a developer, I want to fail to set order as Preparing when is not in Received status", FailToSetOrderAsPreparing)
	t.Run("As a developer, I want to update the status order to Received", SetStatusAsReceived)
	t.Run("As a developer, I want to fail to add items to an order after it is Received", FailToAddItensToOrderAfterReceivedStatus)
	t.Run("As a developer, I want to fail to remove an unexisting item from an order", FailToRemoveUnexistingItemFromOrder)
	t.Run("As a developer, I want to fail to cancel an order that is not in Building status", FailToCancelOrder)
	t.Run("As a developer, I want to fail to update the status order to Received when it has no items", FailToSetStatusAsReceived)
	t.Run("As a developer, I want to fail to update the status order to Building when it has items", FailToSetStatusAsBuilding)
	t.Run("As a developer, I want to fail to set an invalid status for the order", FailToSetAInvalidStatus)
	t.Run("As a developer, I want to update the status order to Preparing", SetStatusAsPreparing)
	t.Run("As a developer, I want to fail to update the status order to Done when it is not in Ready status", FailToSetOrderAsDone)
	t.Run("As a developer, I want to update the status order to Ready", SetStatusAsReady)
	t.Run("As a developer, I want to update the status order to Done", SetStatusAsDone)
	t.Run("As a developer, I want to fail to change the status of an order after it is Done", FailToChangeStatusAfterDone)
}

func CreateOrder(t *testing.T) {
	u, err := NewUser("1", "Fulano da Silva", "fulano@silvas.com", "12345678910")

	if err != nil {
		t.Error("No error expected to create user, got:", err)
		return
	}

	O, err = NewOrder("1", u)

	if err != nil {
		t.Error("No error expected to create order, got:", err)
		return
	}

	if O.GetID() != "1" {
		t.Error("Expected order ID to be '1', got:", O.GetID())
		return
	}

	if O.GetUser().GetID() != "1" {
		t.Error("Expected order user ID to be '1', got:", O.GetUser().GetID())
		return
	}
}

func FailToSetStatusAsReceivedWithoutItens(t *testing.T) {
	err := O.SetStatus(Received)

	if err == nil {
		t.Error("Error expected to set order status to Received, but got nil")
		return
	}

	if !err.Is(InvalidOrderStatus) {
		t.Error("Expected error code InvalidOrderStatus, got:", err.GetCode())
		return
	}
}

func FailToSetStatusAsReceiveWithoutItensWithValue(t *testing.T) {
	i, err := NewItem("1", "Agua", "Agua da Torneira", 0, Drink)
	if err != nil {
		t.Error("No error expected to create item, got:", err)
		return
	}

	err = O.AddItem(i)

	if err != nil {
		t.Error("No error expected to add item to order, got:", err)
		return
	}

	err = O.SetStatus(Received)

	if err == nil {
		t.Error("Error expected to set order status to Received, but got nil")
		return
	}

	err = O.RemoveItem(i.id)

	if err != nil {
		t.Error("No error expected to remove item from order, got:", err)
		return
	}
}

func CancelOrder(t *testing.T) {
	err := O.SetStatus(Cancelled)

	if err != nil {
		t.Error("No error expected to set order status to Cancelled, got:", err)
		return
	}

	if O.GetStatus() != Cancelled {
		t.Error("Expected order status to be Cancelled, got:", O.GetStatus())
		return
	}
}

func FailToUpdateACancelledOrder(t *testing.T) {
	err := O.SetStatus(Preparing)

	if err == nil {
		t.Error("Error expected to set order status to Building, but got nil")
		return
	}

	if !err.Is(InvalidOrderStatus) {
		t.Error("Expected error code InvalidOrderStatus, got:", err.GetCode())
		return
	}

	CreateOrder(t)
}

func AddItensToOrder(t *testing.T) {
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

	err = O.AddItem(i)

	if err != nil {
		t.Error("No error expected to add item to order, got:", err)
		return
	}

	err = O.AddItem(i2)
	if err != nil {
		t.Error("No error expected to add item to order, got:", err)
		return
	}

	err = O.AddItem(i3)
	if err != nil {
		t.Error("No error expected to add item to order, got:", err)
		return
	}

	err = O.AddItem(i4)
	if err != nil {
		t.Error("No error expected to add item to order, got:", err)
		return
	}

	if len(O.GetItems()) != 4 {
		t.Error("Expected 4 items in order, got:", len(O.GetItems()))
		return
	}

	if O.GetPrice() != 19000 {
		t.Error("Expected total price to be 19000, got:", O.GetPrice())
		return
	}

}

func FailToSetOrderAsReady(t *testing.T) {
	err := O.SetStatus(Ready)

	if err == nil {
		t.Error("Error expected to set order status to Ready, but got nil")
		return
	}

	if !err.Is(InvalidOrderStatus) {
		t.Error("Expected error code InvalidOrderStatus, got:", err.GetCode())
		return
	}
}

func FailToSetOrderAsPreparing(t *testing.T) {
	err := O.SetStatus(Preparing)

	if err == nil {
		t.Error("Error expected to set order status to Preparing, but got nil")
		return
	}

	if !err.Is(InvalidOrderStatus) {
		t.Error("Expected error code InvalidOrderStatus, got:", err.GetCode())
		return
	}
}

func SetStatusAsReceived(t *testing.T) {
	err := O.SetStatus(Received)

	if err != nil {
		t.Error("No error expected to set order status to Received, got:", err)
		return
	}
}

func FailToAddItensToOrderAfterReceivedStatus(t *testing.T) {
	i, err := NewItem("5", "Salada", "Fresh salad", 2500, Accompaniment)

	if err != nil {
		t.Error("No error expected to create item, got:", err)
		return
	}

	err = O.AddItem(i)

	if err == nil {
		t.Error("Error expected to add item to order after Received status, but got nil")
		return
	}

	if !err.Is(InvalidOrderStatus) {
		t.Error("Expected error code InvalidOrderStatus, got:", err.GetCode())
		return
	}
}

func FailToRemoveUnexistingItemFromOrder(t *testing.T) {
	err := O.RemoveItem("999")

	if err == nil {
		t.Error("Error expected to remove unexisting item from order, but got nil")
		return
	}

	if !err.Is(ProductNotFoundInOrder) {
		t.Error("Expected error code ProductNotFoundInOrder, got:", err.GetCode())
		return
	}
}

func FailToCancelOrder(t *testing.T) {
	err := O.SetStatus(Cancelled)

	if err == nil {
		t.Error("Error expected to set order status to Cancelled, but got nil")
		return
	}

	if !err.Is(InvalidOrderStatus) {
		t.Error("Expected error code InvalidOrderStatus, got:", err.GetCode())
		return
	}
}

func FailToSetStatusAsReceived(t *testing.T) {
	err := O.SetStatus(Received)

	if err == nil {
		t.Error("Error expected to set order status to Received, but got nil")
		return
	}

	if !err.Is(InvalidOrderStatus) {
		t.Error("Expected error code InvalidOrderStatus, got:", err.GetCode())
		return
	}
}

func FailToSetStatusAsBuilding(t *testing.T) {
	err := O.SetStatus(Building)

	if err == nil {
		t.Error("Error expected to set order status to Building, but got nil")
		return
	}

	if !err.Is(InvalidOrderStatus) {
		t.Error("Expected error code InvalidOrderStatus, got:", err.GetCode())
		return
	}
}

func FailToSetAInvalidStatus(t *testing.T) {
	err := O.SetStatus("InvalidStatus")

	if err == nil {
		t.Error("Error expected to set order status to InvalidStatus, but got nil")
		return
	}

	if !err.Is(InvalidOrderStatus) {
		t.Error("Expected error code InvalidOrderStatus, got:", err.GetCode())
		return
	}
}

func SetStatusAsPreparing(t *testing.T) {
	err := O.SetStatus(Preparing)

	if err != nil {
		t.Error("No error expected to set order status to Preparing, got:", err)
		return
	}

	if O.GetStatus() != Preparing {
		t.Error("Expected order status to be Preparing, got:", O.GetStatus())
		return
	}
}

func FailToSetOrderAsDone(t *testing.T) {
	err := O.SetStatus(Done)

	if err == nil {
		t.Error("Error expected to set order status to Done, but got nil")
		return
	}

	if !err.Is(InvalidOrderStatus) {
		t.Error("Expected error code InvalidOrderStatus, got:", err.GetCode())
		return
	}
}

func SetStatusAsReady(t *testing.T) {
	err := O.SetStatus(Ready)

	if err != nil {
		t.Error("No error expected to set order status to Ready, got:", err)
		return
	}

	if O.GetStatus() != Ready {
		t.Error("Expected order status to be Ready, got:", O.GetStatus())
		return
	}
}

func SetStatusAsDone(t *testing.T) {
	err := O.SetStatus(Done)

	if err != nil {
		t.Error("No error expected to set order status to Done, got:", err)
		return
	}

	if O.GetStatus() != Done {
		t.Error("Expected order status to be Done, got:", O.GetStatus())
		return
	}
}

func FailToChangeStatusAfterDone(t *testing.T) {
	err := O.SetStatus(Preparing)

	if err == nil {
		t.Error("Error expected to set order status to Preparing after Done, but got nil")
		return
	}

	if !err.Is(InvalidOrderStatus) {
		t.Error("Expected error code InvalidOrderStatus, got:", err.GetCode())
		return
	}
}
