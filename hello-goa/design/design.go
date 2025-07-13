package design

import (
    . "goa.design/goa/v3/dsl"
)

var _ = Service("hello", func() {
    Description("A simple service that says hello")

    Method("greetHello", func() {
        Description("To get hello")
        Payload(String, "Name to greet")
        Result(String, "A greeting message")

        HTTP(func() {
            GET(("/greetHello/{name}"))
        })
    })

    Method("respondToHello", func() {
        Description("To respond to hello")
        Payload(String, "Name to greet")
        Result(String, "A response message")

        HTTP(func() {
            GET(("/respondHello/{name}"))
        })
    })
})