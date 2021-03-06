package db

import (
	"context"
	"fmt"
	"github.com/semyon-dev/gpn-tc-backend/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

var (
	client *mongo.Client
	db     *mongo.Database
)

func Connect() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(config.MongoUrl))
	if err != nil {
		log.Fatal(err)
	}

	// Create connect
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println("client MongoDB:", err)
	} else {
		fmt.Println("✔ Подключение client MongoDB успешно")
	}

	db = client.Database("main")

	err = Ping()
	if err == nil {
		fmt.Println("Connected to MongoDB!")
		return
	}
	fmt.Println(err.Error())
}

func Ping() error {
	return client.Ping(context.Background(), readpref.Primary())
}
