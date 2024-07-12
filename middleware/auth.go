package middleware

import (
	"database/sql"
	"errors"
	"log"
	"time"

	"burakforum/database"
	"burakforum/models"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("my_secret_key")

func RegisterUser(user *models.User) error {
	// Kullanıcının adı zaten var mı kontrol et
	var existingUser models.User
	checkQuery := "SELECT id, name, email FROM users WHERE name = ?"
	err := database.DB.QueryRow(checkQuery, user.Name).Scan(&existingUser.ID, &existingUser.Name, &existingUser.Email)

	if err != nil && err != sql.ErrNoRows {
		log.Printf("Error checking for existing user: %v", err)
		return err
	}
	if existingUser.Name != "" {
		return errors.New("user with the same name already exists")
	}

	// Şifreyi şifrele
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	// Yeni kullanıcıyı ekle
	insertQuery := "INSERT INTO users (id, name, email, password) VALUES (?, ?, ?, ?)"
	_, err = database.DB.Exec(insertQuery, user.ID, user.Name, user.Email, user.Password)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		return err
	}
	return nil
}

// login işlemi için yaptım
func AuthenticateUser(name, password string) (string, error) {
	var user models.User
	query := "SELECT id, name, password FROM users WHERE name = ?"
	err := database.DB.QueryRow(query, name).Scan(&user.ID, &user.Name, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("user not found")
		}
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &models.Claims{
		UserID: user.ID,
		Email:  user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) (*models.Claims, error) {
	claims := &models.Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, errors.New("invalid token signature")
		}
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}
