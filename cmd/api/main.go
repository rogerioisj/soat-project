package main

import (
	"database/sql"
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
	createItemService := item.NewCreateItemService(itemRepository)

	createOrderService := order.NewCreateOrderService(orderRepository, userRepository)
	upgradeOrderService := order.NewUpgradeOrderStageService(orderRepository)
	listActiveOrdersService := order.NewListOrdersService(orderRepository)

	guh := handlers.NewGetUserByCpfHandler(identifyWithCpfService)
	cuh := handlers.NewCreateUser(createUserService)

	gih := handlers.NewGetItensHandler(getItensService)
	cih := handlers.NewCreateItemHandler(createItemService)

	coh := handlers.NewCreateOrderHandler(createOrderService)
	uoh := handlers.NewUpgradeOrderStageHandler(upgradeOrderService)
	gaoh := handlers.NewGetOrdersHandler(listActiveOrdersService)

	http.HandleFunc("GET /user", guh.Execute)
	http.HandleFunc("POST /user", cuh.Execute)

	http.HandleFunc("POST /item", cih.Execute)
	http.HandleFunc("GET /itens", gih.Execute)

	http.HandleFunc("POST /order", coh.Execute)
	http.HandleFunc("PATCH /order", uoh.Execute)
	http.HandleFunc("GET /orders", gaoh.Execute)

	log.Println("Listening on port 8080")
	log.Println("GET http://127.0.0.1:8080")

	http.ListenAndServe(":8080", nil)
}

func dataBaseConnection() *sql.DB {
	log.Println("Connecting to database ...")
	connStr := "user=admin password=admin dbname=restaurant_db sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Failed to connect to the database:", err)
	} else {
		log.Println("Database connection established successfully.")
	}

	return db
}
