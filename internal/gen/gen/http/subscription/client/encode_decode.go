// Code generated by goa v3.13.1, DO NOT EDIT.
//
// subscription HTTP client encoders and decoders
//
// Command:
// $ goa gen github.com/tomchan123/go-news-mailer/internal/design --output
// ./internal/gen

package client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"

	subscription "github.com/tomchan123/go-news-mailer/internal/gen/gen/subscription"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// BuildGetAllRequest instantiates a HTTP request object with method and path
// set to call the "subscription" service "getAll" endpoint
func (c *Client) BuildGetAllRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GetAllSubscriptionPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("subscription", "getAll", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeGetAllResponse returns a decoder for responses returned by the
// subscription getAll endpoint. restoreBody controls whether the response body
// should be restored after having been read.
func DecodeGetAllResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body GetAllResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("subscription", "getAll", err)
			}
			for _, e := range body {
				if e != nil {
					if err2 := ValidateSubscriptionResponse(e); err2 != nil {
						err = goa.MergeErrors(err, err2)
					}
				}
			}
			if err != nil {
				return nil, goahttp.ErrValidationError("subscription", "getAll", err)
			}
			res := NewGetAllSubscriptionOK(body)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("subscription", "getAll", resp.StatusCode, string(body))
		}
	}
}

// BuildGetOneByUIDRequest instantiates a HTTP request object with method and
// path set to call the "subscription" service "getOneByUID" endpoint
func (c *Client) BuildGetOneByUIDRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		uid string
	)
	{
		p, ok := v.(string)
		if !ok {
			return nil, goahttp.ErrInvalidType("subscription", "getOneByUID", "string", v)
		}
		uid = p
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GetOneByUIDSubscriptionPath(uid)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("subscription", "getOneByUID", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeGetOneByUIDResponse returns a decoder for responses returned by the
// subscription getOneByUID endpoint. restoreBody controls whether the response
// body should be restored after having been read.
// DecodeGetOneByUIDResponse may return the following errors:
//   - "SubscriptionNotFound" (type *goa.ServiceError): http.StatusNotFound
//   - error: internal error
func DecodeGetOneByUIDResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body GetOneByUIDResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("subscription", "getOneByUID", err)
			}
			err = ValidateGetOneByUIDResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("subscription", "getOneByUID", err)
			}
			res := NewGetOneByUIDSubscriptionOK(&body)
			return res, nil
		case http.StatusNotFound:
			var (
				body GetOneByUIDSubscriptionNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("subscription", "getOneByUID", err)
			}
			err = ValidateGetOneByUIDSubscriptionNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("subscription", "getOneByUID", err)
			}
			return nil, NewGetOneByUIDSubscriptionNotFound(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("subscription", "getOneByUID", resp.StatusCode, string(body))
		}
	}
}

// BuildDeleteOneByUIDRequest instantiates a HTTP request object with method
// and path set to call the "subscription" service "deleteOneByUID" endpoint
func (c *Client) BuildDeleteOneByUIDRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		uid string
	)
	{
		p, ok := v.(string)
		if !ok {
			return nil, goahttp.ErrInvalidType("subscription", "deleteOneByUID", "string", v)
		}
		uid = p
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: DeleteOneByUIDSubscriptionPath(uid)}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("subscription", "deleteOneByUID", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeDeleteOneByUIDResponse returns a decoder for responses returned by the
// subscription deleteOneByUID endpoint. restoreBody controls whether the
// response body should be restored after having been read.
// DecodeDeleteOneByUIDResponse may return the following errors:
//   - "SubscriptionNotFound" (type *goa.ServiceError): http.StatusNotFound
//   - error: internal error
func DecodeDeleteOneByUIDResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			return nil, nil
		case http.StatusNotFound:
			var (
				body DeleteOneByUIDSubscriptionNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("subscription", "deleteOneByUID", err)
			}
			err = ValidateDeleteOneByUIDSubscriptionNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("subscription", "deleteOneByUID", err)
			}
			return nil, NewDeleteOneByUIDSubscriptionNotFound(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("subscription", "deleteOneByUID", resp.StatusCode, string(body))
		}
	}
}

// unmarshalSubscriptionResponseToSubscriptionSubscription builds a value of
// type *subscription.Subscription from a value of type *SubscriptionResponse.
func unmarshalSubscriptionResponseToSubscriptionSubscription(v *SubscriptionResponse) *subscription.Subscription {
	res := &subscription.Subscription{
		UID:   *v.UID,
		Email: *v.Email,
		Name:  v.Name,
		Since: v.Since,
	}

	return res
}
