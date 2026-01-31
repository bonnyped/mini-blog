package mymiddleware

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth/v5"
)

func Authorize(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := chi.URLParam(r, "id")

		_, claims, err := jwtauth.FromContext(r.Context())
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if claims["user_id"] != userID {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
