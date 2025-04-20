package middleware

import (
	"context"
	"log"
	"net/http"
	"os"
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
			log.Println("bearer not found")
			responseUnauthorized(w)
			return
		}
		// decode the token
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return []byte(os.Getenv("JWT_SECRET")), nil
		}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
		// TODO: handle other cases like expiration and timing
		if err != nil {
			log.Println("error while decoding token")
			log.Println(err.Error())
			responseUnauthorized(w)
			return
		}

		if !token.Valid {
			log.Println("the token is not valid")
			responseUnauthorized(w)
			return
		}
		// userID, err := token.Claims.GetIssuer()
		// log.Printf("userID: %v\n", userID)
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			log.Println("not about to extract token claims")
			responseUnauthorized(w)
			return
		}
		ctx := context.WithValue(r.Context(), "userID", claims["sub"])

		// modify the request context to include user id
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
