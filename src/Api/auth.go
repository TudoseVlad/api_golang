package api

import (
	"golang_api/src/data"
	"net/http"
	"strings"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		const prefix = "Basic "
		if !strings.HasPrefix(authHeader, prefix) {
			http.Error(w, "Invalid authorization header format", http.StatusUnauthorized)
			return
		}

		encodedCredentials := strings.TrimPrefix(authHeader, prefix)

		if !data.CheckCredentials(encodedCredentials) {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
