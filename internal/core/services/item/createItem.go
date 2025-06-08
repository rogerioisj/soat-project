package item

import (
	"github.com/rogerioisj/soat-project/internal/core/domain"
	"github.com/rogerioisj/soat-project/internal/core/ports/repositories"
)

type CreateItemService struct {
	ir repositories.ItemRepositoryInterface
}

func NewCreateItemService(ir repositories.ItemRepositoryInterface) *CreateItemService {
	return &CreateItemService{
		ir: ir,
	}
}

func (s *CreateItemService) Execute(item *domain.Item) *domain.DomainError {
	if item == nil {
		return domain.NewDomainError(domain.NilItemInstance, "item cannot be nil")
	}

	if item.GetName() == "" {
		return domain.NewDomainError(domain.InvalidItemName, "item cannot be nil")
	}

	if item.GetPrice() <= 0 {
		return domain.NewDomainError(domain.InvalidItemPrice, "item price must be greater than zero")
	}

	if item.GetProductType() != domain.Drink && item.GetProductType() != domain.Dessert && item.GetProductType() != domain.Snack && item.GetProductType() != domain.Accompaniment {
		return domain.NewDomainError(domain.InvalidProductType, "item category must be one of: drink, dessert, snack, or accompaniment")
	}

	err := s.ir.Create(item)
	if err != nil {
		return err
	}

	return nil
}
