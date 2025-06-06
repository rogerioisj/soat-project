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
	transaction, err := r.db.Begin()
	if err != nil {
		log.Print("Error starting transaction: ", err)
		return domain.NewDomainError("Database Error", "Error starting transaction")
	}

	var orderStatus string

	row := transaction.QueryRow("SELECT status FROM orders WHERE id = $1", id)
	if err := row.Scan(&orderStatus); err != nil {
		if err == sql.ErrNoRows {
			log.Print("No order found with the given ID")
			return domain.NewDomainError("Order Not Found", "No order found with the given ID")
		}
		log.Print("Error retrieving order: ", err)
		transaction.Rollback()
		return domain.NewDomainError("Database Error", "Error retrieving order")
	}

	if err := transaction.Commit(); err != nil {
		log.Print("Error committing transaction: ", err)
		return domain.NewDomainError("Database Error", "Error committing transaction")
	}

	var nextOrderStage domain.OrderStatus

	switch orderStatus {
	case "received":
		nextOrderStage = domain.Received
	case "preparing":
		nextOrderStage = domain.Preparing
	case "ready":
		nextOrderStage = domain.Ready
	case "done":
		nextOrderStage = domain.Done
	case "cancelled":
		nextOrderStage = domain.Cancelled
	case "building":
		nextOrderStage = domain.Building
	case "waiting_payment":
		nextOrderStage = domain.WaitingPayment
	}

	order.SetStatus(nextOrderStage)
	order.SetId(id)

	return nil
}

func (r *OrderRepository) Update(order *domain.Order) *domain.DomainError {
	transaction, err := r.db.Begin()
	if err != nil {
		log.Print("Error starting transaction: ", err)
		return domain.NewDomainError("Database Error", "Error starting transaction")
	}

	_, err = transaction.Exec("UPDATE orders SET status = $1 WHERE id = $2", order.GetStatus(), order.GetID())
	if err != nil {
		log.Print("Error updating order status: ", err)
		transaction.Rollback()
		return domain.NewDomainError("Database Error", "Error updating order status")
	}

	if err := transaction.Commit(); err != nil {
		log.Print("Error committing transaction: ", err)
		return domain.NewDomainError("Database Error", "Error committing transaction")
	}

	return nil
}

func (r *OrderRepository) ListActives(orders *[]domain.Order, offset, limit int) *domain.DomainError {
	transaction, err := r.db.Begin()
	if err != nil {
		log.Print("Error starting transaction: ", err)
		return domain.NewDomainError("Database Error", "Error starting transaction")
	}

	query := `
			SELECT
				o.id,
				o.status,
				o.created_at,
				u.id as user_id,
				u.name as user_name,
				i.id as item_id,
				i.name as item_name,
				i.price as item_price,
				oi.quantity as item_quantity
			FROM
			    orders o
			LEFT JOIN users u ON o.user_id = u.id
			LEFT JOIN orders_itens oi ON o.id = oi.order_id
			LEFT JOIN itens i ON oi.item_id = i.id
			WHERE 
			    o.status <> 'done' and
				o.status <> 'cancelled' and
				o.status <> 'building' and
				o.status <> 'waiting_payment'
			ORDER BY o.created_at
			LIMIT $1 OFFSET $2`

	rows, err := transaction.Query(query, limit, offset)

	if err != nil {
		log.Print("Error querying orders: ", err)
		transaction.Rollback()
		return domain.NewDomainError("Database Error", "Error querying orders")
	}
	defer rows.Close()

	ordersMap := make(map[string]domain.Order)

	for rows.Next() {
		var order domain.Order
		var user domain.User
		var item domain.Item

		var orderId, orderStatus, userId, userName, itemId, itemName string
		var itemPrice int64
		var itemQuantity int64
		var orderCreatedAt sql.NullTime

		if err := rows.Scan(&orderId, &orderStatus, &orderCreatedAt, &userId, &userName, &itemId, &itemName, &itemPrice, &itemQuantity); err != nil {
			log.Print("Error scanning order row: ", err)
			transaction.Rollback()
			return domain.NewDomainError("Database Error", "Error scanning order row")
		}

		user.SetName(userName)
		user.SetID(userId)

		order.SetUser(user)
		order.SetId(orderId)
		order.SetStatus(domain.OrderStatus(orderStatus))

		if existingOrder, exists := ordersMap[orderId]; exists {
			order = existingOrder
		}

		item.SetID(itemId)
		item.SetName(itemName)
		item.SetPrice(itemPrice)
		item.SetQuantity(itemQuantity)

		err2 := order.AddItem(item)

		if err2 != nil {
			log.Print("Error adding item to order: ", err2)
			transaction.Rollback()
			return domain.NewDomainError("Order Error", "Error adding item to order")
		}

		ordersMap[orderId] = order
	}

	if err := rows.Err(); err != nil {
		log.Print("Error iterating over order rows: ", err)
		transaction.Rollback()
		return domain.NewDomainError("Database Error", "Error iterating over order rows")
	}

	if err := transaction.Commit(); err != nil {
		log.Print("Error committing transaction: ", err)
		return domain.NewDomainError("Database Error", "Error committing transaction")
	}

	for _, order := range ordersMap {
		*orders = append(*orders, order)
	}

	return nil
}
