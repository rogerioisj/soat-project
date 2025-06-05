package postgres

import (
	"database/sql"
	"errors"
	"github.com/rogerioisj/soat-project/internal/core/domain"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(user *domain.User) *domain.DomainError {
	// Implement the logic to create a user in the PostgreSQL database
	return nil
}

func (r *UserRepository) GetByCpf(user *domain.User, cpf string) *domain.DomainError {
	row := r.db.QueryRow("SELECT id, name, email FROM users WHERE cpf = $1", cpf)

	var id, name, email string
	if err := row.Scan(&id, &name, &email); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.NewDomainError(domain.UserNotFound, "User with this CPF not found")
		}
		return domain.NewDomainError("Database Error", "Error querying user by CPF")
	}

	user.SetID(id)
	user.SetName(name)
	user.SetEmail(email)
	user.SetCPF(cpf)

	return nil
}

func (r *UserRepository) GetByEmail(user *domain.User) *domain.DomainError {
	// Implement the logic to get a user by email from the PostgreSQL database
	return nil
}

func (r *UserRepository) GetByCpfOrEmail(user *domain.User) *domain.DomainError {
	return nil
}

func (r *UserRepository) GetGuestUser(user *domain.User) *domain.DomainError {
	// Implement the logic to get a guest user from the PostgreSQL database
	return nil
}
