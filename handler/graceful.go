package handler

import (
	"net/http"
	"time"
)

type GracefulHandler struct{}

// ServeHTTP implements http.Handler.
func (p *GracefulHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	time.Sleep(10 * time.Second)
	w.Write([]byte("test for Graceful Shutdown"))
}

func NewGracefulHandler() *GracefulHandler {
	return &GracefulHandler{}
}



