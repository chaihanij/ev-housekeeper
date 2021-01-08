package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// dbConfig for mongodb
type dbConfig struct {
	host     string
	port     string
	user     string
	password string
}

// DBClient is connection of databse
type DBClient struct {
	ctx    context.Context
	client *mongo.Client
}

// NewDBConnection open connection mongodb return context and pointer of mongodb connection
func (d *dbConfig) NewDBConnection() (*DBClient, error) {
	dbURL := fmt.Sprintf("mongodb://%s:%s@%s:%s", d.user, d.password, d.host, d.port)
	var ctx = context.TODO()
	clientOptions := options.Client().ApplyURI(dbURL)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &DBClient{
		ctx:    ctx,
		client: client,
	}, nil
}
