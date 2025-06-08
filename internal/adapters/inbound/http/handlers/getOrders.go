package handlers

import (
	"encoding/json"
	"github.com/rogerioisj/soat-project/internal/core/domain"
	"github.com/rogerioisj/soat-project/internal/core/services/order"
	"net/http"
	"strconv"
)

type GetOrdersStruct struct {
	s *order.ListActiveOrdersService
}

func NewGetOrdersHandler(s *order.ListActiveOrdersService) *GetOrdersStruct {
	return &GetOrdersStruct{
		s: s,
	}
}

func (h *GetOrdersStruct) Execute(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")

	if page == "" || page < "1" {
		page = "1"
	}

	if limit == "" || limit < "1" {
		limit = "10"
	}

	if _, err := strconv.Atoi(page); err != nil {
		http.Error(w, "Invalid page parameter", http.StatusBadRequest)
		return
	}

	if _, err := strconv.Atoi(limit); err != nil {
		http.Error(w, "Invalid limit parameter", http.StatusBadRequest)
		return
	}

	pageInt, _ := strconv.Atoi(page)
	limitInt, _ := strconv.Atoi(limit)

	orders := make([]domain.Order, 0)

	err := h.s.Execute(&orders, pageInt, limitInt)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	jsonOrders := make([]map[string]interface{}, 0)

	for _, order := range orders {
		jsonUser := make(map[string]interface{})
		jsonItens := make([]map[string]interface{}, 0)

		jsonUser["id"] = order.GetUser().GetID()
		jsonUser["name"] = order.GetUser().GetName()

		for _, item := range *order.GetItens() {
			jsonItem := make(map[string]interface{})
			jsonItem["id"] = item.GetID()
			jsonItem["name"] = item.GetName()
			jsonItem["price"] = item.GetPrice()

			jsonItens = append(jsonItens, jsonItem)
		}

		jsonOrder := make(map[string]interface{})

		jsonOrder["id"] = order.GetID()
		jsonOrder["status"] = order.GetStatus()
		jsonOrder["user"] = jsonUser
		jsonOrder["itens"] = jsonItens

		jsonOrders = append(jsonOrders, jsonOrder)
	}

	if len(jsonOrders) == 0 {
		http.Error(w, "No orders found", http.StatusNoContent)
		return
	}

	jsonResponse := map[string]interface{}{
		"orders": jsonOrders,
		"page":   pageInt,
		"limit":  limitInt,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jsonResponse)
}
