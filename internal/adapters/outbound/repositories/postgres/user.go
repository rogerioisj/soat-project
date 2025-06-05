package postgres

import (
	"database/sql"
	"errors"
	"github.com/rogerioisj/soat-project/internal/core/domain"
	"log"
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
	row := r.db.QueryRow("INSERT INTO users (name, email, cpf) VALUES ($1, $2, $3) RETURNING id", user.GetName(), user.GetEmail(), user.GetCPF())

	var id string

	if err := row.Scan(&id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.NewDomainError(domain.UserAlreadyExists, "User with this CPF or email already exists")
		}
		log.Print("Error inserting user into database: ", err)
		return domain.NewDomainError("Database Error", "Error inserting user into database")
	}

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

func (r *UserRepository) GetByEmail(user *domain.User, email string) *domain.DomainError {
	row := r.db.QueryRow("SELECT id, name, cpf FROM users WHERE cpf = $1", email)

	var id, name, cpf string
	if err := row.Scan(&id, &name, &cpf); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.NewDomainError(domain.UserNotFound, "User with this email not found")
		}
		return domain.NewDomainError("Database Error", "Error querying user by email")
	}

	user.SetID(id)
	user.SetName(name)
	user.SetEmail(email)
	user.SetCPF(cpf)

	return nil
}

func (r *UserRepository) GetByCpfOrEmail(user *domain.User) *domain.DomainError {
	row := r.db.QueryRow("SELECT id, name, email, cpf FROM users WHERE cpf = $1 OR email = $2", user.GetCPF(), user.GetEmail())
	var id, name, email, cpf string
	if err := row.Scan(&id, &name, &email, &cpf); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.NewDomainError(domain.UserNotFound, "User with this CPF or email not found")
		}
		return domain.NewDomainError("Database Error", "Error querying user by CPF or email")
	}

	user.SetID(id)
	user.SetName(name)

	return nil
}

func (r *UserRepository) GetGuestUser(user *domain.User) *domain.DomainError {
	row := r.db.QueryRow("SELECT id, name, email, cpf FROM users WHERE name = $1", "Convidado")

	var id, name, email, cpf string
	if err := row.Scan(&id, &name, &email, &cpf); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.NewDomainError(domain.UserNotFound, "Guest user not found")
		}
		return domain.NewDomainError("Database Error", "Error querying guest user")
	}

	user.SetID(id)
	user.SetName(name)
	user.SetEmail(email)
	user.SetCPF(cpf)

	return nil
}
