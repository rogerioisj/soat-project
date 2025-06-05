package postgres

import (
	"database/sql"
	"github.com/rogerioisj/soat-project/internal/core/domain"
	"log"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (r *OrderRepository) Create(order *domain.Order) *domain.DomainError {
	userId := order.GetUser().GetID()

	transaction, err := r.db.Begin()

	if err != nil {
		log.Print("Error starting transaction: ", err)
		return domain.NewDomainError("Database Error", "Error starting transaction")
	}

	row := transaction.QueryRow("INSERT INTO orders (user_id) VALUES ($1) RETURNING id", userId)

	var orderId string

	if err := row.Scan(&orderId); err != nil {
		if err == sql.ErrNoRows {
			log.Print("No rows returned when inserting order")
			return domain.NewDomainError("Order Creation Error", "Failed to create order")
		}
		log.Print("Error inserting order into database: ", err)
		return domain.NewDomainError("Database Error", "Error inserting order into database")
	}

	order.SetId(orderId)

	for _, product := range order.ItemOrder {
		_, err := transaction.Exec("INSERT INTO orders_itens (order_id, item_id, quantity) VALUES ($1, $2, $3)", orderId, product.ItemID, product.Quantity)

		if err != nil {
			log.Print("Error inserting order item into database: ", err)
			transaction.Rollback()
			return domain.NewDomainError("Database Error", "Error inserting order item into database")
		}
	}

	if err := transaction.Commit(); err != nil {
		log.Print("Error committing transaction: ", err)
		return domain.NewDomainError("Database Error", "Error committing transaction")
	}

	return nil
}

func (r *OrderRepository) GetById(id string, order *domain.Order) *domain.DomainError {
	return nil
}

func (r *OrderRepository) Update(order *domain.Order) *domain.DomainError {
	// Implement the update logic here
	return nil
}
