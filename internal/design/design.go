package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("news", func() {
	Title("AI News Mailer")
	Description("A REST service to interact with an emailer that sends AI-summarised news")

	Server("newsSvr", func() {
		Services("subscription")

		Host("development", func() { URI("http://localhost:8888/rest") })
	})
})

var _ = Service("subscription", func() {
	Description("The subscription service provides ways to manage news subscription")

	HTTP(func() {
		Path("/subscription")
	})

	Method("getAll", func() {
		Description("Get all subscriptions")

		Result(func() {
			Attribute("id", String, "Unique identifier of subscription")
		})

		HTTP(func() {
			GET("")
			Response(StatusOK)
		})
	})
})
