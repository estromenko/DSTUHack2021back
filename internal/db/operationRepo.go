package db

import (
	"database/sql"
	"dstuhack/internal/models"
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
	if err := u.db.QueryRow(`INSERT INTO operations (user_id, type, symbol, price, quantity) VALUES ($1, $2, $3, $4, $5) RETURNING id`,
		&operation.UserId,
		&operation.Type,
		&operation.Symbol,
		&operation.Price,
		&operation.Quantity,
	).Scan(&operation.ID); err != nil {
		return err
	}
	return nil
}

// GetAllByUserId ...
func (r *OperationRepo) GetAllByUserId(userId int) ([]*models.Operation, error) {
	rows, _ := r.db.Query(`SELECT * FROM operations WHERE user_id = $1`, userId)

	var operations []*models.Operation
	for rows.Next() {
		var (
			id     int
			userId int
			_type  string
			name   string
			Price  float32
			amount int
		)

		if err := rows.Scan(&id,
			&userId,
			&_type,
			&name,
			&Price,
			&amount,
		); err != nil {
			return nil, err
		}
		_operation := &models.Operation{
			ID:       id,
			UserId:   userId,
			Type:     _type,
			Symbol:   name,
			Price:    Price,
			Quantity: amount,
		}

		operations = append(operations, _operation)
	}
	return operations, nil
}
