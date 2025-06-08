package repositories

import "github.com/rogerioisj/soat-project/internal/core/domain"

type ItemRepositoryInterface interface {
	Create(item *domain.Item) *domain.DomainError
	GetById(id string, item *domain.Item) *domain.DomainError
	ListByType(productType string, page, limit int32, itemList *[]domain.Item) *domain.DomainError
	Update(item *domain.Item, id string) *domain.DomainError
	Delete(itemId string) *domain.DomainError
}
