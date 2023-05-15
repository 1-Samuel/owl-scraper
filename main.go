package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}

	opts := options.Client().ApplyURI(uri)
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

	repo := MongodbRepo{db: database}
	scraper := Scraper{
		repo: repo,
	}

	activeMatchTicker := time.NewTicker(time.Minute)
	dailyTicker := time.NewTicker(time.Hour * 24)
	fmt.Println("Started!")

	go func() {
		for {
			select {
			case <-activeMatchTicker.C:
				if scraper.isMatchActive() {
					scraper.activeMatch()
				}
			}
		}
	}()

	go func() {
		for {
			select {
			case <-dailyTicker.C:
				scraper.start()
			}
		}
	}()

	go scraper.start()

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done // Will block here until user hits ctrl+c
}
