package database

import (
	"context"
	"time"

	"github.com/2paperstar/movie-api/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

var (
	usersCollection *mongo.Collection
)

func init() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(config.DbUrl))
	if err != nil {
		panic(err)
	}

	database := client.Database("movie-api")
	usersCollection = database.Collection("users")
}
