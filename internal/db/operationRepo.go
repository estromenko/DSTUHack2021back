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
	if err := u.db.QueryRow(`INSERT INTO operations (user_id, type, name, purchase_price, amount) VALUES ($1, $2, $3, $4, $5) RETURNING id`,
		&operation.UserId,
		&operation.Type,
		&operation.Name,
		&operation.PurchasePrice,
		&operation.Amount,
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
			id            int
			userId        int
			_type         string
			name          string
			purchasePrice float32
			amount        int
		)

<<<<<<< HEAD
		if err := rows.Scan(
			&id,
=======
		if err := rows.Scan(&id,
>>>>>>> main
			&userId,
			&_type,
			&name,
			&purchasePrice,
			&amount,
		); err != nil {
			return nil, err
		}
		_operation := &models.Operation{
			ID:            id,
			UserId:        userId,
			Type:          _type,
			Name:          name,
			PurchasePrice: purchasePrice,
			Amount:        amount,
		}

		operations = append(operations, _operation)
	}
	return operations, nil
}
<<<<<<< HEAD
 
// Ищем все операции одного типа
func (u *OperationRepo) GetOperationByNameAndUserId(name string, userId int) ([]*models.Operation, error) {
	rows, _ := u.db.Query(`SELECT * FROM operations WHERE user_id = $1 AND name = $2`, userId, name)

	var operations []*models.Operation 
	for rows.Next()	{
=======

// Валюта, акция, облигация
func (u *OperationRepo) GetAllUserValute(searchType string, id int) ([]*models.Operation, error) {
	rows, _ := u.db.Query(`SELECET * FROM operations WHERE type = $1 AND user_id = $2`, searchType, id)

	var operations []*models.Operation
	for rows.Next() {
>>>>>>> main
		var (
			id            int
			userId        int
			_type         string
			name          string
			purchasePrice float32
			amount        int
		)

<<<<<<< HEAD
		if err := rows.Scan(
			&id,
=======
		if err := rows.Scan(&id,
>>>>>>> main
			&userId,
			&_type,
			&name,
			&purchasePrice,
			&amount,
		); err != nil {
			return nil, err
		}
<<<<<<< HEAD
		_operation := &models.Operation{
=======
		operation := &models.Operation{
>>>>>>> main
			ID:            id,
			UserId:        userId,
			Type:          _type,
			Name:          name,
			PurchasePrice: purchasePrice,
			Amount:        amount,
		}

<<<<<<< HEAD
		operations = append(operations, _operation)
	}
	return operations, nil
}

func (u *OperationRepo) ChangeOperation(operation *models.Operation, amount int) error {
	var oldOperation models.Operation

	if err := u.db.QueryRow(`SELECT * FROM operations WHERE type = $1 AND userId = $2 AND purchase_price = $3 AND amount = $4`,
		operation.Type, operation.UserId, operation.PurchasePrice, operation.Amount).Scan(
		&oldOperation.ID,
		&oldOperation.UserId,
		&oldOperation.Type,
		&oldOperation.Name,
		&oldOperation.PurchasePrice,
		&oldOperation.Amount,
	); err != nil {
		return err
	}
	
 	var newAmount int = operation.Amount - amount

	if err := u.db.QueryRow(`UPDATE operations SET amount = $1 WHERE type = $2 AND userId = $3 AND purchase_price = $4`, 
		newAmount,
		operation.UserId, 
		operation.PurchasePrice, 
		operation.Amount).Scan(); err != nil {
		
		return err
	}
	
	return nil
}

func (u *OperationRepo) DeleteOperationById(id int) error{
	if err := u.db.QueryRow(`DELETE FROM Salespeople WHERE id = $1`, 
	id); err != nil {
		return nil
	}

	return nil
=======
		operations = append(operations, operation)
	}

	return operations, nil
>>>>>>> main
}
