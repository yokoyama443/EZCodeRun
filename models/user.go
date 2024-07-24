package models

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"not null" json:"name"`
	Email     string    `gorm:"unique;not null" json:"email"`
	Password  string    `gorm:"not null" json:"password,omitempty"`
	CreatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"created_at"`
}

// JWT用
type Claims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	jwt.StandardClaims
}

func CreateUser(user *User) error {
	result := DB.Create(user)
	return result.Error
}

func GetUserByEmail(email string) (User, error) {
	var user User
	result := DB.Where("email = ?", email).First(&user)
	return user, result.Error
}

func GetUserByID(id uint) (User, error) {
	var user User
	result := DB.First(&user, id)
	return user, result.Error
}
