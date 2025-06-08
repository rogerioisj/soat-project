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

func (s *GetItensBasedOnCategoryService) Execute(category string, page, limit int32, itens *[]domain.Item) *domain.DomainError {

	err := s.ir.ListByType(category, page, limit, itens)
	if err != nil {
		return err
	}

	return nil
}
