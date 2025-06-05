package repositories

import "github.com/rogerioisj/soat-project/internal/core/domain"

type UserRepositoryInterface interface {
	Create(user *domain.User) *domain.DomainError
	GetByCpf(user *domain.User, cpf string) *domain.DomainError
	GetByEmail(user *domain.User, email string) *domain.DomainError
	GetByCpfOrEmail(user *domain.User) *domain.DomainError
	GetGuestUser(user *domain.User) *domain.DomainError
}
