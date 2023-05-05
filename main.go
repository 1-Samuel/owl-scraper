package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	opts := options.Client().ApplyURI("mongodb://root:example@localhost:27017")
	// Create a new client and connect to the server
	dbClient, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = dbClient.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	database := dbClient.Database("owl")
	collection := database.Collection("matches")

	repo := MongodbRepo{coll: collection}
	scraper := Scraper{
		repo: repo,
	}

	scraper.start()
}
