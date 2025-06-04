package services

import (
	"github.com/rogerioisj/soat-project/internal/core/domain"
	"github.com/rogerioisj/soat-project/internal/core/ports/repositories"
)

type CreateUserServiceInterface interface {
	Execute(name, email, cpf string) (*domain.User, error)
}

type CreateUserService struct {
	r repositories.UserRepositoryInterface
}

func NewCreateUserService(r repositories.UserRepositoryInterface) *CreateUserService {
	return &CreateUserService{
		r: r,
	}
}

func (s *CreateUserService) Execute(name, email, cpf string) (*domain.User, *domain.DomainError) {
	u, err := domain.NewUser("", name, email, cpf)

	if err != nil {
		return nil, err
	}

	err = s.r.GetByCpfOrEmail(u)

	if err != nil && err.Code != domain.UserNotFound {
		return nil, err
	}

	if u.GetID() != "" {
		return nil, domain.NewDomainError(domain.UserAlreadyExists, "User with this CPF or email already exists")
	}

	err = s.r.Create(u)

	if err != nil {
		return nil, err
	}

	return u, nil
}
