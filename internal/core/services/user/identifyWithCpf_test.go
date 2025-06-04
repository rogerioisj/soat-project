package user

import (
	"github.com/rogerioisj/soat-project/internal/adapters/outbound/repositories/memory"
	"github.com/rogerioisj/soat-project/internal/core/domain"
	"github.com/rogerioisj/soat-project/internal/core/ports/repositories"
	"testing"
)

var (
	urs  repositories.UserRepositoryInterface
	iwcs *IdentifyWithCpfService
)

func TestIdentifyWithCpfService(t *testing.T) {
	t.Log("IdentifyWithCpfService test suite")
	t.Run("As a developer, I want to test the IdentifyWithCpfService so that I can ensure it works correctly", loadIdentifyWithCpfServiceInstance)
	t.Run("Given a valid CPF, when I identify it, then it should return the user successfully", identifyWithCpfSuccessfully)
	t.Run("Given an invalid CPF, when I try to identify it, then it should return an error", identifyWithCpfFailure)
}

func loadIdentifyWithCpfServiceInstance(t *testing.T) {
	urs = memory.NewUserRepositoryMock()

	if urs == nil {
		t.Fatal("UserRepository should not be nil")
	}

	u, err := domain.NewUser("", "Roger", "teste@teste.com", "12345678900")

	if err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}

	err = urs.Create(u)

	if err != nil {
		t.Fatalf("Failed to create user in repository: %v", err)
	}

	iwcs = NewIdentifyWithCpfService(urs)

	if iwcs == nil {
		t.Fatal("IdentifyWithCpfService should not be nil")
	}
}

func identifyWithCpfSuccessfully(t *testing.T) {
	u, err := iwcs.Execute("12345678900")

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if u == nil {
		t.Error("Expected user to be identified, got nil")
		return
	}

	if u.GetName() != "Roger" {
		t.Errorf("Expected user name to be 'Roger', got '%s'", u.GetName())
	}

	if u.GetCPF() != "12345678900" {
		t.Errorf("Expected user CPF to be '12345678900', got '%s'", u.GetCPF())
	}

	if u.GetEmail() != "teste@teste.com" {
		t.Errorf("Expected user email to be 'teste@teste.com', got %s", u.GetEmail())
	}
}

func identifyWithCpfFailure(t *testing.T) {
	_, err := iwcs.Execute("00000000000")

	if err == nil {
		t.Error("Expected an error, got nil")
		return
	}

	if err.Code != domain.UserNotFound {
		t.Errorf("Expected error code %s, got %s", domain.UserNotFound, err.Code)
	}

	if err.Message != "User with this CPF not found" {
		t.Errorf("Expected error message 'User with this CPF not found', got '%s'", err.Message)
	}
}
