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
func (s *subscriptionsrvc) GetAll(ctx context.Context) (res []*subscription.Subscription, e error) {
	s.logger.Print("subscription.getAll")

	subs, e := s.db.GetAllSubscriptions()
	if e != nil {
		e = subscription.MakeServerError(fmt.Errorf("server error: %v", e)) // development
		return
	}

	res = MarshalSubscriptions(subs)
	return
}

// Get a subscription by UID
func (s *subscriptionsrvc) GetOneByUID(ctx context.Context, p string) (res *subscription.Subscription, e error) {
	s.logger.Print("subscription.getOneByUID")

	sub, e := s.db.GetOneSubscriptionByUid(p)
	if e != nil {
		e = subscription.MakeServerError(fmt.Errorf("server error: %v", e)) // development
		return
	} else if sub == nil {
		e = subscription.MakeSubscriptionNotFound(fmt.Errorf("subscription not found: %v", p))
		return
	}

	res = MarshalSubscription(sub)
	return
}

// Delete a subscription by UID
func (s *subscriptionsrvc) DeleteOneByUID(ctx context.Context, p string) (err error) {
	s.logger.Print("subscription.deleteOneByUID")

	dc, e := s.db.DeleteOneSubscription(p)
	if err != nil {
		err = subscription.MakeServerError(fmt.Errorf("server error: %v", e)) // development
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

	if s, e := s.db.GetOneSubscriptionByEmail(p.Email); err != nil {
		err = subscription.MakeServerError(fmt.Errorf("server error: %v", e)) // development
		return
	} else if s != nil {
		err = subscription.MakeSubscriptionAlreadyExists(fmt.Errorf("subscription with the same email already exists"))
	}

	cs, e := s.db.CreateOneSubscription(UnmarsharlSubscription(p))
	if e != nil {
		err = subscription.MakeServerError(fmt.Errorf("server error: %v", e)) // development
		return
	}

	res = MarshalSubscription(cs)
	return
}
