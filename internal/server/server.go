package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
	"github.com/go-chi/chi/v5"
)

type MyServer struct {
	logger *log.Logger
	server *http.Server
}

func (m *MyServer) ListenAndServe() error {
	err := m.server.ListenAndServe()
	if err != nil {
		m.logger.Fatal(err)
		return err
	}
	return nil
}

func CreateMyServer(lg *log.Logger) *MyServer {
	r := chi.NewRouter()

	r.HandleFunc("/", handlers.Index)
	r.HandleFunc("/upload", handlers.Upload)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ErrorLog:     lg,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &MyServer{logger: lg, server: srv}
}
