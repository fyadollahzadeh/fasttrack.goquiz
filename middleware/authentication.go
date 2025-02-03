package middleware

import (
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userId := r.Header.Get("userId")

		if userId == "" {
			http.Error(w, "Unauthorized: Missing userId", http.StatusUnauthorized)
			return
		}

		// Proceed if userId exists
		next.ServeHTTP(w, r)
	})
}
