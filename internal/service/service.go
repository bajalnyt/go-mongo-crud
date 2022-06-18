package service

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bajalnyt/go-mongo-crud/internal/db"
	"github.com/gobuffalo/logger"
	"go.mongodb.org/mongo-driver/mongo"
)

type config struct {
	Version string `conf:"env:VERSION,default:unknown"`
	Env     string `conf:"env:ENV,default:local,short:e"`
	Port    string `conf:"env:PORT,default:8080,short:p"`
}

type Service struct {
	config      config
	server      *http.Server
	mongoClient *mongo.Client
	logger      logger.Logger
}

func New() (*Service, error) {
	// logger, err := logger.NewLogger()
	// if err != nil {
	// 	logger.Fatalf("logger.NewLogger() error: %s\n", err)
	// }

	cfg := config{
		Port: "8088",
	}

	server := http.Server{
		Addr:    ":" + cfg.Port,
		Handler: Mux(),
	}

	dBclient, err := db.InitDatabase()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err = dBclient.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()

	s := Service{
		config: cfg,
		server: &server,
		//logger: logger,
		mongoClient: dBclient,
	}

	return &s, nil
}

func (s *Service) Run() error {
	fmt.Println("starting offer presentation api...")

	shutdown := make(chan os.Signal, 1)

	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	serverErrors := make(chan error, 1)

	go func() {
		fmt.Printf("api started - host: http://localhost%s", s.server.Addr)
		serverErrors <- s.server.ListenAndServe()
	}()

	select {
	case err := <-serverErrors:
		return fmt.Errorf("server error: %w", err)

	case sig := <-shutdown:
		s.logger.Infof("shutdown started - signal: %s", sig)
		defer s.logger.Infof("shutdown complete - signal: %s", sig)

		shutdownTimeout := 5 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)

		defer cancel()

		if err := s.server.Shutdown(ctx); err != nil {
			s.server.Close()
			return fmt.Errorf("could not stop server gracefully: %w", err)
		}
	}

	return nil
}
