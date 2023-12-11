package handler

import (
	"net/http"
)

type PanicHandler struct{}

// ServeHTTP implements http.Handler.
func (p *PanicHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
	panic("Intentional panic!")
}

func NewPanicHandler() *PanicHandler {
	return &PanicHandler{}
}



