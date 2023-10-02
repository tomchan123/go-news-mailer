package news

import (
	"context"
	"fmt"
	"log"

	subscription "github.com/tomchan123/go-news-mailer/internal/gen/subscription"
	db "github.com/tomchan123/go-news-mailer/internal/mongodb"
)

// subscription service example implementation.
// The example methods log the requests and return zero values.
type subscriptionsrvc struct {
	logger *log.Logger
	db     *db.Mongodb
}

// NewSubscription returns the subscription service implementation.
func NewSubscription(logger *log.Logger, db *db.Mongodb) subscription.Service {
	return &subscriptionsrvc{logger, db}
}

// Get all subscriptions
func (s *subscriptionsrvc) GetAll(ctx context.Context) (res []*subscription.Subscription, err error) {
	res, err = s.db.GetAllSubscriptions()
	if err != nil {
		err = subscription.MakeServerError(fmt.Errorf("server error: %v", err)) // development
		return
	}

	s.logger.Print("subscription.getAll")
	return
}

// Get a subscription by UID
func (s *subscriptionsrvc) GetOneByUID(ctx context.Context, p string) (res *subscription.Subscription, err error) {
	res, err = s.db.GetOneSubscription(p)
	if err != nil {
		err = subscription.MakeServerError(fmt.Errorf("server error: %v", err)) // development
		return
	} else if res == nil {
		err = subscription.MakeSubscriptionNotFound(fmt.Errorf("subscription not found: %v", p))
		return
	}

	s.logger.Print("subscription.getOneByUID")
	return
}

// Delete a subscription by UID
func (s *subscriptionsrvc) DeleteOneByUID(ctx context.Context, p string) (err error) {
	s.logger.Print("subscription.deleteOneByUID")
	return
}

// Create a new subscription
func (s *subscriptionsrvc) CreateOne(ctx context.Context, p *subscription.SubscriptionCreateOnePayload) (res *subscription.Subscription, err error) {
	s.logger.Print("subscription.CreateOne")
	return
}
