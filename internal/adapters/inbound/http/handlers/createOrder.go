package handlers

import (
	"encoding/json"
	"github.com/rogerioisj/soat-project/internal/adapters/inbound/http/dtos"
	"github.com/rogerioisj/soat-project/internal/core/domain"
	"github.com/rogerioisj/soat-project/internal/core/services/order"
	"io"
	"log"
	"net/http"
)

type CreateOrderHandler struct {
	s *order.CreateOrderService
}

func NewCreateOrderHandler(s *order.CreateOrderService) *CreateOrderHandler {
	return &CreateOrderHandler{
		s: s,
	}
}

func (h *CreateOrderHandler) Execute(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)

	if err != nil {
		log.Print("Error reading request body: ", err)
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	dto := &dtos.CreateOrderRequest{}

	err = json.Unmarshal(requestBody, dto)

	if err != nil {
		log.Print("Error unmarshalling request body: ", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if dto.UserID == "" || len(dto.Products) == 0 {
		http.Error(w, "UserID and Products are required", http.StatusBadRequest)
		return
	}

	o, _ := domain.NewOrderWithoutUser()

	err2 := h.s.Execute(dto.UserID, &dto.Products, o)

	if err2 != nil {
		log.Print("Error executing create order service: ", err2)
		http.Error(w, err2.Error(), http.StatusBadRequest)
		return
	}

	responseBody := map[string]interface{}{
		"id": o.GetID(),
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(responseBody)
}
