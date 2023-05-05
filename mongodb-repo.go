package main

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongodbRepo struct {
	coll *mongo.Collection
}

func (m MongodbRepo) Persist(match Match) error {
	filter := bson.D{{"id", match.ID}}
	opts := options.Replace().SetUpsert(true)
	_, err := m.coll.ReplaceOne(context.TODO(), filter, match, opts)
	return err
}
