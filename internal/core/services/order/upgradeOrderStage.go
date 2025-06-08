package order

import (
	"github.com/rogerioisj/soat-project/internal/core/domain"
	"github.com/rogerioisj/soat-project/internal/core/ports/repositories"
)

type UpgradeOrderStageService struct {
	or repositories.OrderRepositoryInterface
}

func NewUpgradeOrderStageService(or repositories.OrderRepositoryInterface) *UpgradeOrderStageService {
	return &UpgradeOrderStageService{
		or: or,
	}
}

func (s *UpgradeOrderStageService) Execute(orderId string) *domain.DomainError {
	var o domain.Order

	err := s.or.GetById(orderId, &o)
	if err != nil {
		return err
	}

	stage := o.GetStatus()

	if stage == domain.Done {
		return domain.NewDomainError("Invalid Stage", "Order is already completed")
	}

	if stage == domain.WaitingPayment || stage == domain.Cancelled {
		return domain.NewDomainError("Invalid Stage", "Order is not in a stage that can be upgraded")
	}

	stage = defineNewStage(stage)

	o.SetStatus(stage)

	err = s.or.Update(&o)

	if err != nil {
		return err
	}

	return nil
}

func defineNewStage(stage domain.OrderStatus) domain.OrderStatus {
	switch stage {
	case domain.Building:
		return domain.Received
	case domain.Received:
		return domain.Preparing
	case domain.Preparing:
		return domain.Ready
	case domain.Ready:
		return domain.Done
	default:
		return stage
	}
}
