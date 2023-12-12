package middleware

import (
	"context"
	"net/http"
	"time"
	"fmt"

	"github.com/mileusna/useragent"
)

type userOSKey string 
const uoskey = userOSKey("UserOS")

type RequestInfo struct {
	Timestamp time.Time
	Latency int64
	Path string
	OS string
}

func ParseUserOS(h http.Handler) http.Handler {

	fn := func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()
		
		ua := useragent.Parse(r.UserAgent())
		ctx := context.WithValue(r.Context(), uoskey, ua.OS)
		h.ServeHTTP(w, r.WithContext(ctx))

		finish := time.Now()
		latency := finish.Sub(start).Milliseconds()

		reqInfo := &RequestInfo {
			Timestamp: finish,
			Latency: latency,
			Path: r.URL.Path,
			OS: ua.OS,
		}
		fmt.Println(reqInfo)
	}

	return http.HandlerFunc(fn)
}