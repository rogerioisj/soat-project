package item

import (
	"github.com/rogerioisj/soat-project/internal/core/domain"
	"github.com/rogerioisj/soat-project/internal/core/ports/repositories"
)

type GetItensBasedOnCategoryService struct {
	ir repositories.ItemRepositoryInterface
}

func NewGetItensBasedOnCategoryService(ir repositories.ItemRepositoryInterface) *GetItensBasedOnCategoryService {
	return &GetItensBasedOnCategoryService{
		ir: ir,
	}
}

func (s *GetItensBasedOnCategoryService) Execute(category string, page, limit int32) ([]*domain.Item, *domain.DomainError) {
	items, err := s.ir.ListByType(category, page, limit)
	if err != nil {
		return nil, err
	}

	return items, nil
}
