package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	Client *mongo.Client
	DBName string
}

func NewMongoInstance(connectionURL string, dbName string) (*MongoInstance, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(connectionURL)
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return &MongoInstance{
		Client: client,
		DBName: dbName,
	}, err
}

func (m *MongoInstance) GetCollection(collection string) *mongo.Collection {
	return m.Client.Database(m.DBName).Collection(collection)
}
