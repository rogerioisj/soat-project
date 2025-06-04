package services

import (
	"github.com/rogerioisj/soat-project/internal/adapters/secondary/repositories/memory"
	"github.com/rogerioisj/soat-project/internal/core/domain"
	"github.com/rogerioisj/soat-project/internal/core/ports/repositories"
	"strings"
	"testing"
)

var (
	ur repositories.UserRepositoryInterface
	s  *CreateUserService
)

func TestCreateUserService(t *testing.T) {
	t.Log("CreateUserService test suite")
	t.Run("As a developer, I want to test the CreateUserService so that I can ensure it works correctly", createService)
	t.Run("Given a valid user, when I create it, then it should be created successfully", createUserSuccessfully)
	t.Run("Given an existing user, when I try to create it again, then it should return an error", createUserFailure)
	t.Run("Given an invalid name, when I try to create it, then it should return an error", createUserWithInvalidName)
	t.Run("Given an invalid email, when I try to create it, then it should return an error", createUserWithInvalidEmail)
	t.Run("Given an invalid cpf, when I try to create it, then it should return an error", createUserWithInvalidCpf)
}

func createService(t *testing.T) {
	ur = memory.NewUserRepositoryMock()

	if ur == nil {
		t.Fatal("UserRepository should not be nil")
	}

	s = NewCreateUserService(ur)
	if s == nil {
		t.Fatal("CreateUserService should not be nil")
	}
}

func createUserSuccessfully(t *testing.T) {

	u, err := s.Execute("Roger", "teste@teste.com", "12345678900")

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if u == nil {
		t.Error("Expected user to be created, got nil")
		return
	}

	if u.GetName() != "Roger" {
		t.Errorf("Expected user name to be 'Roger', got '%s'", u.GetName())
	}

	if u.GetEmail() != "teste@teste.com" {
		t.Errorf("Expected user email to be 'teste@teste.com', got '%s'", u.GetEmail())
	}

	if u.GetCPF() != "12345678900" {
		t.Errorf("Expected user CPF to be '12345678900', got '%s'", u.GetCPF())
	}

	if u.GetID() == "" {
		t.Error("Expected user ID to be set, got empty string")
	}
}

func createUserFailure(t *testing.T) {
	_, err := s.Execute("Roger", "teste@teste.com", "12345678900")

	if err == nil {
		t.Error("Expected error for existing user, got nil")
		return
	}

	if err.GetCode() != domain.UserAlreadyExists {
		t.Errorf("Expected error code UserAlreadyExists, got %v", err.GetCode())
	}

	if err.Error() != "User with this CPF or email already exists" {
		t.Errorf("Expected error message 'User with this CPF or email already exists', got '%s'", err.Error())
	}
}

func createUserWithInvalidName(t *testing.T) {
	_, err := s.Execute("Ro", "invalid-email", "12345")

	if err == nil {
		t.Error("Expected error for invalid user data, got nil")
		return
	}

	if err.GetCode() != domain.InvalidNameRange {
		t.Errorf("Expected error code InvalidUserData, got %v", err.GetCode())
	}

	if !strings.Contains(err.Error(), "Name must be between 3 and 100 characters") {
		t.Errorf("Expected error message to contain 'Name must be between 3 and 100 characters', got '%s'", err.Error())
	}
}

func createUserWithInvalidEmail(t *testing.T) {
	_, err := s.Execute("Roger Inacio", "invalid-email", "12345")

	if err == nil {
		t.Error("Expected error for invalid user data, got nil")
		return
	}

	if err.GetCode() != domain.InvalidEmailFormat {
		t.Errorf("Expected error code InvalidEmailFormat, got %v", err.GetCode())
	}

	if !strings.Contains(err.Error(), "Email format is invalid") {
		t.Errorf("Expected error message to contain 'Email format is invalid', got '%s'", err.Error())
	}
}

func createUserWithInvalidCpf(t *testing.T) {
	_, err := s.Execute("Roger Inacio", "teste2@teste.com", "12345")

	if err == nil {
		t.Error("Expected error for invalid user data, got nil")
		return
	}

	if err.GetCode() != domain.InvalidCPF {
		t.Errorf("Expected error code InvalidCPF, got %v", err.GetCode())
	}

	if !strings.Contains(err.Error(), "CPF must contain exactly 11 digits") {
		t.Errorf("Expected error message to contain 'CPF must contain exactly 11 digits', got '%s'", err.Error())
	}
}
