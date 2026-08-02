package main

import (
	"log"
	"os"

	"myServ/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "[MORSE] ", log.LstdFlags|log.Lshortfile)

	srv := server.NewServer(logger)

	if err := srv.Start(); err != nil {
		logger.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
