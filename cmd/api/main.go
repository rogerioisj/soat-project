package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/rogerioisj/soat-project/internal/adapters/inbound/http/handlers"
	"github.com/rogerioisj/soat-project/internal/adapters/outbound/repositories/postgres"
	"github.com/rogerioisj/soat-project/internal/core/services/user"
	"log"
	"net/http"
)

func main() {
	db := dataBaseConnection()
	defer db.Close()

	userRepository := postgres.NewUserRepository(db)

	identifyWithCpfService := user.NewIdentifyWithCpfService(userRepository)

	uh := handlers.NewGetUserByCpfHandler(identifyWithCpfService)

	http.HandleFunc("GET /user", uh.Execute)

	fmt.Println("Listening on port 8080")
	fmt.Println("GET http://127.0.0.1:8080")

	http.ListenAndServe(":8080", nil)
}

func dataBaseConnection() *sql.DB {
	connStr := "user=admin password=admin dbname=restaurant_db sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
