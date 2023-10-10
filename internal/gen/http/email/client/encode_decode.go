// Code generated by goa v3.13.1, DO NOT EDIT.
//
// email HTTP client encoders and decoders
//
// Command:
// $ goa gen github.com/tomchan123/go-news-mailer/internal/design --output
// ./internal

package client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"

	email "github.com/tomchan123/go-news-mailer/internal/gen/email"
	goahttp "goa.design/goa/v3/http"
)

// BuildSendOneEmailRequest instantiates a HTTP request object with method and
// path set to call the "email" service "sendOneEmail" endpoint
func (c *Client) BuildSendOneEmailRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: SendOneEmailEmailPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("email", "sendOneEmail", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeSendOneEmailRequest returns an encoder for requests sent to the email
// sendOneEmail server.
func EncodeSendOneEmailRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*email.SubscriptionSendOneEmailPayload)
		if !ok {
			return goahttp.ErrInvalidType("email", "sendOneEmail", "*email.SubscriptionSendOneEmailPayload", v)
		}
		body := NewSendOneEmailRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("email", "sendOneEmail", err)
		}
		return nil
	}
}

// DecodeSendOneEmailResponse returns a decoder for responses returned by the
// email sendOneEmail endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeSendOneEmailResponse may return the following errors:
//   - "SubscriptionFieldMissing" (type *goa.ServiceError): http.StatusBadRequest
//   - "ServerError" (type *goa.ServiceError): http.StatusInternalServerError
//   - error: internal error
func DecodeSendOneEmailResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
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
		case http.StatusBadRequest:
			var (
				body SendOneEmailSubscriptionFieldMissingResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("email", "sendOneEmail", err)
			}
			err = ValidateSendOneEmailSubscriptionFieldMissingResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("email", "sendOneEmail", err)
			}
			return nil, NewSendOneEmailSubscriptionFieldMissing(&body)
		case http.StatusInternalServerError:
			var (
				body SendOneEmailServerErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("email", "sendOneEmail", err)
			}
			err = ValidateSendOneEmailServerErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("email", "sendOneEmail", err)
			}
			return nil, NewSendOneEmailServerError(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("email", "sendOneEmail", resp.StatusCode, string(body))
		}
	}
}
