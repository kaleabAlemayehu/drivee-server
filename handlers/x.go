package handlers

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"net/http"

	"github.com/bradfitz/gomemcache/memcache"
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
	state := make([]byte, 32)
	rand.Read(state)
	stateStr := base64.URLEncoding.EncodeToString(state)
	h.cache.Set(&memcache.Item{Key: stateStr, Value: []byte(stateStr), Expiration: 600})
	utils.SendResponse(w, "success", http.StatusOK, map[string]string{
		"state":     stateStr,
		"verifier":  encodedVerifier,
		"challenge": challenge,
	})
}
