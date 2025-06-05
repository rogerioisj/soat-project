package handlers

import (
	"encoding/json"
	"github.com/rogerioisj/soat-project/internal/adapters/inbound/http/dtos"
	"github.com/rogerioisj/soat-project/internal/core/domain"
	"github.com/rogerioisj/soat-project/internal/core/services/user"
	"io"
	"log"
	"net/http"
	"regexp"
)

type CreateUserHandler struct {
	s user.CreateUserServiceInterface
}

func NewCreateUser(s user.CreateUserServiceInterface) *CreateUserHandler {
	return &CreateUserHandler{
		s: s,
	}
}

func (h *CreateUserHandler) Execute(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)

	if err != nil {
		log.Print("Error reading request body: ", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	defer r.Body.Close()

	dto := &dtos.CreateUserRequest{}

	err = json.Unmarshal(requestBody, dto)

	if err != nil {
		log.Print("Error unmarshalling request body: ", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if dto.CPF == "" || dto.Name == "" || dto.Email == "" {
		http.Error(w, "CPF, Name and Email are required", http.StatusBadRequest)
		return
	}

	if len(dto.CPF) < 11 {
		http.Error(w, "CPF must be at least 11 characters long", http.StatusBadRequest)
		return
	}

	if len(dto.Name) < 3 {
		http.Error(w, "Name must be at least 3 characters long", http.StatusBadRequest)
		return
	}

	if len(dto.Email) < 5 {
		http.Error(w, "Email must be at least 5 characters long", http.StatusBadRequest)
		return
	}

	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	regex, _ := regexp.Compile(emailRegex)

	result := regex.Match([]byte(dto.Email))

	if !result {
		http.Error(w, "Email format is invalid", http.StatusBadRequest)
		return
	}

	_, err2 := h.s.Execute(dto.Name, dto.Email, dto.CPF)

	if err2 != nil {
		log.Print("Error creating user: ", err2)
		if err2.Is(domain.UserAlreadyExists) {
			http.Error(w, "User already exists", http.StatusConflict)
			return
		}
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
