package db

import (
	"database/sql"
	"dstuhack/internal/models"
	"time"
)

// Operation repo
type OperationRepo struct {
	db *sql.DB
}

// New opertaion
func NewOperationRepo(db *sql.DB) *OperationRepo {
	return &OperationRepo{
		db: db,
	}
}

// Create operation
func (u *OperationRepo) Create(operation *models.Operation) error {
	if err := u.db.QueryRow(`INSERT INTO operations (user_id, type, symbol, name, price, date, quantity) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`,
		&operation.UserId,
		&operation.Type,
		&operation.Symbol,
		&operation.Name,
		&operation.Price,
		&operation.Date,
		&operation.Quantity,
	).Scan(&operation.ID); err != nil {
		return err
	}
	return nil
}

// GetAllByUserId ...
func (r *OperationRepo) GetAllByUserId(userID int) ([]*models.Operation, error) {
	rows, _ := r.db.Query(`SELECT * FROM operations WHERE user_id = $1`, userID)

	var operations []*models.Operation
	for rows.Next() {
		var (
			id       int
			userId   int
			_type    string
			symbol   string
			name     string
			price    float32
			date     time.Time
			quantity int
		)

		if err := rows.Scan(
			&id,
			&userId,
			&_type,
			&symbol,
			&name,
			&price,
			&date,
			&quantity,
		); err != nil {
			return nil, err
		}
		_operation := &models.Operation{
			ID:       id,
			UserId:   userId,
			Type:     _type,
			Symbol:   symbol,
			Name:     name,
			Date:     date,
			Price:    price,
			Quantity: quantity,
		}

		operations = append(operations, _operation)
	}
	return operations, nil
}
