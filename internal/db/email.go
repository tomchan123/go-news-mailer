package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Email struct {
	Uid              primitive.ObjectID   `bson:"_id,omitempty"`
	SentAt           time.Time            `bson:"sentAt"`
	Articles         []NewsAritcle        `bson:"articles"`
	SubscriptionUIDs []primitive.ObjectID `bson:"subscriptionUids"`
}

func (m *Mongodb) emails() *mongo.Collection {
	return m.NewsDB().Collection("emails")
}

func (m *Mongodb) GetAllEmails() ([]*Email, error) {
	panic("Not Implemented")
}

func (m *Mongodb) getOneEmailByUid(uid primitive.ObjectID) (*Email, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var res Email
	err := m.emails().FindOne(ctx, bson.M{"_id": uid}).Decode(&res)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}

	return &res, err
}

func (m *Mongodb) CreateOneEmail(e *Email) (*Email, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := m.emails().InsertOne(ctx, &e)
	if err != nil {
		return nil, fmt.Errorf("db.Mongodb.CreateOneEmail > failed to insert email: %v", err)
	}

	uid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, fmt.Errorf("db.Mongodb.CreateOneEmail > failed to convert object id: %v", err)
	}

	newe, err := m.getOneEmailByUid(uid)
	if newe == nil {
		return nil, fmt.Errorf("db.Mongodb.CreateOneEmail > failed to retrieve new email: %v", err)
	}

	return newe, nil
}
