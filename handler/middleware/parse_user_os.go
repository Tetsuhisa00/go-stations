package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/mileusna/useragent"
)

type userOSKey string

const uoskey = userOSKey("UserOS")

type RequestInfo struct {
	Timestamp time.Time `json:"timestamp"`
	Latency   int64     `json:"latency"`
	Path      string    `json:"path"`
	OS        string    `json:"os"`
}

func AccessLog(h http.Handler) http.Handler {

	fn := func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()

		ua := useragent.Parse(r.UserAgent())

		ctx := context.WithValue(r.Context(), uoskey, ua.OS)
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)

		latency := time.Since(start).Milliseconds()

		userOS, ok := r.Context().Value(uoskey).(string)
		if !ok {
			userOS = "Unknown"
		}

		reqInfo := &RequestInfo{
			Timestamp: start,
			Latency:   latency,
			Path:      r.URL.Path,
			OS:        userOS,
		}

		reqInfoJson, err := json.Marshal(reqInfo)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println(string(reqInfoJson))
	}

	return http.HandlerFunc(fn)
}
