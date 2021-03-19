package db

import (
	"database/sql"
	"dstuhack/internal/models"
)

type OperationRepo struct {
	db *sql.DB
}

func NewOperationRepo (db *sql.DB) *OperationRepo {
	return &OperationRepo{
		db: db,
	}
}

// Find by id
func (u *OperationRepo) FindOperationsById (id int) (*models.Operation) {
	var operation models.Operation

	if err := u.db.QueryRow(`SELECT * FROM operations WHERE user_id = $1`, id).Scan (
		&operation.ID,
		&operation.UserId,
		&operation.Type,
		&operation.Name,
		&operation.PurchasePrice,
		&operation.Amount,
	); err != nil {
		return nil, err
	}

	return &operation, nil
}

// Create operation
func (u *OpertionRepo) Create (operation *models.Operation) error {
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
/*
func (u *OperationRepo) GetAll(userId int) ([]*models.Operation, error) {
	rows, err := u.db.Query(`SELECT * FROM operations WHERE user_id = $1`, userId)
	if err != nil {
		return nil, err
	}

	var users []*models.Operation
	for rows.Next() {
		var (
			id       int
			email    string
			username string
			password string
			isOnline bool
		)

		if err := rows.Scan(&id, &email, &username, &password, &isOnline); err != nil {
			return nil, err
		}
		user := &entities.User{
			ID:       id,
			Email:    email,
			Username: username,
			IsOnline: isOnline,
			Password: password,
		}

		users = append(users, user)
	}
	return users, nil
}

*/