package middleware

import (
	"log"
	"net/http"
	"os"
)

func UserAuthentication(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("WWW-Authenticate", `Basic realm="Enter userid and password"`)
		userid, password, ok := r.BasicAuth()
		if !ok || userid == "" || password == "" {
			w.WriteHeader(http.StatusUnauthorized)
			log.Println("Basic認証エラー: 認証情報が提供されていないです。")
			return
		}
		if userid != os.Getenv("BASIC_AUTH_USER_ID") || password != os.Getenv("BASIC_AUTH_PASSWORD") {
			w.WriteHeader(http.StatusUnauthorized)
			log.Println("Basic認証エラー: ユーザーIDまたはパスワードが間違っています。")
		}
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
