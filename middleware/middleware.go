package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

type FuncMiddleware func(http.HandlerFunc) http.HandlerFunc

func CreateStack(xs ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for i := len(xs) - 1; i >= 0; i-- {
			x := xs[i]
			next = x(next)
		}
		return next
	}
}

func CreateStackFunc(xs ...Middleware) FuncMiddleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		var h http.Handler = next
		for i := len(xs) - 1; i >= 0; i-- {
			h = xs[i](h)
		}
		return func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
}
