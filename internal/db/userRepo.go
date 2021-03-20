package db

import (
	"database/sql"
	"dstuhack/internal/models"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (u *UserRepo) FindByID(id int) (*models.User, error) {
	var user models.User

	if err := u.db.QueryRow(`SELECT * FROM users WHERE id = $1`, id).Scan(
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

func (u *UserRepo) Update(user *models.User) error {
	if _, err := u.FindByEmail(user.Email); err != nil && err != sql.ErrNoRows {
		return err
	}

	return u.db.QueryRow(`UPDATE users SET first_name = $1, last_name = $2, password = $3, balance = $4`,
		&user.FirstName,
		&user.LastName,
		&user.Password,
		&user.Balance,
	).Scan()
}
