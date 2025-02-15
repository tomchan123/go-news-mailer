// Code generated by goa v3.13.1, DO NOT EDIT.
//
// subscription client
//
// Command:
// $ goa gen github.com/tomchan123/go-news-mailer/internal/design --output
// ./internal

package subscription

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "subscription" service client.
type Client struct {
	GetAllEndpoint         goa.Endpoint
	GetOneByUIDEndpoint    goa.Endpoint
	DeleteOneByUIDEndpoint goa.Endpoint
	CreateOneEndpoint      goa.Endpoint
}

// NewClient initializes a "subscription" service client given the endpoints.
func NewClient(getAll, getOneByUID, deleteOneByUID, createOne goa.Endpoint) *Client {
	return &Client{
		GetAllEndpoint:         getAll,
		GetOneByUIDEndpoint:    getOneByUID,
		DeleteOneByUIDEndpoint: deleteOneByUID,
		CreateOneEndpoint:      createOne,
	}
}

// GetAll calls the "getAll" endpoint of the "subscription" service.
// GetAll may return the following errors:
//   - "ServerError" (type *goa.ServiceError): Error returned when there is an internal server error
//   - "SubscriptionNotFound" (type *goa.ServiceError): Error returned when the specified subscription does not exist
//   - "SubscriptionFieldMissing" (type *goa.ServiceError): Error returned when the subscription has missing field(s)
//   - "SubscriptionAlreadyExists" (type *goa.ServiceError): Error returned when the subscription already exists
//   - error: internal error
func (c *Client) GetAll(ctx context.Context) (res []*Subscription, err error) {
	var ires any
	ires, err = c.GetAllEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.([]*Subscription), nil
}

// GetOneByUID calls the "getOneByUID" endpoint of the "subscription" service.
// GetOneByUID may return the following errors:
//   - "ServerError" (type *goa.ServiceError): Error returned when there is an internal server error
//   - "SubscriptionNotFound" (type *goa.ServiceError): Error returned when the specified subscription does not exist
//   - "SubscriptionFieldMissing" (type *goa.ServiceError): Error returned when the subscription has missing field(s)
//   - "SubscriptionAlreadyExists" (type *goa.ServiceError): Error returned when the subscription already exists
//   - error: internal error
func (c *Client) GetOneByUID(ctx context.Context, p string) (res *Subscription, err error) {
	var ires any
	ires, err = c.GetOneByUIDEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Subscription), nil
}

// DeleteOneByUID calls the "deleteOneByUID" endpoint of the "subscription"
// service.
// DeleteOneByUID may return the following errors:
//   - "ServerError" (type *goa.ServiceError): Error returned when there is an internal server error
//   - "SubscriptionNotFound" (type *goa.ServiceError): Error returned when the specified subscription does not exist
//   - "SubscriptionFieldMissing" (type *goa.ServiceError): Error returned when the subscription has missing field(s)
//   - "SubscriptionAlreadyExists" (type *goa.ServiceError): Error returned when the subscription already exists
//   - error: internal error
func (c *Client) DeleteOneByUID(ctx context.Context, p string) (err error) {
	_, err = c.DeleteOneByUIDEndpoint(ctx, p)
	return
}

// CreateOne calls the "createOne" endpoint of the "subscription" service.
// CreateOne may return the following errors:
//   - "ServerError" (type *goa.ServiceError): Error returned when there is an internal server error
//   - "SubscriptionNotFound" (type *goa.ServiceError): Error returned when the specified subscription does not exist
//   - "SubscriptionFieldMissing" (type *goa.ServiceError): Error returned when the subscription has missing field(s)
//   - "SubscriptionAlreadyExists" (type *goa.ServiceError): Error returned when the subscription already exists
//   - error: internal error
func (c *Client) CreateOne(ctx context.Context, p *SubscriptionCreateOnePayload) (res *Subscription, err error) {
	var ires any
	ires, err = c.CreateOneEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Subscription), nil
}
