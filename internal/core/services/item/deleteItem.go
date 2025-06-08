package item

import (
	"github.com/rogerioisj/soat-project/internal/core/domain"
	"github.com/rogerioisj/soat-project/internal/core/ports/repositories"
)

type DeleteItemService struct {
	ir repositories.ItemRepositoryInterface
}

func NewDeleteItemService(ir repositories.ItemRepositoryInterface) *DeleteItemService {
	return &DeleteItemService{
		ir: ir,
	}
}

func (s *DeleteItemService) Execute(itemId string) error {
	if itemId == "" {
		return domain.NewDomainError(domain.InvalidItemId, "item ID cannot be empty")
	}

	err := s.ir.Delete(itemId)
	if err != nil {
		return domain.NewDomainError(domain.ItemNotFound, "error deleting item: "+err.Error())
	}

	return nil
}
