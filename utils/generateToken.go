package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base32"
	"encoding/hex"
	jwt "github.com/golang-jwt/jwt/v5"
	"os"
	"time"

	"github.com/google/uuid"
)

func GenerateToken() (string, string, error) {
	bytes := make([]byte, 15)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", "", err
	}
	token := base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(bytes)
	hash := sha256.Sum256([]byte(token))
	hashed := hex.EncodeToString(hash[:])
	return token, hashed, nil
}

func GenerateJWT(id uuid.UUID, name string, email string, rememberMe bool) (string, error) {

	var date int
	if rememberMe {
		date = 30
	} else {
		date = 7
	}
	// generate jwt token and attach to response
	tokenStr := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   id,
		"name":  name,
		"email": email,
		"iat":   time.Now().Unix(),
		"exp":   time.Now().AddDate(0, 0, date).Unix(),
	})
	token, err := tokenStr.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}
