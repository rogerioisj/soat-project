package order

import (
	"github.com/rogerioisj/soat-project/internal/adapters/inbound/http/dtos"
	"github.com/rogerioisj/soat-project/internal/core/domain"
	"github.com/rogerioisj/soat-project/internal/core/ports/repositories"
)

type CreateOrderService struct {
	r  repositories.OrderRepositoryInterface
	ur repositories.UserRepositoryInterface
}

func NewCreateOrderService(r repositories.OrderRepositoryInterface, ur repositories.UserRepositoryInterface) *CreateOrderService {
	return &CreateOrderService{
		r:  r,
		ur: ur,
	}
}

func (s *CreateOrderService) Execute(userId string, itens *[]dtos.Product, o *domain.Order) *domain.DomainError {
	u := &domain.User{}

	err := s.ur.GetByID(u, userId)

	if err != nil {
		return err
	}

	o.SetUser(*u)

	var itemOrder []domain.ItemOrderElement

	for _, item := range *itens {
		io := domain.ItemOrderElement{
			ItemID:   item.ID,
			Quantity: item.Quantity,
		}
		itemOrder = append(itemOrder, io)
	}

	o.AddItemOrder(itemOrder)

	err = s.r.Create(o)

	if err != nil {
		return err
	}

	return nil
}
