package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func GetUserByCpfHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("Method not allowed:", r.Method)

	cpf := r.URL.Query().Get("cpf")

	fmt.Println("Received request to get user by CPF:", cpf)

	if cpf == "" {
		http.Error(w, "CPF is required", http.StatusBadRequest)
		return
	}

	response := map[string]interface{}{
		"CPF": cpf,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
