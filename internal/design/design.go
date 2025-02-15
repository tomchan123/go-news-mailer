package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("news", func() {
	Title("AI News Mailer")
	Description("A REST service to interact with an emailer that sends AI-summarised news")
	Version("1.0")

	HTTP(func() {
		Path("/api")
	})

	Server("newsSvr", func() {
		Services("subscription", "email")

		Host("development", func() { URI("http://localhost:8888") })
	})
})

var _ = Service("subscription", func() {
	Description("The subscription service provides ways to manage news subscription")

	HTTP(func() {
		Path("/subscriptions")
		Response("ServerError", StatusInternalServerError)
	})

	Error("ServerError", func() {
		Description("Error returned when there is an internal server error")
	})
	Error("SubscriptionNotFound", func() {
		Description("Error returned when the specified subscription does not exist")
	})
	Error("SubscriptionFieldMissing", func() {
		Description("Error returned when the subscription has missing field(s)")
	})
	Error("SubscriptionAlreadyExists", func() {
		Description("Error returned when the subscription already exists")
	})

	Method("getAll", func() {
		Description("Get all subscriptions")

		Result(ArrayOf(TSubscription), "A list of subscriptions")

		HTTP(func() {
			GET("")

			Response(StatusOK)
		})
	})

	Method("getOneByUID", func() {
		Description("Get a subscription by UID")

		Payload(String, "UID of subscription", func() {
			Example("abcd1234")
		})
		Result(TSubscription, "A subscription with the same UID")

		HTTP(func() {
			GET("/{uid}")

			Response(StatusOK)
			Response("SubscriptionNotFound", StatusNotFound)
		})
	})

	Method("deleteOneByUID", func() {
		Description("Delete a subscription by UID")

		Payload(String, "UID of subscription", func() {
			Example("abcd1234")
		})
		Result(Empty)

		HTTP(func() {
			DELETE("/{uid}")

			Response(StatusOK)
			Response("SubscriptionNotFound", StatusNotFound)
		})
	})

	Method("createOne", func() {
		Description("Create a new subscription")

		Payload(TSubscription, "Subscription to be created", func() {
			Required("email")
		})
		Result(TSubscription, "Created subscription")

		HTTP(func() {
			POST("")

			Response(StatusOK)
			Response("SubscriptionFieldMissing", StatusBadRequest)
			Response("SubscriptionAlreadyExists", StatusConflict)
		})
	})

	// Method("updateOne", func() {
	// 	Description("Update a single subscription")

	// 	Payload(TSubscription, "Subscription to be updated")
	// 	Result(TSubscription, "Updated subscription")

	// 	HTTP(func() {
	// 		PUT("/{uid}")
	// 		Response(StatusOK)
	// 	})
	// })
})

var _ = Service("email", func() {
	Description("The email service allow sending of news email through API")

	HTTP(func() {
		Path("/email")
		Response("ServerError", StatusInternalServerError)
	})

	Error("ServerError", func() {
		Description("Error returned when there is an internal server error")
	})
	Error("SubscriptionFieldMissing", func() {
		Description("Error returned when the subscription has missing field(s)")
	})

	Method("sendOneEmail", func() {
		Description("Send an email to the target subscription")

		Payload(TSubscription, "Subscription to send an email to", func() {
			Required("email")
		})

		Result(Empty)

		HTTP(func() {
			POST("")

			Response(StatusOK)
			Response("SubscriptionFieldMissing", StatusBadRequest)
		})
	})
})
