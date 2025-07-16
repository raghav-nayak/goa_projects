package design

import (
    . "goa.design/goa/v3/dsl"
)

var _ = Service("greeter", func() {
    Description("A simple gRPC service that says hello.")

    // Defines the remote procedure call (RPC) method
    Method("SayHello", func() {
        Description("Send a greeting to a user")

        // Define the request payload (what the client sends)
        Payload(func ()  {
            Field(1, "name", String, "Name of the user to greet", func() { // number is the tag in the generated .proto file
                Example("Amex")
                MinLength(1)
            })
            Required("name")
        })

        // Note: Methods that support both HTTP and gRPC transports can use `Field` for defining fields (the tag is ignored for HTTP)

        Result(func() {
            Field(1, "greeting", String, "A friendly greeting message")
            Required("greeting")
        })

        // Indicate this method should be exposed over gRPC
        // ensures the generated code include .proto definitions and stubs for this method
        GRPC(func() {
            // The default code for a successful response is CodeOK(0)
            // We can also define custom mapping if needed
            // Response(CodeOk)
        })
    })

})