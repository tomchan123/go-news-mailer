package design

import (
	. "goa.design/goa/v3/dsl"
)

var TSubscription = Type("Subscription", func() {
	Description("Representation of the subscription")

	Attribute("uid", String, "Unique identifier of subcription", func() {
		Example("abcd1234")
	})
	Attribute("email", String, "Subscribed email", func() {
		Example("user@email.com")
	})
	Attribute("name", String, "Name of the subscriber", func() {
		Example("John Doe")
	})
	Attribute("since", String, "Datetime when the subscription was made", func() {
		Example("2012-04-23T18:25:43.511Z")
	})
})
