package middleware

import (
	"net/http"
	"os"
)

func UserAuthentication(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		userid, password, ok := r.BasicAuth()
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		if userid != os.Getenv("BASIC_AUTH_USER_ID") || password != os.Getenv("BASIC_AUTH_PASSWORD") {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
