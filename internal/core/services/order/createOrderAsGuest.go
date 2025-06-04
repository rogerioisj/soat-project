package order

import (
	"github.com/rogerioisj/soat-project/internal/core/domain"
	"github.com/rogerioisj/soat-project/internal/core/ports/repositories"
)

type CreateOrderAsGuestService struct {
	r  repositories.OrderRepositoryInterface
	ur repositories.UserRepositoryInterface
}

func NewCreateOrderAsGuestService(r repositories.OrderRepositoryInterface, ur repositories.UserRepositoryInterface) *CreateOrderAsGuestService {
	return &CreateOrderAsGuestService{
		r:  r,
		ur: ur,
	}
}

func (s *CreateOrderAsGuestService) Execute() (*domain.Order, *domain.DomainError) {
	u := &domain.User{}

	err := s.ur.GetGuestUser(u)

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
