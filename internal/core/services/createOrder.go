package services

import (
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

func (s *CreateOrderService) Execute() (*domain.Order, *domain.DomainError) {
	u := &domain.User{}

	err := s.ur.GetByCpf(u, u.GetCPF())

	if err != nil {
		return nil, err
	}

	o, err := domain.NewOrder(u)

	if err != nil {
		return nil, err
	}

	err = s.r.Create(o)

	if err != nil {
		return nil, err
	}

	return o, nil
}
