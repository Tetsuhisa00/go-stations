package router

import (
	"database/sql"
	"net/http"
   
	"github.com/TechBowl-japan/go-stations/handler"
	"github.com/TechBowl-japan/go-stations/service"
	"github.com/TechBowl-japan/go-stations/handler/middleware"
	
)

func NewRouter(todoDB *sql.DB) *http.ServeMux {
	
	todoService := service.NewTODOService(todoDB)
	todoHandler := handler.NewTODOHandler(todoService)
	panicHandler := handler.NewPanicHandler()
	gracefulHandler := handler.NewGracefulHandler()

	mux := http.NewServeMux()
	mux.Handle("/healthz", middleware.UserAuthentication(&handler.HealthzHandler{}))
	mux.Handle("/todos", todoHandler)
	mux.Handle("/do-panic", middleware.Recovery(panicHandler))
	mux.Handle("/graceful", gracefulHandler)
	return mux
}
