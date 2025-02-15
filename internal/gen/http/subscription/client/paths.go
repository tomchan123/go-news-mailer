// Code generated by goa v3.13.1, DO NOT EDIT.
//
// HTTP request path constructors for the subscription service.
//
// Command:
// $ goa gen github.com/tomchan123/go-news-mailer/internal/design --output
// ./internal

package client

import (
	"fmt"
)

// GetAllSubscriptionPath returns the URL path to the subscription service getAll HTTP endpoint.
func GetAllSubscriptionPath() string {
	return "/api/subscriptions"
}

// GetOneByUIDSubscriptionPath returns the URL path to the subscription service getOneByUID HTTP endpoint.
func GetOneByUIDSubscriptionPath(uid string) string {
	return fmt.Sprintf("/api/subscriptions/%v", uid)
}

// DeleteOneByUIDSubscriptionPath returns the URL path to the subscription service deleteOneByUID HTTP endpoint.
func DeleteOneByUIDSubscriptionPath(uid string) string {
	return fmt.Sprintf("/api/subscriptions/%v", uid)
}

// CreateOneSubscriptionPath returns the URL path to the subscription service createOne HTTP endpoint.
func CreateOneSubscriptionPath() string {
	return "/api/subscriptions"
}
