package domain

import (
	"testing"
)

func TestUserDomain(t *testing.T) {
	t.Log("User Domain Test Suite")
	t.Run("As a developer, I want to create a user domain model with validation rules", CreateUser)
	t.Run("As a developer, I want to create a user domain model with invalid name", CreateUserWithInvalidName)
	t.Run("As a developer, I want to create a user domain model with invalid email", CreateUserWithInvalidEmail)
	t.Run("As a developer, I want to create a user domain model with invalid CPF", CreateUserWithInvalidCPF)
}

func CreateUser(t *testing.T) {
	user, err := NewUser("1", "Fulano da Silva", "fulano@silvas.com", "12345678910")

	if err != nil {
		t.Error("No error expected, got:", err)
	}

	if user.GetID() != "1" {
		t.Error("Expected user ID to be '1', got:", user.GetID())
	}

	if user.GetName() != "Fulano da Silva" {
		t.Error("Expected user name to be 'Fulano da Silva', got:", user.GetName())
	}

	if user.GetEmail() != "fulano@silvas.com" {
		t.Error("Expected email to be 'fulano@silvas.com', got:", user.GetEmail())
	}

	if user.GetCPF() != "12345678910" {
		t.Error("Expected CPF to be '12345678910', got:", user.GetCPF())
	}
}

func CreateUserWithInvalidName(t *testing.T) {
	_, err := NewUser("1", "te", "fulano@silvas.com", "12345678910")

	if err == nil {
		t.Error("Error expected for invalid name, but got nil")
		return
	}
}

func CreateUserWithInvalidEmail(t *testing.T) {
	_, err := NewUser("1", "Fulano da Silva", "invalid-email", "12345678910")

	if err == nil {
		t.Error("Error expected for invalid email, but got nil")
		return
	}
}

func CreateUserWithInvalidCPF(t *testing.T) {
	_, err := NewUser("1", "Fulano da Silva", "fulano@silvas.com", "1")

	if err == nil {
		t.Error("Error expected for invalid cpf, but got nil")
		return
	}
}
