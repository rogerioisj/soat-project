package memory

import (
	"github.com/google/uuid"
	"github.com/rogerioisj/soat-project/internal/core/domain"
)

type UserRepositoryMock struct {
	users map[string]*domain.User
}

func NewUserRepositoryMock() *UserRepositoryMock {
	return &UserRepositoryMock{
		users: make(map[string]*domain.User),
	}
}

func (r *UserRepositoryMock) Create(user *domain.User) *domain.DomainError {
	if user.GetEmail() == "error@error.com" {
		return domain.NewDomainError(domain.InvalidEmailFormat, "Invalid user email")
	}

	checkForCpf, _ := r.searchUserByCpf(user.GetCPF())
	checkForEmail, _ := r.searchUserByEmail(user.GetEmail())

	if checkForCpf || checkForEmail {
		return domain.NewDomainError(domain.UserAlreadyExists, "User with this CPF or email already exists")
	}

	user.SetID(uuid.New().String())
	r.users[user.GetID()] = user

	return nil
}

func (r *UserRepositoryMock) GetByCpf(user *domain.User, cpf string) *domain.DomainError {
	exists, foundUser := r.searchUserByCpf(cpf)

	if !exists {
		return domain.NewDomainError(domain.UserNotFound, "User with this CPF not found")
	}

	user.SetID(foundUser.GetID())
	user.SetName(foundUser.GetName())
	user.SetEmail(foundUser.GetEmail())
	user.SetCPF(foundUser.GetCPF())

	return nil
}

func (r *UserRepositoryMock) GetByEmail(user *domain.User) *domain.DomainError {
	exists, foundUser := r.searchUserByEmail(user.GetEmail())

	if !exists {
		return domain.NewDomainError(domain.UserNotFound, "User with this email not found")
	}

	user.SetID(foundUser.GetID())
	user.SetName(foundUser.GetName())
	user.SetEmail(foundUser.GetEmail())
	user.SetCPF(foundUser.GetCPF())

	return nil
}

func (r *UserRepositoryMock) GetByCpfOrEmail(user *domain.User) *domain.DomainError {
	if user.GetName() == "Error" {
		return domain.NewDomainError(domain.InvalidNameRange, "Invalid user name")
	}

	existsByCpf, foundUserByCpf := r.searchUserByCpf(user.GetCPF())

	if existsByCpf {
		user.SetID(foundUserByCpf.GetID())
		user.SetName(foundUserByCpf.GetName())
		user.SetEmail(foundUserByCpf.GetEmail())
		user.SetCPF(foundUserByCpf.GetCPF())
		return nil
	}

	existsByEmail, foundUserByEmail := r.searchUserByEmail(user.GetEmail())

	if existsByEmail {
		user.SetID(foundUserByEmail.GetID())
		user.SetName(foundUserByEmail.GetName())
		user.SetEmail(foundUserByEmail.GetEmail())
		user.SetCPF(foundUserByEmail.GetCPF())
		return nil
	}

	return domain.NewDomainError(domain.UserNotFound, "User with this CPF or email not found")
}

func (r *UserRepositoryMock) searchUserByCpf(cpf string) (bool, *domain.User) {
	for user := range r.users {
		if r.users[user].GetCPF() == cpf {
			return true, r.users[user]
		}
	}

	return false, nil
}

func (r *UserRepositoryMock) searchUserByEmail(email string) (bool, *domain.User) {
	for user := range r.users {
		if r.users[user].GetEmail() == email {
			return true, r.users[user]
		}
	}

	return false, nil
}
