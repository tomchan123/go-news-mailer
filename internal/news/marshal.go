package news

import (
	"strings"
	"time"

	"github.com/tomchan123/go-news-mailer/internal/db"
	"github.com/tomchan123/go-news-mailer/internal/gen/subscription"
)

func MarshalSubscription(s *db.Subscription) *subscription.Subscription {
	ms := &subscription.Subscription{}

	uid := s.Uid.Hex()
	ms.UID = &uid
	email := strings.Clone(s.Email)
	ms.Email = &email
	name := strings.Clone(s.Name)
	ms.Name = &name
	cAt := s.CreatedAt.Format(time.RFC3339)
	ms.CreatedAt = &cAt

	return ms
}

func MarshalSubscriptions(ss []*db.Subscription) []*subscription.Subscription {
	var mss []*subscription.Subscription
	for _, s := range ss {
		mss = append(mss, MarshalSubscription(s))
	}
	return mss
}
