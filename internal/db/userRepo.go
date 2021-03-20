package db

import (
	"database/sql"
	"dstuhack/internal/models"
)

// UserRepo ...
type UserRepo struct {
	db *sql.DB
}

// NewUserRepo ...
func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

// FindByID ...
func (u *UserRepo) FindByID(id int) (*models.User, error) {
	var user models.User

	if err := u.db.QueryRow(`SELECT * FROM users git `, id).Scan(
		&user.ID,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Password,
		&user.Balance,
	); err != nil {
		return nil, err
	}

	return &user, nil
}

// FindByEmail ...
func (u *UserRepo) FindByEmail(email string) (*models.User, error) {
	var user models.User

	if err := u.db.QueryRow(`SELECT * FROM users WHERE email = $1`, email).Scan(
		&user.ID,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Password,
		&user.Balance,
	); err != nil {
		return nil, err
	}

	return &user, nil
}

// Create ...
func (u *UserRepo) Create(user *models.User) error {
	if _, err := u.FindByEmail(user.Email); err != nil && err != sql.ErrNoRows {
		return err
	}
	return u.db.QueryRow(`INSERT INTO users (email, first_name, last_name, password, balance) VALUES ($1, $2, $3, $4, 0) RETURNING id`,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Password,
	).Scan(&user.ID)
}
