package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/rogerioisj/soat-project/config"
	"github.com/rogerioisj/soat-project/internal/adapters/inbound/http/handlers"
	"github.com/rogerioisj/soat-project/internal/adapters/outbound/repositories/postgres"
	"github.com/rogerioisj/soat-project/internal/core/services/item"
	"github.com/rogerioisj/soat-project/internal/core/services/order"
	"github.com/rogerioisj/soat-project/internal/core/services/user"
	"log"
	"net/http"
)

func main() {
	config, err := config.Load()

	if err != nil {
		log.Fatal("Failed to load configuration: ", err)
	}

	db := dataBaseConnection(&config)
	defer db.Close()

	userRepository := postgres.NewUserRepository(db)
	itemRepository := postgres.NewItemRepository(db)
	orderRepository := postgres.NewOrderRepository(db)

	identifyWithCpfService := user.NewIdentifyWithCpfService(userRepository)
	createUserService := user.NewCreateUserService(userRepository)

	getItensService := item.NewGetItensBasedOnCategoryService(itemRepository)
	createItemService := item.NewCreateItemService(itemRepository)
	updateItemService := item.NewUpdateItemService(itemRepository)
	deleteItemService := item.NewDeleteItemService(itemRepository)

	createOrderService := order.NewCreateOrderService(orderRepository, userRepository)
	upgradeOrderService := order.NewUpgradeOrderStageService(orderRepository)
	listActiveOrdersService := order.NewListOrdersService(orderRepository)

	guh := handlers.NewGetUserByCpfHandler(identifyWithCpfService)
	cuh := handlers.NewCreateUser(createUserService)

	gih := handlers.NewGetItensHandler(getItensService)
	cih := handlers.NewCreateItemHandler(createItemService)
	uih := handlers.NewUpdateItemHandler(updateItemService)
	dih := handlers.NewDeleteItemHandler(deleteItemService)

	coh := handlers.NewCreateOrderHandler(createOrderService)
	uoh := handlers.NewUpgradeOrderStageHandler(upgradeOrderService)
	gaoh := handlers.NewGetOrdersHandler(listActiveOrdersService)

	ssdh := handlers.NewShowSwaggerDocHandler()

	prefix := "/api/v1"

	http.HandleFunc("GET "+prefix+"/user/{cpf}", guh.Execute)
	http.HandleFunc("POST "+prefix+"/user", cuh.Execute)

	http.HandleFunc("POST "+prefix+"/item", cih.Execute)
	http.HandleFunc("GET "+prefix+"/itens", gih.Execute)
	http.HandleFunc("PUT "+prefix+"/item/{id}", uih.Execute)
	http.HandleFunc("DELETE "+prefix+"/item/{id}", dih.Execute)

	http.HandleFunc("POST "+prefix+"/order", coh.Execute)
	http.HandleFunc("PATCH "+prefix+"/order/{id}", uoh.Execute)
	http.HandleFunc("GET "+prefix+"/orders", gaoh.Execute)

	http.HandleFunc("GET "+prefix+"/docs/openapi.yaml", ssdh.File)
	http.HandleFunc("GET /", ssdh.ServeBasicSwaggerUI)

	log.Println("Listening on port " + config.Port + " ...")
	log.Println("Read docs " + config.Host + ":" + config.Port + "/")

	http.ListenAndServe(":"+string(config.Port), nil)
}

func dataBaseConnection(config *config.Configuration) *sql.DB {
	log.Println("Connecting to database ...")

	log.Println("Database URL:", config.DatabaseUrl)

	db, err := sql.Open("postgres", config.DatabaseUrl)
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
