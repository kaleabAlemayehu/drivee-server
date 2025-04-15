package middleware

import (
	"context"
	"net/http"
	"strings"

	jwt "github.com/golang-jwt/jwt/v5"
)

func responseUnauthorized(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")

		// check Bearer prefix pull out token
		tokenStr, found := strings.CutPrefix(authorization, "Bearer ")
		if !found {
			responseUnauthorized(w)
			return
		}
		// decode the token
		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (any, error) {
			return nil, nil
		})

		// TODO: handle other cases like expiration and timing
		if err != nil {
			responseUnauthorized(w)
			return
		}

		if !token.Valid {
			responseUnauthorized(w)
			return
		}
		userID, err := token.Claims.GetIssuer()
		if err != nil {
			responseUnauthorized(w)
			return
		}
		ctx := context.WithValue(r.Context(), "userID", userID)

		// modify the request context to include user id
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
