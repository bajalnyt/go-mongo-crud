package main

import (
	"log"

	"github.com/bajalnyt/go-mongo-crud/internal/service"
)

func main() {
	s, err := service.New()
	if err != nil {
		log.Fatalf("service.New() error: %s\n", err)
	}

	if err := s.Run(); err != nil {
		log.Fatalf("service.Run() error: %s\n", err)
	}
}
