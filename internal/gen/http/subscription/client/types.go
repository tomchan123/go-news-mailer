// Code generated by goa v3.13.1, DO NOT EDIT.
//
// subscription HTTP client types
//
// Command:
// $ goa gen github.com/tomchan123/go-news-mailer/internal/design --output
// ./internal

package client

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
	Email string `form:"email" json:"email" xml:"email"`
	// Name of the subscriber
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Datetime when the subscription was made
	CreatedAt *string `form:"createdAt,omitempty" json:"createdAt,omitempty" xml:"createdAt,omitempty"`
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
	CreatedAt *string `form:"createdAt,omitempty" json:"createdAt,omitempty" xml:"createdAt,omitempty"`
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
	CreatedAt *string `form:"createdAt,omitempty" json:"createdAt,omitempty" xml:"createdAt,omitempty"`
}

// GetAllServerErrorResponseBody is the type of the "subscription" service
// "getAll" endpoint HTTP response body for the "ServerError" error.
type GetAllServerErrorResponseBody struct {
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

// GetOneByUIDSubscriptionNotFoundResponseBody is the type of the
// "subscription" service "getOneByUID" endpoint HTTP response body for the
// "SubscriptionNotFound" error.
type GetOneByUIDSubscriptionNotFoundResponseBody struct {
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

// GetOneByUIDServerErrorResponseBody is the type of the "subscription" service
// "getOneByUID" endpoint HTTP response body for the "ServerError" error.
type GetOneByUIDServerErrorResponseBody struct {
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

// DeleteOneByUIDSubscriptionNotFoundResponseBody is the type of the
// "subscription" service "deleteOneByUID" endpoint HTTP response body for the
// "SubscriptionNotFound" error.
type DeleteOneByUIDSubscriptionNotFoundResponseBody struct {
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

// DeleteOneByUIDServerErrorResponseBody is the type of the "subscription"
// service "deleteOneByUID" endpoint HTTP response body for the "ServerError"
// error.
type DeleteOneByUIDServerErrorResponseBody struct {
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

// CreateOneSubscriptionFieldMissingResponseBody is the type of the
// "subscription" service "createOne" endpoint HTTP response body for the
// "SubscriptionFieldMissing" error.
type CreateOneSubscriptionFieldMissingResponseBody struct {
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

// CreateOneSubscriptionAlreadyExistsResponseBody is the type of the
// "subscription" service "createOne" endpoint HTTP response body for the
// "SubscriptionAlreadyExists" error.
type CreateOneSubscriptionAlreadyExistsResponseBody struct {
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

// CreateOneServerErrorResponseBody is the type of the "subscription" service
// "createOne" endpoint HTTP response body for the "ServerError" error.
type CreateOneServerErrorResponseBody struct {
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

// SubscriptionResponse is used to define fields on response body types.
type SubscriptionResponse struct {
	// Unique identifier of subcription
	UID *string `form:"uid,omitempty" json:"uid,omitempty" xml:"uid,omitempty"`
	// Subscribed email
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// Name of the subscriber
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Datetime when the subscription was made
	CreatedAt *string `form:"createdAt,omitempty" json:"createdAt,omitempty" xml:"createdAt,omitempty"`
}

// NewCreateOneRequestBody builds the HTTP request body from the payload of the
// "createOne" endpoint of the "subscription" service.
func NewCreateOneRequestBody(p *subscription.SubscriptionCreateOnePayload) *CreateOneRequestBody {
	body := &CreateOneRequestBody{
		UID:       p.UID,
		Email:     p.Email,
		Name:      p.Name,
		CreatedAt: p.CreatedAt,
	}
	return body
}

// NewGetAllSubscriptionOK builds a "subscription" service "getAll" endpoint
// result from a HTTP "OK" response.
func NewGetAllSubscriptionOK(body []*SubscriptionResponse) []*subscription.Subscription {
	v := make([]*subscription.Subscription, len(body))
	for i, val := range body {
		v[i] = unmarshalSubscriptionResponseToSubscriptionSubscription(val)
	}

	return v
}

// NewGetAllServerError builds a subscription service getAll endpoint
// ServerError error.
func NewGetAllServerError(body *GetAllServerErrorResponseBody) *goa.ServiceError {
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

// NewGetOneByUIDSubscriptionOK builds a "subscription" service "getOneByUID"
// endpoint result from a HTTP "OK" response.
func NewGetOneByUIDSubscriptionOK(body *GetOneByUIDResponseBody) *subscription.Subscription {
	v := &subscription.Subscription{
		UID:       body.UID,
		Email:     body.Email,
		Name:      body.Name,
		CreatedAt: body.CreatedAt,
	}

	return v
}

// NewGetOneByUIDSubscriptionNotFound builds a subscription service getOneByUID
// endpoint SubscriptionNotFound error.
func NewGetOneByUIDSubscriptionNotFound(body *GetOneByUIDSubscriptionNotFoundResponseBody) *goa.ServiceError {
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

// NewGetOneByUIDServerError builds a subscription service getOneByUID endpoint
// ServerError error.
func NewGetOneByUIDServerError(body *GetOneByUIDServerErrorResponseBody) *goa.ServiceError {
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

// NewDeleteOneByUIDSubscriptionNotFound builds a subscription service
// deleteOneByUID endpoint SubscriptionNotFound error.
func NewDeleteOneByUIDSubscriptionNotFound(body *DeleteOneByUIDSubscriptionNotFoundResponseBody) *goa.ServiceError {
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

// NewDeleteOneByUIDServerError builds a subscription service deleteOneByUID
// endpoint ServerError error.
func NewDeleteOneByUIDServerError(body *DeleteOneByUIDServerErrorResponseBody) *goa.ServiceError {
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

// NewCreateOneSubscriptionOK builds a "subscription" service "createOne"
// endpoint result from a HTTP "OK" response.
func NewCreateOneSubscriptionOK(body *CreateOneResponseBody) *subscription.Subscription {
	v := &subscription.Subscription{
		UID:       body.UID,
		Email:     body.Email,
		Name:      body.Name,
		CreatedAt: body.CreatedAt,
	}

	return v
}

// NewCreateOneSubscriptionFieldMissing builds a subscription service createOne
// endpoint SubscriptionFieldMissing error.
func NewCreateOneSubscriptionFieldMissing(body *CreateOneSubscriptionFieldMissingResponseBody) *goa.ServiceError {
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

// NewCreateOneSubscriptionAlreadyExists builds a subscription service
// createOne endpoint SubscriptionAlreadyExists error.
func NewCreateOneSubscriptionAlreadyExists(body *CreateOneSubscriptionAlreadyExistsResponseBody) *goa.ServiceError {
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

// NewCreateOneServerError builds a subscription service createOne endpoint
// ServerError error.
func NewCreateOneServerError(body *CreateOneServerErrorResponseBody) *goa.ServiceError {
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

// ValidateGetAllServerErrorResponseBody runs the validations defined on
// getAll_ServerError_response_body
func ValidateGetAllServerErrorResponseBody(body *GetAllServerErrorResponseBody) (err error) {
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

// ValidateGetOneByUIDSubscriptionNotFoundResponseBody runs the validations
// defined on getOneByUID_SubscriptionNotFound_response_body
func ValidateGetOneByUIDSubscriptionNotFoundResponseBody(body *GetOneByUIDSubscriptionNotFoundResponseBody) (err error) {
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

// ValidateGetOneByUIDServerErrorResponseBody runs the validations defined on
// getOneByUID_ServerError_response_body
func ValidateGetOneByUIDServerErrorResponseBody(body *GetOneByUIDServerErrorResponseBody) (err error) {
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

// ValidateDeleteOneByUIDSubscriptionNotFoundResponseBody runs the validations
// defined on deleteOneByUID_SubscriptionNotFound_response_body
func ValidateDeleteOneByUIDSubscriptionNotFoundResponseBody(body *DeleteOneByUIDSubscriptionNotFoundResponseBody) (err error) {
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

// ValidateDeleteOneByUIDServerErrorResponseBody runs the validations defined
// on deleteOneByUID_ServerError_response_body
func ValidateDeleteOneByUIDServerErrorResponseBody(body *DeleteOneByUIDServerErrorResponseBody) (err error) {
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

// ValidateCreateOneSubscriptionFieldMissingResponseBody runs the validations
// defined on createOne_SubscriptionFieldMissing_response_body
func ValidateCreateOneSubscriptionFieldMissingResponseBody(body *CreateOneSubscriptionFieldMissingResponseBody) (err error) {
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

// ValidateCreateOneSubscriptionAlreadyExistsResponseBody runs the validations
// defined on createOne_SubscriptionAlreadyExists_response_body
func ValidateCreateOneSubscriptionAlreadyExistsResponseBody(body *CreateOneSubscriptionAlreadyExistsResponseBody) (err error) {
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

// ValidateCreateOneServerErrorResponseBody runs the validations defined on
// createOne_ServerError_response_body
func ValidateCreateOneServerErrorResponseBody(body *CreateOneServerErrorResponseBody) (err error) {
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
