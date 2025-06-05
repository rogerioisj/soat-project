package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/rogerioisj/soat-project/internal/adapters/inbound/http/handlers"
	"github.com/rogerioisj/soat-project/internal/adapters/outbound/repositories/postgres"
	"github.com/rogerioisj/soat-project/internal/core/services/item"
	"github.com/rogerioisj/soat-project/internal/core/services/order"
	"github.com/rogerioisj/soat-project/internal/core/services/user"
	"log"
	"net/http"
)

func main() {
	db := dataBaseConnection()
	defer db.Close()

	userRepository := postgres.NewUserRepository(db)
	itemRepository := postgres.NewItemRepository(db)
	orderRepository := postgres.NewOrderRepository(db)

	identifyWithCpfService := user.NewIdentifyWithCpfService(userRepository)
	createUserService := user.NewCreateUserService(userRepository)

	getItensService := item.NewGetItensBasedOnCategoryService(itemRepository)

	createOrderService := order.NewCreateOrderService(orderRepository, userRepository)

	guh := handlers.NewGetUserByCpfHandler(identifyWithCpfService)
	cuh := handlers.NewCreateUser(createUserService)
	gih := handlers.NewGetItensHandler(getItensService)
	coh := handlers.NewCreateOrderHandler(createOrderService)

	http.HandleFunc("GET /user", guh.Execute)
	http.HandleFunc("POST /user", cuh.Execute)

	http.HandleFunc("GET /itens", gih.Execute)

	http.HandleFunc("POST /order", coh.Execute)

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
