// Code generated by goa v3.13.1, DO NOT EDIT.
//
// email HTTP client types
//
// Command:
// $ goa gen github.com/tomchan123/go-news-mailer/internal/design --output
// ./internal

package client

import (
	email "github.com/tomchan123/go-news-mailer/internal/gen/email"
	goa "goa.design/goa/v3/pkg"
)

// SendOneEmailRequestBody is the type of the "email" service "sendOneEmail"
// endpoint HTTP request body.
type SendOneEmailRequestBody struct {
	// Unique identifier of subcription
	UID *string `form:"uid,omitempty" json:"uid,omitempty" xml:"uid,omitempty"`
	// Subscribed email
	Email string `form:"email" json:"email" xml:"email"`
	// Name of the subscriber
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Datetime when the subscription was made
	CreatedAt *string `form:"createdAt,omitempty" json:"createdAt,omitempty" xml:"createdAt,omitempty"`
}

// SendOneEmailSubscriptionFieldMissingResponseBody is the type of the "email"
// service "sendOneEmail" endpoint HTTP response body for the
// "SubscriptionFieldMissing" error.
type SendOneEmailSubscriptionFieldMissingResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// SendOneEmailServerErrorResponseBody is the type of the "email" service
// "sendOneEmail" endpoint HTTP response body for the "ServerError" error.
type SendOneEmailServerErrorResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// NewSendOneEmailRequestBody builds the HTTP request body from the payload of
// the "sendOneEmail" endpoint of the "email" service.
func NewSendOneEmailRequestBody(p *email.SubscriptionSendOneEmailPayload) *SendOneEmailRequestBody {
	body := &SendOneEmailRequestBody{
		UID:       p.UID,
		Email:     p.Email,
		Name:      p.Name,
		CreatedAt: p.CreatedAt,
	}
	return body
}

// NewSendOneEmailSubscriptionFieldMissing builds a email service sendOneEmail
// endpoint SubscriptionFieldMissing error.
func NewSendOneEmailSubscriptionFieldMissing(body *SendOneEmailSubscriptionFieldMissingResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewSendOneEmailServerError builds a email service sendOneEmail endpoint
// ServerError error.
func NewSendOneEmailServerError(body *SendOneEmailServerErrorResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// ValidateSendOneEmailSubscriptionFieldMissingResponseBody runs the
// validations defined on sendOneEmail_SubscriptionFieldMissing_response_body
func ValidateSendOneEmailSubscriptionFieldMissingResponseBody(body *SendOneEmailSubscriptionFieldMissingResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateSendOneEmailServerErrorResponseBody runs the validations defined on
// sendOneEmail_ServerError_response_body
func ValidateSendOneEmailServerErrorResponseBody(body *SendOneEmailServerErrorResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}
