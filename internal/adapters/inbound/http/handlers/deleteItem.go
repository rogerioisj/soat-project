package handlers

import (
	"github.com/rogerioisj/soat-project/internal/core/services/item"
	"net/http"
)

type DeleteItemHandler struct {
	s *item.DeleteItemService
}

func NewDeleteItemHandler(s *item.DeleteItemService) *DeleteItemHandler {
	return &DeleteItemHandler{
		s: s,
	}
}

func (h *DeleteItemHandler) Execute(w http.ResponseWriter, r *http.Request) {
	itemId := r.PathValue("id")

	if itemId == "" {
		http.Error(w, "Item ID is required", http.StatusBadRequest)
		return
	}

	err := h.s.Execute(itemId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
