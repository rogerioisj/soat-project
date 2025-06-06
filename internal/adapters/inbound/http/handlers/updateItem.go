package handlers

import (
	"encoding/json"
	"github.com/rogerioisj/soat-project/internal/adapters/inbound/http/dtos"
	"github.com/rogerioisj/soat-project/internal/core/domain"
	"github.com/rogerioisj/soat-project/internal/core/services/item"
	"io"
	"net/http"
	"strconv"
)

type UpdateItemHandler struct {
	updateItemService *item.UpdateItemService
}

func NewUpdateItemHandler(updateItemService *item.UpdateItemService) *UpdateItemHandler {
	return &UpdateItemHandler{
		updateItemService: updateItemService,
	}
}

func (h *UpdateItemHandler) Execute(w http.ResponseWriter, r *http.Request) {
	itemId := r.PathValue("id")

	if itemId == "" {
		http.Error(w, "Item ID is required", http.StatusBadRequest)
		return
	}

	requestBody, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	dto := &dtos.UpdateItemDto{}

	err = json.Unmarshal(requestBody, &dto)

	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if dto.Name == "" || dto.Price <= 0 || dto.Category == "" {
		http.Error(w, "Name, Price and Category are required", http.StatusBadRequest)
		return
	}

	item := &domain.Item{}

	item.SetName(dto.Name)
	item.SetPrice(dto.Price)
	item.SetProductType(domain.ProductType(dto.Category))
	item.SetDescription(dto.Description)

	err2 := h.updateItemService.Execute(item, itemId)

	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusBadRequest)
		return
	}

	jsonResponse := map[string]string{
		"id":          itemId,
		"name":        item.GetName(),
		"description": item.GetDescription(),
		"price":       strconv.FormatInt(item.GetPrice(), 10),
		"category":    string(item.GetProductType()),
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(jsonResponse)
}
