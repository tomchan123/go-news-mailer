package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Subscription struct {
	Uid       primitive.ObjectID `bson:"_id"`
	Name      string             `bson:"name"`
	Email     string             `bson:"email"`
	CreatedAt time.Time          `bson:"createdAt"`
}

func (m *Mongodb) subscriptions() *mongo.Collection {
	return m.NewsDB().Collection("subscriptions")
}

func (m *Mongodb) GetAllSubscriptions() ([]*Subscription, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cur, err := m.subscriptions().Find(ctx, bson.D{})
	if err != nil {
		return nil, fmt.Errorf("get all subscriptions failed: %v", err)
	}
	defer cur.Close(ctx)

	var subs []*Subscription
	for cur.Next(ctx) {
		var res Subscription
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

func (m *Mongodb) GetOneSubscription(uid string) (*Subscription, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	oid, err := convertObjectId(uid)
	if err != nil {
		return nil, fmt.Errorf("get one subscription failed: %v", err)
	}

	var res Subscription
	err = m.subscriptions().FindOne(ctx, bson.M{"_id": *oid}).Decode(&res)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("get one subscription failed: %v", err)
	}

	return &res, nil
}

// return the number of documents deleted
// return -1 when error
// return 0 when no document deleted
func (m *Mongodb) DeleteOneSubscription(uid string) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	oid, err := convertObjectId(uid)
	if err != nil {
		return -1, fmt.Errorf("delete one subscription failed: %v", err)
	}

	res, err := m.subscriptions().DeleteOne(ctx, bson.M{"_id": *oid})
	if err != nil {
		return -1, fmt.Errorf("delete one subscription failed %v", err)
	}

	return int(res.DeletedCount), nil
}

func (m *Mongodb) CreateOneSubscription(s *Subscription) {

}

// convert id string to ObjectId
func convertObjectId(id string) (*primitive.ObjectID, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("convert object id failed: %v", err)
	}
	return &oid, nil
}
