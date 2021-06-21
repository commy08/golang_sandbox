package models

import (
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username    string    `grom:"primaryKey; type:text CHARACTER SET utf8 COLLATE uft8_general_ci" json:"username"`
	Password    string    `gorm:"type:text CHARACTER SET utf8 COLLATE utf8_general_ci" json:"password"`
	Firstname   string    `gorm:"type:text CHARACTER SET utf8 COLLATE utf8_general_ci" json:"firstname"`
	Lastname    string    `gorm:"type:text CHARACTER SET utf8 COLLATE utf8_general_ci" json:"lastname"`
	Age         int       `json:"age"`
	DateOfBirth time.Time `json:"date_of_birth"`
}

var (
	jwtKey = os.Getenv("JWT_KEY")
)

// HashPassword : Hash Password
func (u *User) HashPassword() {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	u.Password = string(bytes)
}

// GenerateToken : Generate Token
func (u *User) GenerateToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": u.Username,
	})

	tokenString, err := token.SignedString(jwtKey)
	return tokenString, err
}
