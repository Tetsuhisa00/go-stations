package middleware

import (
	"context"
	"net/http"
	"github.com/mileusna/useragent"
)

type UserOSKey string

func ParseUserOS(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ua := useragent.Parse(r.UserAgent())
		ctx := context.WithValue(r.Context(), UserOSKey("UserOS"), ua.OS)
		h.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}