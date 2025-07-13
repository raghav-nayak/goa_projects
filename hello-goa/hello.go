package helloapi

import (
	"context"
	"fmt"
	hello "hello/gen/hello"

	"goa.design/clue/log"
)

// hello service example implementation.
// The example methods log the requests and return zero values.
type hellosrvc struct{}

// NewHello returns the hello service implementation.
func NewHello() hello.Service {
	return &hellosrvc{}
}

// To get hello
func (s *hellosrvc) GreetHello(ctx context.Context, p string) (res string, err error) {
	log.Printf(ctx, "hello.greetHello")
	return fmt.Sprintf("Hello, %s! from GreetHello", p), nil
}

// To respond to hello
func (s *hellosrvc) RespondToHello(ctx context.Context, p string) (res string, err error) {
	log.Printf(ctx, "hello.respondToHello")
	return fmt.Sprintf("Hello, %s! from RespondToHello", p), nil
}
