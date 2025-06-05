package handlers

import (
	"encoding/json"
	"github.com/rogerioisj/soat-project/internal/core/domain"
	"github.com/rogerioisj/soat-project/internal/core/services/user"
	"log"
	"net/http"
)

type GetUserByCpfHandler struct {
	s user.IdentifyWithCpfServiceInterface
}

func NewGetUserByCpfHandler(s user.IdentifyWithCpfServiceInterface) *GetUserByCpfHandler {
	return &GetUserByCpfHandler{
		s: s,
	}
}

func (h *GetUserByCpfHandler) Execute(w http.ResponseWriter, r *http.Request) {
	cpf := r.URL.Query().Get("cpf")

	if cpf == "" {
		http.Error(w, "CPF is required", http.StatusBadRequest)
		return
	}

	if len(cpf) < 11 {
		http.Error(w, "CPF must be at least 11 characters long", http.StatusBadRequest)
		return
	}

	u, err := h.s.Execute(cpf)

	if err != nil && err.Is(domain.UserNotFound) {
		http.Error(w, "", http.StatusNotFound)
		return
	}

	if err != nil {
		log.Print("Error retrieving user by CPF: ", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"ID":    u.GetID(),
		"CPF":   u.GetCPF(),
		"Name":  u.GetName(),
		"Email": u.GetEmail(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
