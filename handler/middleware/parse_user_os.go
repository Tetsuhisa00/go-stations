package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/mileusna/useragent"
)

type userOSKey string 
const uoskey = userOSKey("UserOS")

type RequestInfo struct {
	Timestamp time.Time
	Latency int64
	OS string
}

func ParseUserOS(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ua := useragent.Parse(r.UserAgent())
		ctx := context.WithValue(r.Context(), uoskey, ua.OS)
		h.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}