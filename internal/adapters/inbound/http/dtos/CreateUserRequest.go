package dtos

type CreateUserRequest struct {
	Name  string `json:"name" validate:"required,min=3,max=100"`
	Email string `json:"email" validate:"required,email"`
	CPF   string `json:"cpf" validate:"required,len=11,cpf"`
}
