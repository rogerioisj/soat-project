package repositories

import "github.com/rogerioisj/soat-project/internal/core/domain"

type ItemRepositoryInterface interface {
	Create(item *domain.Item) error
	GetById(id string) (*domain.Item, error)
	ListByType(productType string) ([]*domain.Item, error)
}
