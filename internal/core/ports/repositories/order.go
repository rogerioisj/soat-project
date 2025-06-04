package repositories

import "github.com/rogerioisj/soat-project/internal/core/domain"

type OrderRepositoryInterface interface {
	Create(order *domain.Order) *domain.DomainError
	GetById(id string) (*domain.Order, *domain.DomainError)
	UpdateStatus(id string, status domain.OrderStatus) *domain.DomainError
}
