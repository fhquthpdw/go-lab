package mgo

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// BaseModel ...
type BaseModel struct {
}

var MongoDB *mongo.Database

func init() {
	// init ENV
	_, IsExists := os.LookupEnv("APP_ENV")
	if IsExists == false {
		err := godotenv.Load()
		if err != nil {
			log.Printf("load config error: %s", err.Error())
		}
	}

	// connect to mongo
	mongoHost := os.Getenv("mongo_host")
	mongoPort := os.Getenv("mongo_port")
	mongoUser := os.Getenv("mongo_user")
	mongoPass := os.Getenv("mongo_pass")
	mongoDBName := os.Getenv("mongo_db")

	mongoURI := fmt.Sprintf("mongodb://%s:%s@%s:%s", mongoUser, mongoPass, mongoHost, mongoPort)
	mongoClient, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		fmt.Errorf("create mongodb client error: %s", err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = mongoClient.Connect(ctx)
	if err != nil {
		log.Printf("mongodb client connect error: %s", err.Error())
	}

	MongoDB = mongoClient.Database(mongoDBName)
}