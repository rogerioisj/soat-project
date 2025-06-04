package services

import (
	"github.com/rogerioisj/soat-project/internal/core/domain"
	"github.com/rogerioisj/soat-project/internal/core/ports/repositories"
)

type IdentifyWithCpfServiceInterface interface {
	Execute(cpf string) (*domain.User, error)
}
type IdentifyWithCpfService struct {
	r repositories.UserRepositoryInterface
}

func NewIdentifyWithCpfService(r repositories.UserRepositoryInterface) *IdentifyWithCpfService {
	return &IdentifyWithCpfService{
		r: r,
	}
}

func (s *IdentifyWithCpfService) Execute(cpf string) (*domain.User, *domain.DomainError) {
	u := &domain.User{}

	err := s.r.GetByCpf(u, cpf)

	if err != nil {
		return nil, err
	}
	return u, nil
}
