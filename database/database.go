package database

import (
	"authorization/custom_models"
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	client *mongo.Client
}

func Connect() *DB {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mongoUri := os.Getenv("MONGO_DB_URI")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUri))
	if err != nil {
		panic(err)
	}
	db := DB{client: client}
	return &db
}

func (d *DB) Disconnect() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return d.client.Disconnect(ctx)
}

func (d *DB) InsertToken(token *custom_models.Token) (bool, error) {
	coll := d.client.Database("authorization").Collection("session")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := coll.InsertOne(ctx, token)
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

func (d *DB) FindOneAndDeleteToken(token string) (bool, error) {
	coll := d.client.Database("authorization").Collection("session")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.M{"token": token}
	_, err := coll.DeleteOne(ctx, filter)
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

func (d *DB) IsTokenExists(token string) (bool, error) {
	coll := d.client.Database("authorization").Collection("session")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.M{"token": token}
	var res *custom_models.Token
	err := coll.FindOne(ctx, filter).Decode(&res)
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}