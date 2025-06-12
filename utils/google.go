package utils

import (
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	jwt "github.com/golang-jwt/jwt/v5"
	"math/big"
	"net/http"
	"sync"
	"time"
)

const (
	googleJWKSURL   = "https://www.googleapis.com/oauth2/v3/certs"
	googleIssuer    = "https://accounts.google.com"
	tokenExpiry     = 5 * time.Minute // Max allowed token age
	jwksRefreshTime = 1 * time.Hour   // How often to refresh Google certs
)

type GoogleClaims struct {
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	jwt.RegisteredClaims
}

type JWK struct {
	Kid string `json:"kid"`
	Kty string `json:"kty"`
	Alg string `json:"alg"`
	Use string `json:"use"`
	N   string `json:"n"`
	E   string `json:"e"`
}

type JWKS struct {
	Keys []JWK `json:"keys"`
}

type GoogleTokenVerifier struct {
	clientID  string
	jwks      *JWKS
	mu        sync.RWMutex
	lastFetch time.Time
}

func NewGoogleTokenVerifier(clientID string) *GoogleTokenVerifier {
	return &GoogleTokenVerifier{
		clientID: clientID,
	}
}

func (v *GoogleTokenVerifier) VerifyToken(tokenString string) (*GoogleClaims, error) {
	// Parse token without verification to get header
	token, _, err := jwt.NewParser().ParseUnverified(tokenString, &GoogleClaims{})
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	// Get key ID from header
	kid, ok := token.Header["kid"].(string)
	if !ok {
		return nil, errors.New("kid not found in token header")
	}

	// Get public key
	publicKey, err := v.getPublicKey(kid)
	if err != nil {
		return nil, fmt.Errorf("failed to get public key: %w", err)
	}

	// Parse and verify token
	claims := &GoogleClaims{}
	token, err = jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})
	if err != nil {
		return nil, fmt.Errorf("token verification failed: %w", err)
	}

	// Validate claims
	if err := v.validateClaims(claims); err != nil {
		return nil, fmt.Errorf("invalid claims: %w", err)
	}

	return claims, nil
}

// validateClaims checks the token claims
func (v *GoogleTokenVerifier) validateClaims(claims *GoogleClaims) error {
	// Validate issuer
	if claims.Issuer != googleIssuer && claims.Issuer != "accounts.google.com" {
		return errors.New("invalid issuer")
	}

	// Validate audience
	// if !claims.Audience.Contains(v.clientID) {
	// 	return errors.New("invalid audience")
	// }

	// Validate expiration
	if claims.ExpiresAt != nil && claims.ExpiresAt.Before(time.Now()) {
		return errors.New("token expired")
	}

	// Validate issued at time (optional but recommended)
	if claims.IssuedAt != nil && time.Since(claims.IssuedAt.Time) > tokenExpiry {
		return errors.New("token too old")
	}

	// Validate email verification
	if !claims.EmailVerified {
		return errors.New("email not verified")
	}

	return nil
}

// getPublicKey retrieves the public key for a given kid
func (v *GoogleTokenVerifier) getPublicKey(kid string) (*rsa.PublicKey, error) {
	v.mu.RLock()
	jwks := v.jwks
	lastFetch := v.lastFetch
	v.mu.RUnlock()

	// Refresh JWKS if needed
	if jwks == nil || time.Since(lastFetch) > jwksRefreshTime {
		newJWKS, err := fetchJWKS()
		if err != nil {
			return nil, err
		}

		v.mu.Lock()
		v.jwks = newJWKS
		v.lastFetch = time.Now()
		jwks = newJWKS
		v.mu.Unlock()
	}

	// Find matching key
	for _, key := range jwks.Keys {
		if key.Kid == kid {
			return jwkToRSAPublicKey(key)
		}
	}

	return nil, errors.New("public key not found for kid")
}

// fetchJWKS retrieves Google's public keys
func fetchJWKS() (*JWKS, error) {
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(googleJWKSURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch JWKS: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}

	var jwks JWKS
	if err := json.NewDecoder(resp.Body).Decode(&jwks); err != nil {
		return nil, fmt.Errorf("failed to decode JWKS: %w", err)
	}

	return &jwks, nil
}

// jwkToRSAPublicKey converts a JWK to an RSA public key
func jwkToRSAPublicKey(jwk JWK) (*rsa.PublicKey, error) {
	if jwk.Kty != "RSA" {
		return nil, errors.New("non-RSA key is not supported")
	}

	// Decode base64url-encoded modulus
	nBytes, err := base64.RawURLEncoding.DecodeString(jwk.N)
	if err != nil {
		return nil, fmt.Errorf("failed to decode modulus: %w", err)
	}
	n := new(big.Int).SetBytes(nBytes)

	// Decode base64url-encoded exponent
	eBytes, err := base64.RawURLEncoding.DecodeString(jwk.E)
	if err != nil {
		return nil, fmt.Errorf("failed to decode exponent: %w", err)
	}
	e := new(big.Int).SetBytes(eBytes).Int64()

	return &rsa.PublicKey{
		N: n,
		E: int(e),
	}, nil
}
