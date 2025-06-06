package postgres

import (
	"database/sql"
	"github.com/rogerioisj/soat-project/internal/core/domain"
	"log"
)

type ItemRepository struct {
	db *sql.DB
}

func NewItemRepository(db *sql.DB) *ItemRepository {
	return &ItemRepository{
		db: db,
	}
}

func (r *ItemRepository) Create(item *domain.Item) *domain.DomainError {
	transaction, err := r.db.Begin()
	if err != nil {
		log.Printf("Error starting transaction: %v", err)
		return domain.NewDomainError("Database Error", "Error starting transaction")
	}

	_, err = transaction.Exec("INSERT INTO itens (name, description, price, type) VALUES ($1, $2, $3, $4)",
		item.GetName(),
		item.GetDescription(),
		item.GetPrice(),
		item.GetProductType(),
	)
	if err != nil {
		log.Printf("Error inserting item: %v", err)
		if rollbackErr := transaction.Rollback(); rollbackErr != nil {
			log.Printf("Error rolling back transaction: %v", rollbackErr)
		}
		return domain.NewDomainError("Database Error", "Error inserting item")
	}

	if err := transaction.Commit(); err != nil {
		log.Printf("Error committing transaction: %v", err)
		return domain.NewDomainError("Database Error", "Error committing transaction")
	}

	log.Printf("Item %s created successfully", item.GetName())

	return nil
}

func (r *ItemRepository) GetById(id string, item *domain.Item) *domain.DomainError {
	return nil
}

func (r *ItemRepository) ListByType(productType string, page, limit int32, itemList *[]domain.Item) *domain.DomainError {
	rows, err := r.db.Query("SELECT id, name, description, price FROM itens WHERE type = $1 LIMIT $2 OFFSET $3", productType, limit, (page-1)*limit)

	if err != nil {
		log.Printf("Error querying items by type: %v", err)
		return domain.NewDomainError("Database Error", "Error querying items by type")
	}
	defer rows.Close()

	for rows.Next() {
		var item domain.Item
		var id, name, description string
		var price int
		if err := rows.Scan(&id, &name, &description, &price); err != nil {
			log.Printf("Error scanning item: %v", err)
			return domain.NewDomainError("Database Error", "Error scanning item")
		}

		item.SetID(id)
		item.SetName(name)
		item.SetDescription(description)
		item.SetPrice(int64(price))

		*itemList = append(*itemList, item)
	}

	return nil
}
