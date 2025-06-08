package domain

import (
	"regexp"
	"strings"
)

type User struct {
	id    string
	name  string `validate:"required,min=3,max=100"`
	email string
	cpf   string
}

func NewUser(id, name, email, cpf string) (*User, *DomainError) {
	u := &User{
		id:    id,
		name:  name,
		email: email,
		cpf:   cpf,
	}

	err := u.Validate()

	if err != nil {
		return nil, err
	}

	return u, nil
}

func (u *User) GetID() string {
	return u.id
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) GetEmail() string {
	return u.email
}

func (u *User) GetCPF() string {
	return u.cpf
}

func (u *User) SetID(id string) {
	u.id = id
}

func (u *User) SetName(name string) {
	u.name = name
}

func (u *User) SetEmail(email string) {
	u.email = email
}

func (u *User) SetCPF(cpf string) {
	u.cpf = cpf
}

func (u *User) validateName() *DomainError {
	if len(u.name) < 3 || len(u.name) > 100 {
		return NewDomainError(InvalidNameRange, "Name must be between 3 and 100 characters")
	}

	return nil
}

func (u *User) validateEmail() *DomainError {
	if len(u.email) < 3 || len(u.email) > 100 {
		return NewDomainError(InvalidEmailRange, "Email must be between 3 and 100 characters")
	}

	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched := regexp.MustCompile(emailRegex).MatchString(u.email)
	if !matched {
		return NewDomainError(InvalidEmailFormat, "Email format is invalid")
	}

	return nil
}

func (u *User) validateCPF() *DomainError {
	cpf := strings.ReplaceAll(u.cpf, ".", "")
	cpf = strings.ReplaceAll(cpf, "-", "")

	const cpfRegex = `^\d{11}$`
	matched := regexp.MustCompile(cpfRegex).MatchString(cpf)

	if !matched {
		return NewDomainError(InvalidCPF, "CPF must contain exactly 11 digits")
	}

	return nil
}

func (u *User) Validate() *DomainError {
	err := u.validateName()
	if err != nil {
		return err
	}

	err = u.validateEmail()
	if err != nil {
		return err
	}

	err = u.validateCPF()
	if err != nil {
		return err
	}

	return nil
}
