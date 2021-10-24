package main

import (
	"os"

	server "github.com/brittonhayes/warhammer/internal/service"
	"github.com/joho/godotenv"
)

func init() {
	_ = godotenv.Load()
}

func main() {
	host := "127.0.0.1"
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	s := server.New(host + ":" + port)
	s.Listen()
}
