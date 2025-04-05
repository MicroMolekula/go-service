package mongo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	defaultTimeout = 10 * time.Second
)

type MongoClient struct {
	client   *mongo.Client
	database *mongo.Database
	timeout  time.Duration
}

func NewMongoClient(uri, dbName string) (*MongoClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	return &MongoClient{
		client:   client,
		database: client.Database(dbName),
		timeout:  defaultTimeout,
	}, nil
}

func (mc *MongoClient) GetCollection(name string) *mongo.Collection {
	return mc.database.Collection(name)
}

func (mc *MongoClient) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), mc.timeout)
	defer cancel()
	return mc.client.Disconnect(ctx)
}
