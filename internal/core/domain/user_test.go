package domain

import (
	"testing"
)

func TestUserDomain(t *testing.T) {
	t.Log("User Domain Test Suite")
	t.Run("As a developer, I want to create a user domain model with validation rules", CreateUser)
	t.Run("As a developer, I want to create a user domain model with invalid name", CreateUserWithInvalidName)
	t.Run("As a developer, I want to create a user domain model with invalid email", CreateUserWithInvalidEmail)
	t.Run("As a developer, I want to create a user domain model with invalid email range", CreateUserWithInvalidEmailRange)
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

	if !err.Is(InvalidNameRange) {
		t.Error("Expected error code InvalidNameRange, got:", err.GetCode())
	}

	if err.Error() != "Name must be between 3 and 100 characters" {
		t.Error("Expected error message 'Name must be between 3 and 100 characters', got:", err.Error())
	}
}

func CreateUserWithInvalidEmail(t *testing.T) {
	_, err := NewUser("1", "Fulano da Silva", "invalid-email", "12345678910")

	if err == nil {
		t.Error("Error expected for invalid email, but got nil")
		return
	}

	if !err.Is(InvalidEmailFormat) {
		t.Error("Expected error code InvalidEmailFormat, got:", err.GetCode())
	}

	if err.Error() != "Email format is invalid" {
		t.Error("Expected error message 'Email format is invalid', got:", err.Error())
	}
}

func CreateUserWithInvalidEmailRange(t *testing.T) {
	_, err := NewUser("1", "Fulano da Silva", "i", "12345678910")

	if err == nil {
		t.Error("Error expected for invalid email, but got nil")
		return
	}

	if !err.Is(InvalidEmailRange) {
		t.Error("Expected error code InvalidEmailRange, got:", err.GetCode())
	}

	if err.Error() != "Email must be between 3 and 100 characters" {
		t.Error("Expected error message 'Email must be between 3 and 100 characters', got:", err.Error())
	}
}

func CreateUserWithInvalidCPF(t *testing.T) {
	_, err := NewUser("1", "Fulano da Silva", "fulano@silvas.com", "1")

	if err == nil {
		t.Error("Error expected for invalid cpf, but got nil")
		return
	}

	if !err.Is(InvalidCPF) {
		t.Error("Expected error code InvalidCPF, got:", err.GetCode())
	}

	if err.Error() != "CPF must contain exactly 11 digits" {
		t.Error("Expected error message 'CPF must contain exactly 11 digits', got:", err.Error())
	}
}
