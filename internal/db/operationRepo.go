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
func (u *OpertionRepo) Create(operation *models.Operation) error {
	if _, err := u.db.QueryRow(`INSERT INTO operations (user_id, type, name, purchase_price, amount) VALUES ($1, $2, $3, $4, $5)`,
		&operation.UserId,
		&operation.Type,
		&operation.Name,
		&operation.PurchasePrice,
		&operation.Amount,
	); if err != nil {
		return err
	}
}

// GetAllByUserId ...
func (u *OperationRepo) GetAllByUserId(userId int) ([]*models.Operation, error) {
	rows, err := u.db.Query(`SELECT * FROM operations WHERE user_id = $1`, userId)

	var operations []*models.Operation
	for rows.Next() {
		var (
			id       	  int
			userId   	  int
			_type	 	  string
			name 	 	  string
			purchasePrice float32
			amount		  int
		)

		if err := rows.Scan(&id, 
							&userId, 
							&_type, 
							&name, 
							&purchasePrice, 
							&amount); err != nil {
			return nil, err
		}
		_operation := &models.Operation{
			ID: id
			UserId: userId
			Type: type_
			Name: name
	 		PurchasePrice: purchasePrice
			Amount: amount
		}

		operations = append(operations, _opertaion)
	}
	return &operations, nil
}