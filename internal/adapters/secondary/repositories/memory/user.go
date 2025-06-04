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
