// Code generated by goa v3.13.1, DO NOT EDIT.
//
// subscription HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/tomchan123/go-news-mailer/internal/design --output
// ./internal/gen

package server

import (
	"context"
	"errors"
	"net/http"

	subscription "github.com/tomchan123/go-news-mailer/internal/gen/gen/subscription"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeGetAllResponse returns an encoder for responses returned by the
// subscription getAll endpoint.
func EncodeGetAllResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.([]*subscription.Subscription)
		enc := encoder(ctx, w)
		body := NewGetAllResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// EncodeGetOneByUIDResponse returns an encoder for responses returned by the
// subscription getOneByUID endpoint.
func EncodeGetOneByUIDResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.(*subscription.Subscription)
		enc := encoder(ctx, w)
		body := NewGetOneByUIDResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeGetOneByUIDRequest returns a decoder for requests sent to the
// subscription getOneByUID endpoint.
func DecodeGetOneByUIDRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			uid string

			params = mux.Vars(r)
		)
		uid = params["uid"]
		payload := uid

		return payload, nil
	}
}

// EncodeGetOneByUIDError returns an encoder for errors returned by the
// getOneByUID subscription endpoint.
func EncodeGetOneByUIDError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "SubscriptionNotFound":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewGetOneByUIDSubscriptionNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeDeleteOneByUIDResponse returns an encoder for responses returned by
// the subscription deleteOneByUID endpoint.
func EncodeDeleteOneByUIDResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		w.WriteHeader(http.StatusOK)
		return nil
	}
}

// DecodeDeleteOneByUIDRequest returns a decoder for requests sent to the
// subscription deleteOneByUID endpoint.
func DecodeDeleteOneByUIDRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			uid string

			params = mux.Vars(r)
		)
		uid = params["uid"]
		payload := uid

		return payload, nil
	}
}

// EncodeDeleteOneByUIDError returns an encoder for errors returned by the
// deleteOneByUID subscription endpoint.
func EncodeDeleteOneByUIDError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "SubscriptionNotFound":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewDeleteOneByUIDSubscriptionNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// marshalSubscriptionSubscriptionToSubscriptionResponse builds a value of type
// *SubscriptionResponse from a value of type *subscription.Subscription.
func marshalSubscriptionSubscriptionToSubscriptionResponse(v *subscription.Subscription) *SubscriptionResponse {
	res := &SubscriptionResponse{
		UID:   v.UID,
		Email: v.Email,
		Name:  v.Name,
		Since: v.Since,
	}

	return res
}
