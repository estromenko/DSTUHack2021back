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
	if err := u.db.QueryRow(`INSERT INTO operations (user_id, type, name, price, amount) VALUES ($1, $2, $3, $4, $5) RETURNING id`,
		&operation.UserId,
		&operation.Type,
		&operation.Name,
		&operation.Price,
		&operation.Quantity,
	).Scan(&operation.UserId); err != nil {
		return err
	}
	return nil
}

// GetAllByUserId ...
func (u *OperationRepo) GetAllByUserId(userId int) ([]*models.Operation, error) {
	rows, _ := u.db.Query(`SELECT * FROM operations WHERE user_id = $1`, userId)

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
			ID:     id,
			UserId: userId,
			Type:   _type,
			Name:   name,
			Price:  Price,
			Quantity: amount,
		}

		operations = append(operations, _operation)
	}
	return operations, nil
}

// Валюта, акция, облигация
func (u *OperationRepo) GetAllUserValute(searchType string, id int) ([]*models.Operation, error) {
	rows, _ := u.db.Query(`SELECET * FROM operations WHERE type = $1 AND user_id = $2`, searchType, id)

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
		operation := &models.Operation{
			ID:     id,
			UserId: userId,
			Type:   _type,
			Name:   name,
			Price:  Price,
			Quantity: amount,
		}

		operations = append(operations, operation)
	}

	return operations, nil
}
