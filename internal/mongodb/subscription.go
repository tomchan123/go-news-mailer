package db

import (
	"context"
	"fmt"
	"time"

	"github.com/tomchan123/go-news-mailer/internal/gen/subscription"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (m *Mongodb) subscriptions() *mongo.Collection {
	return m.NewsDB().Collection("subscriptions")
}

func (m *Mongodb) GetAllSubscriptions() ([]*subscription.Subscription, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cur, err := m.subscriptions().Find(ctx, bson.D{})
	if err != nil {
		return nil, fmt.Errorf("get all subscriptions failed: %v", err)
	}
	defer cur.Close(ctx)

	var subs []*subscription.Subscription
	for cur.Next(ctx) {
		var res subscription.Subscription
		err := cur.Decode(&res)
		if err != nil {
			return nil, fmt.Errorf("get all subscriptions failed: %v", err)
		}
		subs = append(subs, &res)
	}
	if err := cur.Err(); err != nil {
		return nil, fmt.Errorf("get all subscriptions failed: %v", err)
	}

	return subs, nil
}

func (m *Mongodb) GetOneSubscription(uid string) (*subscription.Subscription, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var res subscription.Subscription
	err := m.subscriptions().FindOne(ctx, bson.D{{"Uid", uid}}).Decode(&res)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("get one subscription failed: %v", err)
	}

	return &res, err
}
