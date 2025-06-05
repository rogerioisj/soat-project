package handlers

import (
	"github.com/rogerioisj/soat-project/internal/core/services/order"
	"log"
	"net/http"
)

type UpgradeOrderStageHandler struct {
	s *order.UpgradeOrderStageService
}

func NewUpgradeOrderStageHandler(s *order.UpgradeOrderStageService) *UpgradeOrderStageHandler {
	return &UpgradeOrderStageHandler{
		s: s,
	}
}

func (h *UpgradeOrderStageHandler) Execute(w http.ResponseWriter, r *http.Request) {
	orderID := r.URL.Query().Get("id")
	if orderID == "" {
		http.Error(w, "order_id is required", http.StatusBadRequest)
		return
	}

	err := h.s.Execute(orderID)
	if err != nil && err.Error() != "No order found with the given ID" && err.Error() != "Order is already completed" {
		log.Print("Error upgrading order stage: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
