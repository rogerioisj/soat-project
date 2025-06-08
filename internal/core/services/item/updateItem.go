package item

import (
	"github.com/rogerioisj/soat-project/internal/core/domain"
	"github.com/rogerioisj/soat-project/internal/core/ports/repositories"
)

type UpdateItemService struct {
	ir repositories.ItemRepositoryInterface
}

func NewUpdateItemService(ir repositories.ItemRepositoryInterface) *UpdateItemService {
	return &UpdateItemService{
		ir: ir,
	}
}

func (s *UpdateItemService) Execute(item *domain.Item, itemId string) *domain.DomainError {
	if item == nil {
		return domain.NewDomainError(domain.NilItemInstance, "item cannot be nil")
	}

	if item.GetName() == "" {
		return domain.NewDomainError(domain.InvalidItemName, "item name cannot be empty")
	}

	if item.GetPrice() <= 0 {
		return domain.NewDomainError(domain.InvalidItemPrice, "item price must be greater than zero")
	}

	if item.GetProductType() != domain.Drink && item.GetProductType() != domain.Dessert && item.GetProductType() != domain.Snack && item.GetProductType() != domain.Accompaniment {
		return domain.NewDomainError(domain.InvalidProductType, "item category must be one of: drink, dessert, snack, or accompaniment")
	}

	loadedItem := &domain.Item{}

	err := s.ir.GetById(itemId, loadedItem)

	if err != nil {
		return domain.NewDomainError(domain.ItemNotFound, "item not found")
	}

	err = s.ir.Update(item, itemId)

	if err != nil {
		return domain.NewDomainError(domain.ItemNotFound, "error updating item")
	}

	return nil
}
