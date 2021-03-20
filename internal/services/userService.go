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

// UserService ...
type UserService struct {
	db *db.Database
}

// NewUserService ...
func NewUserService(db *db.Database) *UserService {
	return &UserService{
		db: db,
	}
}

// Repo ...
func (u *UserService) Repo() *db.UserRepo {
	return u.db.User()
}

func (u *UserService) hashPassword(password string) string {
	// Helper struct for password hashing
	type passwordConfig struct {
		time    uint32
		memory  uint32
		threads uint8
		keyLen  uint32
	}

	c := &passwordConfig{
		time:    1,
		memory:  64 * 1024,
		threads: 4,
		keyLen:  32,
	}

	hash := argon2.IDKey(
		[]byte(password),
		[]byte(viper.GetString("salt")),
		c.time,
		c.memory,
		c.threads,
		c.keyLen,
	)

	b64Salt := base64.RawStdEncoding.EncodeToString([]byte(viper.GetString("salt")))
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	format := "$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s"
	return fmt.Sprintf(format, argon2.Version, c.memory, c.time, c.threads, b64Salt, b64Hash)
}

// ComparePasswords ...
func (u *UserService) ComparePasswords(user *models.User, pass string) bool {
	return u.hashPassword(pass) == user.Password
}

func (u *UserService) validate(user *models.User) string {
	email := validation.Validate(user.Email, is.Email)
	firstName := validation.Validate(user.FirstName, validation.Length(2, 50))
	lastName := validation.Validate(user.LastName, validation.Length(2, 50))
	password := validation.Validate(user.Password, validation.Length(6, 50))

	errors := ""

	if user.Password == "" || user.FirstName == "" {
		errors += "Must provide all data"
	}

	if email != nil {
		errors += email.Error() + ". Got: " + user.Email + ". "
	}

	if firstName != nil {
		errors += firstName.Error() + ". Got: " + user.FirstName + ". "
	}

	if lastName != nil {
		errors += lastName.Error() + ". Got: " + user.LastName + ". "
	}

	if password != nil {
		errors += password.Error() + ". Got: " + user.Password + ". "
	}

	return errors
}

// Create ...
func (u *UserService) Create(user *models.User) (string, error) {

	// Validation
	if message := u.validate(user); message != "" {
		return "", fmt.Errorf(message)
	}

	user.Password = u.hashPassword(user.Password)

	if err := u.db.User().Create(user); err != nil {
		return "", err
	}

	return u.GenerateToken(user)
}

// GenerateToken ...
func (u *UserService) GenerateToken(user *models.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["email"] = user.Email

	return token.SignedString([]byte(viper.GetString("jwt_secret")))
}

func (u *UserService) GetPortfolio(userID int) (map[string]int, error) {
	operations, err := u.db.Operation().GetAllByUserId(userID)
	if err != nil {
		return nil, err
	}

	processedOperation := make(map[string]int)

	for _, v := range operations {
		if _, ok := processedOperation[v.Symbol]; !ok {
			processedOperation[v.Symbol] = 0
		}

		if v.Type == "buy" {
			processedOperation[v.Symbol] += v.Quantity
		} else {
			processedOperation[v.Symbol] -= v.Quantity
		}
	}

	return processedOperation, nil
}
