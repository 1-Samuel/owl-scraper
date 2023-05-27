package main

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongodbRepo struct {
	db *mongo.Database
}

func (m MongodbRepo) Persist(match Match) error {
	filter := bson.D{{"id", match.ID}}
	opts := options.Replace().SetUpsert(true)
	_, err := m.db.Collection("matches").ReplaceOne(context.TODO(), filter, match, opts)
	return err
}

func (m MongodbRepo) PersistActive(activeMatch ActiveMatch) error {
	filter := bson.D{{"uid", activeMatch.UID}}
	opts := options.Replace().SetUpsert(true)
	_, err := m.db.Collection("activeMatches").ReplaceOne(context.TODO(), filter, activeMatch, opts)
	return err
}

func (m MongodbRepo) DeleteActiveMatch(uid string) error {
	filter := bson.D{{"uid", uid}}
	_, err := m.db.Collection("activeMatches").DeleteOne(context.TODO(), filter)
	return err
}

func (m MongodbRepo) Get() ([]Match, error) {
	opts := options.Find().SetSort(bson.D{{"start", 1}})
	cur, err := m.db.Collection("matches").Find(context.TODO(), bson.D{}, opts)
	if err != nil {
		return nil, err
	}
	var results []Match
	if err = cur.All(context.TODO(), &results); err != nil {
		return nil, err
	}
	return results, nil
}

func (m MongodbRepo) GetActive() (*ActiveMatch, error) {
	filter := bson.D{{"uid", uid}}
	var result ActiveMatch
	err := m.db.Collection("activeMatches").FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
