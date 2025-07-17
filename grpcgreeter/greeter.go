package grpcgreeter

import (
    "context"
    "fmt"

    gengreeter "grpcgreeter/gen/greeter"
)

// GreeterService implements the greeter Service interface
type GreeterService struct{}

// NewGreeterService creates a new GreeterService instance
func NewGreeterService() *GreeterService {
    return &GreeterService{}
}

// SayHello implements the SayHello logic the greeter service
func (s *GreeterService) SayHello(ctx context.Context, p *gengreeter.SayHelloPayload) (*gengreeter.SayHelloResult, error) {
    // Add input validation if needed
    if p.Name == "" {
        return nil, fmt.Errorf("name cannot be empty")
    }

    // build the greeting
    greeting := fmt.Sprintf("Hello, %s!", p.Name)

    // Return the result
    return &gengreeter.SayHelloResult{
        Greeting: greeting,
    }, nil
}