// Code generated by goa v3.13.1, DO NOT EDIT.
//
// subscription HTTP server types
//
// Command:
// $ goa gen github.com/tomchan123/go-news-mailer/internal/design --output
// ./internal

package server

import (
	subscription "github.com/tomchan123/go-news-mailer/internal/gen/subscription"
	goa "goa.design/goa/v3/pkg"
)

// CreateOneRequestBody is the type of the "subscription" service "createOne"
// endpoint HTTP request body.
type CreateOneRequestBody struct {
	// Unique identifier of subcription
	UID *string `form:"uid,omitempty" json:"uid,omitempty" xml:"uid,omitempty"`
	// Subscribed email
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// Name of the subscriber
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Datetime when the subscription was made
	Since *string `form:"since,omitempty" json:"since,omitempty" xml:"since,omitempty"`
}

// GetAllResponseBody is the type of the "subscription" service "getAll"
// endpoint HTTP response body.
type GetAllResponseBody []*SubscriptionResponse

// GetOneByUIDResponseBody is the type of the "subscription" service
// "getOneByUID" endpoint HTTP response body.
type GetOneByUIDResponseBody struct {
	// Unique identifier of subcription
	UID *string `form:"uid,omitempty" json:"uid,omitempty" xml:"uid,omitempty"`
	// Subscribed email
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// Name of the subscriber
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Datetime when the subscription was made
	Since *string `form:"since,omitempty" json:"since,omitempty" xml:"since,omitempty"`
}

// CreateOneResponseBody is the type of the "subscription" service "createOne"
// endpoint HTTP response body.
type CreateOneResponseBody struct {
	// Unique identifier of subcription
	UID *string `form:"uid,omitempty" json:"uid,omitempty" xml:"uid,omitempty"`
	// Subscribed email
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// Name of the subscriber
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Datetime when the subscription was made
	Since *string `form:"since,omitempty" json:"since,omitempty" xml:"since,omitempty"`
}

// GetOneByUIDSubscriptionNotFoundResponseBody is the type of the
// "subscription" service "getOneByUID" endpoint HTTP response body for the
// "SubscriptionNotFound" error.
type GetOneByUIDSubscriptionNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// DeleteOneByUIDSubscriptionNotFoundResponseBody is the type of the
// "subscription" service "deleteOneByUID" endpoint HTTP response body for the
// "SubscriptionNotFound" error.
type DeleteOneByUIDSubscriptionNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// CreateOneSubscriptionFieldMissingResponseBody is the type of the
// "subscription" service "createOne" endpoint HTTP response body for the
// "SubscriptionFieldMissing" error.
type CreateOneSubscriptionFieldMissingResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// CreateOneSubscriptionAlreadyExistsResponseBody is the type of the
// "subscription" service "createOne" endpoint HTTP response body for the
// "SubscriptionAlreadyExists" error.
type CreateOneSubscriptionAlreadyExistsResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// SubscriptionResponse is used to define fields on response body types.
type SubscriptionResponse struct {
	// Unique identifier of subcription
	UID *string `form:"uid,omitempty" json:"uid,omitempty" xml:"uid,omitempty"`
	// Subscribed email
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// Name of the subscriber
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Datetime when the subscription was made
	Since *string `form:"since,omitempty" json:"since,omitempty" xml:"since,omitempty"`
}

// NewGetAllResponseBody builds the HTTP response body from the result of the
// "getAll" endpoint of the "subscription" service.
func NewGetAllResponseBody(res []*subscription.Subscription) GetAllResponseBody {
	body := make([]*SubscriptionResponse, len(res))
	for i, val := range res {
		body[i] = marshalSubscriptionSubscriptionToSubscriptionResponse(val)
	}
	return body
}

// NewGetOneByUIDResponseBody builds the HTTP response body from the result of
// the "getOneByUID" endpoint of the "subscription" service.
func NewGetOneByUIDResponseBody(res *subscription.Subscription) *GetOneByUIDResponseBody {
	body := &GetOneByUIDResponseBody{
		UID:   res.UID,
		Email: res.Email,
		Name:  res.Name,
		Since: res.Since,
	}
	return body
}

// NewCreateOneResponseBody builds the HTTP response body from the result of
// the "createOne" endpoint of the "subscription" service.
func NewCreateOneResponseBody(res *subscription.Subscription) *CreateOneResponseBody {
	body := &CreateOneResponseBody{
		UID:   res.UID,
		Email: res.Email,
		Name:  res.Name,
		Since: res.Since,
	}
	return body
}

// NewGetOneByUIDSubscriptionNotFoundResponseBody builds the HTTP response body
// from the result of the "getOneByUID" endpoint of the "subscription" service.
func NewGetOneByUIDSubscriptionNotFoundResponseBody(res *goa.ServiceError) *GetOneByUIDSubscriptionNotFoundResponseBody {
	body := &GetOneByUIDSubscriptionNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewDeleteOneByUIDSubscriptionNotFoundResponseBody builds the HTTP response
// body from the result of the "deleteOneByUID" endpoint of the "subscription"
// service.
func NewDeleteOneByUIDSubscriptionNotFoundResponseBody(res *goa.ServiceError) *DeleteOneByUIDSubscriptionNotFoundResponseBody {
	body := &DeleteOneByUIDSubscriptionNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewCreateOneSubscriptionFieldMissingResponseBody builds the HTTP response
// body from the result of the "createOne" endpoint of the "subscription"
// service.
func NewCreateOneSubscriptionFieldMissingResponseBody(res *goa.ServiceError) *CreateOneSubscriptionFieldMissingResponseBody {
	body := &CreateOneSubscriptionFieldMissingResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewCreateOneSubscriptionAlreadyExistsResponseBody builds the HTTP response
// body from the result of the "createOne" endpoint of the "subscription"
// service.
func NewCreateOneSubscriptionAlreadyExistsResponseBody(res *goa.ServiceError) *CreateOneSubscriptionAlreadyExistsResponseBody {
	body := &CreateOneSubscriptionAlreadyExistsResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewCreateOneSubscriptionCreateOnePayload builds a subscription service
// createOne endpoint payload.
func NewCreateOneSubscriptionCreateOnePayload(body *CreateOneRequestBody) *subscription.SubscriptionCreateOnePayload {
	v := &subscription.SubscriptionCreateOnePayload{
		UID:   body.UID,
		Email: *body.Email,
		Name:  body.Name,
		Since: body.Since,
	}

	return v
}

// ValidateCreateOneRequestBody runs the validations defined on
// CreateOneRequestBody
func ValidateCreateOneRequestBody(body *CreateOneRequestBody) (err error) {
	if body.Email == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("email", "body"))
	}
	return
}
