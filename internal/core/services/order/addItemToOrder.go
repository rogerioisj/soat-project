package order

import (
	"github.com/rogerioisj/soat-project/internal/core/domain"
	"github.com/rogerioisj/soat-project/internal/core/ports/repositories"
)

type AddItemToOrderService struct {
	or repositories.OrderRepositoryInterface
	ir repositories.ItemRepositoryInterface
}

func NewAddItemToOrderService(or repositories.OrderRepositoryInterface, ir repositories.ItemRepositoryInterface) *AddItemToOrderService {
	return &AddItemToOrderService{
		or: or,
		ir: ir,
	}
}

func (s *AddItemToOrderService) Execute(orderId string, itemIDs []string) *domain.DomainError {
	o, err := s.or.GetById(orderId)
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

	err = s.or.Update(o)

	if err != nil {
		return err
	}

	return nil
}
