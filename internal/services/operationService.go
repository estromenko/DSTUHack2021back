package services

import (
	"dstuhack/internal/db"
	"dstuhack/internal/models"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/argon2"

	"github.com/dgrijalva/jwt-go"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/spf13/viper"
)

// OperationService
type OperationService struct {
	db *db.Database
} 

// New operation service
func NewOperationService(db *db.Database) *OperationRepo {
	return &OperationRepo {
		db: db
	}
}

// Repo
func (u *UserService) Repo() *db.OperationRepo {
	return u.db.OpretaionRepoCreate()
}

// Create 
func (u *UserService) Create(opertaion *models.Opertaion) error {
	_, err := u.db.operationRepo.Create(opertaion)
	if err != nil {
		return err
	} else {
		return nil
	}
}

// GetAllByUserId
func (u *UserService) GetAllByUserId(userId int) ([]*models.Opertaion, error) {
	[]operations, err := u.db.operationRepo.FindOperationsById(userId)

	if err == nil {
		return nil, err
	}

	return operations, nil
}