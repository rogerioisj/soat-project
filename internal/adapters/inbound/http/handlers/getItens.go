package handlers

import (
	"encoding/json"
	"github.com/rogerioisj/soat-project/internal/core/domain"
	"github.com/rogerioisj/soat-project/internal/core/services/item"
	"net/http"
	"strconv"
)

type GetItensHandler struct {
	s *item.GetItensBasedOnCategoryService
}

func NewGetItensHandler(s *item.GetItensBasedOnCategoryService) *GetItensHandler {
	return &GetItensHandler{
		s: s,
	}
}

func (h *GetItensHandler) Execute(w http.ResponseWriter, r *http.Request) {
	category := r.URL.Query().Get("category")
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")

	if page == "" {
		page = "1"
	}

	if limit == "" {
		limit = "10"
	}

	pageInt, err := strconv.Atoi(page)

	if err != nil {
		http.Error(w, "Invalid page number", http.StatusBadRequest)
		return
	}

	limitInt, err := strconv.Atoi(limit)

	if err != nil {
		http.Error(w, "Invalid limit number", http.StatusBadRequest)
		return
	}

	if category != "snack" && category != "drink" && category != "dessert" && category != "accompaniment" {
		http.Error(w, "Invalid category", http.StatusBadRequest)
		return
	}

	var itens []domain.Item

	err2 := h.s.Execute(category, int32(pageInt), int32(limitInt), &itens)

	if err2 != nil {
		http.Error(w, "Error retrieving items: "+err2.Error(), http.StatusInternalServerError)
		return
	}

	if len(itens) == 0 {
		http.Error(w, "No items found for the given category", http.StatusNoContent)
		return
	}

	responseArray := make([]map[string]interface{}, len(itens))

	for i, item := range itens {
		responseArray[i] = map[string]interface{}{
			"ID":          item.GetID(),
			"Name":        item.GetName(),
			"Price":       item.GetPrice(),
			"Description": item.GetDescription(),
		}
	}

	responseBody := map[string]interface{}{
		"items": responseArray,
		"page":  pageInt,
		"limit": limitInt,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(responseBody)
	if err != nil {
		http.Error(w, "Error encoding response: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
