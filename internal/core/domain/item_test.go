package domain

import "testing"

func TestItemDomain(t *testing.T) {
	t.Log("Item Domain Test Suite")
	t.Run("As a developer, I want to create an item domain model with validation rules", CreateItem)
	t.Run("As a developer, I want to create an item domain model with invalid description", CreateItemWithInvalidDescription)
	t.Run("As a developer, I want to create an item domain model with invalid price", CreateItemWithInvalidPrice)
	t.Run("As a developer, I want to create an item domain model with invalid category", CreateItemWithInvalidCategory)
	t.Run("As a developer, I want to create an item domain model with invalid name", CreateItemWithInvalidName)
}

func CreateItem(t *testing.T) {
	item, err := NewItem("1", "Coxinha", "Delicious chicken coxinha", 5000, Snack)

	if err != nil {
		t.Error("No error expected, got:", err)
	}

	if item.GetID() != "1" {
		t.Error("Expected item ID to be '1', got:", item.GetID())
	}

	if item.GetName() != "Coxinha" {
		t.Error("Expected item name to be 'Coxinha', got:", item.GetName())
	}

	if item.GetDescription() != "Delicious chicken coxinha" {
		t.Error("Expected description to be 'Delicious chicken coxinha', got:", item.GetDescription())
	}

	if item.GetPrice() != 5000 {
		t.Error("Expected price to be 5000, got:", item.GetPrice())
	}

	if item.GetProductType() != Snack {
		t.Error("Expected product type to be 'snack', got:", item.GetProductType())
	}
}

func CreateItemWithInvalidDescription(t *testing.T) {
	_, err := NewItem("1", "Coxinha", "A", 5000, Snack)

	if err == nil {
		t.Error("Error expected for invalid description, but got nil")
		return
	}

	if !err.Is(InvalidDescriptionRange) {
		t.Error("Expected error code InvalidDescriptionRange, got:", err.GetCode())
	}

	if err.GetCode() != InvalidDescriptionRange {
		t.Error("Expected error code InvalidDescriptionRange, got:", err.GetCode())
	}

	if err.Error() != "Description must be between 3 and 500 characters" {
		t.Error("Expected error message 'Description must be between 3 and 500 characters', got:", err.Error())
	}
}

func CreateItemWithInvalidPrice(t *testing.T) {
	_, err := NewItem("1", "Coxinha", "Delicious chicken coxinha", -100, Snack)

	if err == nil {
		t.Error("Error expected for invalid price, but got nil")
		return
	}

	if !err.Is(InvalidPriceRange) {
		t.Error("Expected error code InvalidPriceRange, got:", err.GetCode())
	}

	if err.GetCode() != InvalidPriceRange {
		t.Error("Expected error code InvalidPriceRange, got:", err.GetCode())
	}

	if err.Error() != "Price must be a positive value" {
		t.Error("Expected error message 'Price must be a positive value', got:", err.Error())
	}
}

func CreateItemWithInvalidCategory(t *testing.T) {
	_, err := NewItem("1", "Coxinha", "Delicious chicken coxinha", 5000, "invalid_category")

	if err == nil {
		t.Error("Error expected for invalid category, but got nil")
		return
	}

	if !err.Is(InvalidProductType) {
		t.Error("Expected error code InvalidProductType, got:", err.GetCode())
	}

	if err.GetCode() != InvalidProductType {
		t.Error("Expected error code InvalidProductType, got:", err.GetCode())
	}

	if err.Error() != "Invalid product type" {
		t.Error("Expected error message 'Invalid product type', got:", err.Error())
	}
}

func CreateItemWithInvalidName(t *testing.T) {
	_, err := NewItem("1", "C", "Delicious chicken coxinha", 5000, Snack)

	if err == nil {
		t.Error("Error expected for invalid name, but got nil")
		return
	}

	if !err.Is(InvalidNameRange) {
		t.Error("Expected error code InvalidNameRange, got:", err.GetCode())
	}

	if err.GetCode() != InvalidNameRange {
		t.Error("Expected error code InvalidNameRange, got:", err.GetCode())
	}

	if err.Error() != "Name must be between 3 and 100 characters" {
		t.Error("Expected error message 'Name must be between 3 and 100 characters', got:", err.Error())
	}
}
