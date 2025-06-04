package repositories

import "github.com/rogerioisj/soat-project/internal/core/domain"

type UserRepositoryInterface interface {
	Create(user *domain.User) error
	GetByCpf(cpf string) (*domain.User, error)
	GetByEmail(email string) (*domain.User, error)
}
