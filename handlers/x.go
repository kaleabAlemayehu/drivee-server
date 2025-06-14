package handlers

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"net/http"

	"github.com/kaleabAlemayehu/drivee-server/utils"
)

func (h *handler) GeneratePKCE(w http.ResponseWriter, r *http.Request) {
	verifier := make([]byte, 32)
	if _, err := rand.Read(verifier); err != nil {
		http.Error(w, "Failed to generate PKCE", http.StatusInternalServerError)
		return
	}
	encodedVerifier := base64.RawURLEncoding.EncodeToString(verifier)
	challenge := base64.RawURLEncoding.EncodeToString(sha256.New().Sum([]byte(encodedVerifier)))
	utils.SendResponse(w, "success", http.StatusOK, map[string]string{
		"verifier":  encodedVerifier,
		"challenge": challenge,
	})
}
