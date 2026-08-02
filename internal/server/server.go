package server

import (
	"log"
	"net/http"
	"time"

	"myServ/internal/handlers"
)

// Server содержит логгер и HTTP-сервер
type Server struct {
	Logger *log.Logger
	HTTP   *http.Server
}

// NewServer создаёт и настраивает новый сервер
func NewServer(logger *log.Logger) *Server {
	mux := http.NewServeMux()

	// Регистрируем хендлеры
	mux.HandleFunc("/", handlers.IndexHandler)
	mux.HandleFunc("/upload", handlers.UploadHandler(logger))

	httpServer := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{
		Logger: logger,
		HTTP:   httpServer,
	}
}

// Start запускает HTTP-сервер
func (s *Server) Start() error {
	s.Logger.Println("Сервер запущен на http://localhost:8080")
	return s.HTTP.ListenAndServe()
}
