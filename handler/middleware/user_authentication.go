package middleware

import (
	"net/http"
	"os"
)

func UserAuthentication(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		if username == os.Getenv("BASIC_AUTH_USER_ID") && password == os.Getenv("BASIC_AUTH_PASSWORD") {
			h.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			return 
		}
	}
	return http.HandlerFunc(fn)
}