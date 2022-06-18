package service

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/bajalnyt/go-mongo-crud/internal/db"
	"go.mongodb.org/mongo-driver/bson"
)

func MongoCrudHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		version := "test"
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		client, err := db.InitDatabase()
		if err != nil {
			fmt.Println("Error writing response, %w", err)
		}
		client.Connect(ctx)
		collection := client.Database("testing").Collection("numbers")
		collection.InsertOne(ctx, bson.D{{"name", "pi"}, {"value", 3.14159}})

		_, err = w.Write([]byte(version))
		if err != nil {
			fmt.Println("Error writing response, %w", err)
		}
	}
}
