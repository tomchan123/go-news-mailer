package news

import (
	"context"
	"fmt"
	"log"

	db "github.com/tomchan123/go-news-mailer/internal/db"
	subscription "github.com/tomchan123/go-news-mailer/internal/gen/subscription"
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
	s.logger.Print("subscription.getAll")

	subs, err := s.db.GetAllSubscriptions()
	if err != nil {
		err = subscription.MakeServerError(fmt.Errorf("server error: %v", err)) // development
		return
	}

	res = MarshalSubscriptions(subs)
	return
}

// Get a subscription by UID
func (s *subscriptionsrvc) GetOneByUID(ctx context.Context, p string) (res *subscription.Subscription, err error) {
	s.logger.Print("subscription.getOneByUID")

	sub, err := s.db.GetOneSubscription(p)
	if err != nil {
		err = subscription.MakeServerError(fmt.Errorf("server error: %v", err)) // development
		return
	} else if sub == nil {
		err = subscription.MakeSubscriptionNotFound(fmt.Errorf("subscription not found: %v", p))
		return
	}

	res = MarshalSubscription(sub)

	return
}

// Delete a subscription by UID
func (s *subscriptionsrvc) DeleteOneByUID(ctx context.Context, p string) (err error) {
	s.logger.Print("subscription.deleteOneByUID")

	dc, err := s.db.DeleteOneSubscription(p)
	if err != nil {
		err = subscription.MakeServerError(fmt.Errorf("server error: %v", err)) // development
		return
	}
	if dc <= 0 {
		err = subscription.MakeSubscriptionNotFound(fmt.Errorf("subscription not found: %v", p))
		return
	}

	return
}

// Create a new subscription
func (s *subscriptionsrvc) CreateOne(ctx context.Context, p *subscription.SubscriptionCreateOnePayload) (res *subscription.Subscription, err error) {
	s.logger.Print("subscription.CreateOne")
	return
}
