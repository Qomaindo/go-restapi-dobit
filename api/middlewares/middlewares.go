package middlewares

import (
	"errors"
	"net/http"

	"github.com/hdyro/go-restapi-bit/api/auth"
	"github.com/hdyro/go-restapi-bit/api/responses"
)

func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

func SetMiddlewareAuthenticationBibite(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.TokenValidBibite(r)
		if err != nil {
			responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized Bibite"))
			return
		}
		next(w, r)
	}
}

func SetMiddlewareAuthenticationMitra(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.TokenValidMitra(r)
		if err != nil {
			responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized Mitra"))
			return
		}
		next(w, r)
	}
}

func SetMiddlewareAuthenticationCustomer(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.TokenValidCustomer(r)
		if err != nil {
			responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized Customer"))
			return
		}
		next(w, r)
	}
}
