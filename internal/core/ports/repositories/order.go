package repositories

import "github.com/rogerioisj/soat-project/internal/core/domain"

type OrderRepositoryInterface interface {
	Create(order *domain.Order) error
	GetById(id string) (*domain.Order, error)
	UpdateStatus(id string, status domain.OrderStatus) error
}
