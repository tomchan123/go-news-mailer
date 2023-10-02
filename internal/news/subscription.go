package news

import (
	"context"
	"log"

	subscription "github.com/tomchan123/go-news-mailer/internal/gen/subscription"
)

// subscription service example implementation.
// The example methods log the requests and return zero values.
type subscriptionsrvc struct {
	logger *log.Logger
}

// NewSubscription returns the subscription service implementation.
func NewSubscription(logger *log.Logger) subscription.Service {
	return &subscriptionsrvc{logger}
}

// Get all subscriptions
func (s *subscriptionsrvc) GetAll(ctx context.Context) (res []*subscription.Subscription, err error) {
	s.logger.Print("subscription.getAll")
	return
}

// Get a subscription by UID
func (s *subscriptionsrvc) GetOneByUID(ctx context.Context, p string) (res *subscription.Subscription, err error) {
	res = &subscription.Subscription{}
	s.logger.Print("subscription.getOneByUID")
	return
}

// Delete a subscription by UID
func (s *subscriptionsrvc) DeleteOneByUID(ctx context.Context, p string) (err error) {
	s.logger.Print("subscription.deleteOneByUID")
	return
}

// Create a new subscription
func (s *subscriptionsrvc) CreateOne(context.Context, *subscription.SubscriptionCreateOnePayload) (res *subscription.Subscription, err error) {
	s.logger.Print("subscription.CreateOne")
	return
}
