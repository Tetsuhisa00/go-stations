package middleware

import (
	"log"
	"net/http"
	"os"
)

func UserAuthentication(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("WWW-Authenticate", `Basic realm="Enter username and password"`)
		username, password, ok := r.BasicAuth()
		if !ok || username == "" || password == "" {
			w.WriteHeader(http.StatusUnauthorized)
			log.Println("Basic認証エラー: 認証情報が提供されていないです。")
			return
		}

		if username == os.Getenv("BASIC_AUTH_USER_ID") && password == os.Getenv("BASIC_AUTH_PASSWORD") {
			h.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			log.Println("Basic認証エラー: ユーザーIDまたはパスワードが間違っています。")
			return
		}
	}
	return http.HandlerFunc(fn)
}
