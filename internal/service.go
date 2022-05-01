package service

import (
	"net/http"

	"github.com/gobuffalo/logger"
)

type config struct {
	Version string `conf:"env:VERSION,default:unknown"`
	Env     string `conf:"env:ENV,default:local,short:e"`
	Port    string `conf:"env:PORT,default:8080,short:p"`
}

type Service struct {
	config config
	server *http.Server
	logger logger.Logger
}

func New() (*Service, error) {
	logger, err := logger.NewLogger()
	if err != nil {
		logger.Fatalf("logger.NewLogger() error: %s\n", err)
	}

	cfg := config{}

	server := http.Server{
		Addr:    ":" + cfg.Port,
		Handler: Mux(logger),
	}

	s := Service{
		config: cfg,
		server: &server,
		logger: logger,
	}

	return &s, nil
}
