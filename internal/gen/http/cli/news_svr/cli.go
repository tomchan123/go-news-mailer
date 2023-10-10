// Code generated by goa v3.13.1, DO NOT EDIT.
//
// newsSvr HTTP client CLI support package
//
// Command:
// $ goa gen github.com/tomchan123/go-news-mailer/internal/design --output
// ./internal

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	emailc "github.com/tomchan123/go-news-mailer/internal/gen/http/email/client"
	subscriptionc "github.com/tomchan123/go-news-mailer/internal/gen/http/subscription/client"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//	command (subcommand1|subcommand2|...)
func UsageCommands() string {
	return `subscription (get-all|get-one-by-uid|delete-one-by-uid|create-one)
email send-one-email
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` subscription get-all` + "\n" +
		os.Args[0] + ` email send-one-email --body '{
      "createdAt": "2012-04-23T18:25:43.511Z",
      "email": "user@email.com",
      "name": "John Doe",
      "uid": "abcd1234"
   }'` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, any, error) {
	var (
		subscriptionFlags = flag.NewFlagSet("subscription", flag.ContinueOnError)

		subscriptionGetAllFlags = flag.NewFlagSet("get-all", flag.ExitOnError)

		subscriptionGetOneByUIDFlags = flag.NewFlagSet("get-one-by-uid", flag.ExitOnError)
		subscriptionGetOneByUIDPFlag = subscriptionGetOneByUIDFlags.String("p", "REQUIRED", "UID of subscription")

		subscriptionDeleteOneByUIDFlags = flag.NewFlagSet("delete-one-by-uid", flag.ExitOnError)
		subscriptionDeleteOneByUIDPFlag = subscriptionDeleteOneByUIDFlags.String("p", "REQUIRED", "UID of subscription")

		subscriptionCreateOneFlags    = flag.NewFlagSet("create-one", flag.ExitOnError)
		subscriptionCreateOneBodyFlag = subscriptionCreateOneFlags.String("body", "REQUIRED", "")

		emailFlags = flag.NewFlagSet("email", flag.ContinueOnError)

		emailSendOneEmailFlags    = flag.NewFlagSet("send-one-email", flag.ExitOnError)
		emailSendOneEmailBodyFlag = emailSendOneEmailFlags.String("body", "REQUIRED", "")
	)
	subscriptionFlags.Usage = subscriptionUsage
	subscriptionGetAllFlags.Usage = subscriptionGetAllUsage
	subscriptionGetOneByUIDFlags.Usage = subscriptionGetOneByUIDUsage
	subscriptionDeleteOneByUIDFlags.Usage = subscriptionDeleteOneByUIDUsage
	subscriptionCreateOneFlags.Usage = subscriptionCreateOneUsage

	emailFlags.Usage = emailUsage
	emailSendOneEmailFlags.Usage = emailSendOneEmailUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "subscription":
			svcf = subscriptionFlags
		case "email":
			svcf = emailFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "subscription":
			switch epn {
			case "get-all":
				epf = subscriptionGetAllFlags

			case "get-one-by-uid":
				epf = subscriptionGetOneByUIDFlags

			case "delete-one-by-uid":
				epf = subscriptionDeleteOneByUIDFlags

			case "create-one":
				epf = subscriptionCreateOneFlags

			}

		case "email":
			switch epn {
			case "send-one-email":
				epf = emailSendOneEmailFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     any
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "subscription":
			c := subscriptionc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "get-all":
				endpoint = c.GetAll()
				data = nil
			case "get-one-by-uid":
				endpoint = c.GetOneByUID()
				data = *subscriptionGetOneByUIDPFlag
			case "delete-one-by-uid":
				endpoint = c.DeleteOneByUID()
				data = *subscriptionDeleteOneByUIDPFlag
			case "create-one":
				endpoint = c.CreateOne()
				data, err = subscriptionc.BuildCreateOnePayload(*subscriptionCreateOneBodyFlag)
			}
		case "email":
			c := emailc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "send-one-email":
				endpoint = c.SendOneEmail()
				data, err = emailc.BuildSendOneEmailPayload(*emailSendOneEmailBodyFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// subscriptionUsage displays the usage of the subscription command and its
// subcommands.
func subscriptionUsage() {
	fmt.Fprintf(os.Stderr, `The subscription service provides ways to manage news subscription
Usage:
    %[1]s [globalflags] subscription COMMAND [flags]

COMMAND:
    get-all: Get all subscriptions
    get-one-by-uid: Get a subscription by UID
    delete-one-by-uid: Delete a subscription by UID
    create-one: Create a new subscription

Additional help:
    %[1]s subscription COMMAND --help
`, os.Args[0])
}
func subscriptionGetAllUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] subscription get-all

Get all subscriptions

Example:
    %[1]s subscription get-all
`, os.Args[0])
}

func subscriptionGetOneByUIDUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] subscription get-one-by-uid -p STRING

Get a subscription by UID
    -p STRING: UID of subscription

Example:
    %[1]s subscription get-one-by-uid --p "abcd1234"
`, os.Args[0])
}

func subscriptionDeleteOneByUIDUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] subscription delete-one-by-uid -p STRING

Delete a subscription by UID
    -p STRING: UID of subscription

Example:
    %[1]s subscription delete-one-by-uid --p "abcd1234"
`, os.Args[0])
}

func subscriptionCreateOneUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] subscription create-one -body JSON

Create a new subscription
    -body JSON: 

Example:
    %[1]s subscription create-one --body '{
      "createdAt": "2012-04-23T18:25:43.511Z",
      "email": "user@email.com",
      "name": "John Doe",
      "uid": "abcd1234"
   }'
`, os.Args[0])
}

// emailUsage displays the usage of the email command and its subcommands.
func emailUsage() {
	fmt.Fprintf(os.Stderr, `The email service allow sending of news email through API
Usage:
    %[1]s [globalflags] email COMMAND [flags]

COMMAND:
    send-one-email: Send an email to the target subscription

Additional help:
    %[1]s email COMMAND --help
`, os.Args[0])
}
func emailSendOneEmailUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] email send-one-email -body JSON

Send an email to the target subscription
    -body JSON: 

Example:
    %[1]s email send-one-email --body '{
      "createdAt": "2012-04-23T18:25:43.511Z",
      "email": "user@email.com",
      "name": "John Doe",
      "uid": "abcd1234"
   }'
`, os.Args[0])
}
