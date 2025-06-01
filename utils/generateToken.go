package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base32"
	"encoding/hex"
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
