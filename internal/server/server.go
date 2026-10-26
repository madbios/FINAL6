package server

import (
	"log"
	"net/http"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Server struct {
	Logger *log.Logger
	HTTP   *http.Server
}

func ReturnRouter(logger *log.Logger) *Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/upload", handlers.UploadHandler)
	mux.HandleFunc("/", handlers.RootHandler)
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     logger,
		ReadTimeout:  5,
		WriteTimeout: 10,
		IdleTimeout:  15,
	}
	return &Server{
		Logger: logger,
		HTTP:   srv,
	}
}
