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
