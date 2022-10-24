package utils

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDB = InitMongoDb()

func InitMongoDb() *mongo.Database {
	config, err := LoadConfig("./.")
	if err != nil {
		log.Fatal("load config err = ", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.DBSource))
	if err != nil {
		log.Println("Connection MonDB Error:", err)
	}
	return client.Database("im")
}
