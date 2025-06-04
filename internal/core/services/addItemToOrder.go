package services

import (
	"github.com/rogerioisj/soat-project/internal/core/domain"
	"github.com/rogerioisj/soat-project/internal/core/ports/repositories"
)

type AddItemToOrderService struct {
	r  repositories.OrderRepositoryInterface
	ir repositories.ItemRepositoryInterface
}

func NewAddItemToOrderService() *AddItemToOrderService {
	return &AddItemToOrderService{}
}

func (s *AddItemToOrderService) Execute(orderId string, itemIDs []string) *domain.DomainError {
	o, err := s.r.GetById(orderId)
	if err != nil {
		return err
	}

	for _, itemID := range itemIDs {
		item, err := s.ir.GetById(itemID)

		if err != nil {
			return domain.NewDomainError(domain.ItemNotFound, "Item not found with ID: "+itemID)
		}

		err = o.AddItem(item)

		if err != nil {
			return err
		}
	}

	err = s.r.Update(o)

	if err != nil {
		return err
	}

	return nil
}
