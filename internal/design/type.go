package design

import (
	. "goa.design/goa/v3/dsl"
)

var TSubscription = Type("Subscription", func() {
	Description("Representation of the subscription")

	Attribute("uid", String, "Unique identifier of subcription")
	Attribute("email", String, "Subscribed email")
	Attribute("name", String, "Name of the subscriber")
	Attribute("since", String, "Datetime when the subscription was made")

	Required("uid", "email")
})
