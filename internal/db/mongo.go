package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongodb struct {
	client *mongo.Client
}

func (m *Mongodb) ConnectDB(uri string) (func() error, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, fmt.Errorf("mongodb connect failed: %v", err)
	}

	m.client = client
	d := func() error {
		if err = client.Disconnect(ctx); err != nil {
			return fmt.Errorf("mongodb disconnect failed: %v", err)
		}

		return nil
	}
	return d, nil
}

func (m *Mongodb) Client() *mongo.Client {
	if m.client == nil {
		log.Fatal("invalid mongodb client instance")
	}

	return m.client
}

func (m *Mongodb) NewsDB() *mongo.Database {
	return m.Client().Database("news")
}
