package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/mileusna/useragent"
)

type userOSKey string 
const uoskey = userOSKey("UserOS")

type RequestInfo struct {
	Timestamp time.Time `json:"timestamp"`
	Latency int64 `json:"latency"`
	Path string `json:"path"`
	OS string `json:"os"`
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

		reqInfoJson, err := json.Marshal(reqInfo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return 
		}
		fmt.Println(string(reqInfoJson))
	}

	return http.HandlerFunc(fn)
}