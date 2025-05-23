package repository

import (
    "context"
    "log"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
    Client     *mongo.Client
    Collection *mongo.Collection
}

func NewMongoDB(uri, dbName, collectionName string) (*MongoDB, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    clientOptions := options.Client().ApplyURI(uri)
    client, err := mongo.Connect(ctx, clientOptions)
    if err != nil {
        return nil, err
    }

    err = client.Ping(ctx, nil)
    if err != nil {
        return nil, err
    }

    log.Println("MongoDB connected")

    collection := client.Database(dbName).Collection(collectionName)
    return &MongoDB{Client: client, Collection: collection}, nil
}

