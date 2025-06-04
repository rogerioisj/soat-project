package memory

import (
	"github.com/google/uuid"
	"github.com/rogerioisj/soat-project/internal/core/domain"
)

type OrderRepositoryMock struct {
	orders map[string]*domain.Order
}

func NewOrderRepositoryMock() *OrderRepositoryMock {
	return &OrderRepositoryMock{
		orders: make(map[string]*domain.Order),
	}
}

func (r *OrderRepositoryMock) Create(order *domain.Order) *domain.DomainError {
	if order == nil {
		return domain.NewDomainError(domain.InvalidOrder, "Order cannot be nil")
	}

	id := uuid.New().String()

	order.SetId(id)

	r.orders[order.GetID()] = order
	return nil
}
