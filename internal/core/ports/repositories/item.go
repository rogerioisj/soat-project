package repositories

import "github.com/rogerioisj/soat-project/internal/core/domain"

type ItemRepositoryInterface interface {
	Create(item *domain.Item) *domain.DomainError
	GetById(id string) (*domain.Item, *domain.DomainError)
	ListByType(productType string, page, limit int32) ([]*domain.Item, *domain.DomainError)
}
