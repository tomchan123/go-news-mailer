// Code generated by goa v3.13.1, DO NOT EDIT.
//
// email HTTP client CLI support package
//
// Command:
// $ goa gen github.com/tomchan123/go-news-mailer/internal/design --output
// ./internal

package client

import (
	"encoding/json"
	"fmt"

	email "github.com/tomchan123/go-news-mailer/internal/gen/email"
)

// BuildSendOneEmailPayload builds the payload for the email sendOneEmail
// endpoint from CLI flags.
func BuildSendOneEmailPayload(emailSendOneEmailBody string) (*email.SubscriptionSendOneEmailPayload, error) {
	var err error
	var body SendOneEmailRequestBody
	{
		err = json.Unmarshal([]byte(emailSendOneEmailBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"createdAt\": \"2012-04-23T18:25:43.511Z\",\n      \"email\": \"user@email.com\",\n      \"name\": \"John Doe\",\n      \"uid\": \"abcd1234\"\n   }'")
		}
	}
	v := &email.SubscriptionSendOneEmailPayload{
		UID:       body.UID,
		Email:     body.Email,
		Name:      body.Name,
		CreatedAt: body.CreatedAt,
	}

	return v, nil
}
