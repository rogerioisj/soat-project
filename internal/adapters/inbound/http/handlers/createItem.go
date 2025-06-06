package handlers

import (
	"encoding/json"
	"github.com/rogerioisj/soat-project/internal/adapters/inbound/http/dtos"
	"github.com/rogerioisj/soat-project/internal/core/domain"
	"github.com/rogerioisj/soat-project/internal/core/services/item"
	"io"
	"net/http"
)

type CreateItemHandler struct {
	createItemService *item.CreateItemService
}

func NewCreateItemHandler(createItemService *item.CreateItemService) *CreateItemHandler {
	return &CreateItemHandler{
		createItemService: createItemService,
	}
}

func (h *CreateItemHandler) Execute(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	dto := &dtos.CreateItemDto{}

	err = json.Unmarshal(requestBody, &dto)

	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if dto.Name == "" || dto.Price <= 0 || dto.Category == "" {
		http.Error(w, "Name, Price and Category are required", http.StatusBadRequest)
		return
	}

	item := domain.Item{}

	item.SetName(dto.Name)
	item.SetPrice(dto.Price)
	item.SetProductType(domain.ProductType(dto.Category))
	item.SetDescription(dto.Description)

	err2 := h.createItemService.Execute(&item)

	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
