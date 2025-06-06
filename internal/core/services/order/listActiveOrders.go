package order

import (
	"github.com/rogerioisj/soat-project/internal/core/domain"
	"github.com/rogerioisj/soat-project/internal/core/ports/repositories"
)

type ListActiveOrdersService struct {
	r repositories.OrderRepositoryInterface
}

func NewListOrdersService(r repositories.OrderRepositoryInterface) *ListActiveOrdersService {
	return &ListActiveOrdersService{
		r: r,
	}
}

func (s *ListActiveOrdersService) Execute(orders *[]domain.Order, page, limit int) *domain.DomainError {
	if page < 1 {
		page = 1
	}

	if limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit

	err := s.r.ListActives(orders, offset, limit)
	if err != nil {
		return domain.NewDomainError("error listing orders", err.Error())
	}

	return nil
}
