package greeter

import (
	"context"
	"net"
	"os"
	"os/signal"
	"syscall"

	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"grpcgreeter"

	gengreeter "grpcgreeter/gen/greeter"
	genpub "grpcgreeter/gen/grpc/greeter/pb"
	genserver "grpcgreeter/gen/grpc/greeter/server"
)

// loggingInterceptor is a simple gRPC interceptor that logs incoming requests
// intercepts every gRPC call before it reaches the service
// logs the method being called
// useful for debugging and monitoring
// can be extended for metrics, authentication, or other cross-cutting concerns
func loggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Printf("Handling request %s", info.FullMethod)
	return handler(ctx, req) // Call the handler to process the request and return the response or error
}

func main() {
	// create a TCP listener for incoming gRPC requests
	// this is where the gRPC server will listen for incoming requests
	listener, err := net.Listen("tcp", ":8090")
	if err != nil {
		log.Fatalf("failed to listen on port 8090: %v", err)
	}

	// create a new gRPC server with middleware(interceptor) support
	// the loggingInterceptor will log each request
	srv := grpc.NewServer(
		grpc.UnaryInterceptor(loggingInterceptor),
	)

	// Initialize the service
	svc := grpcgreeter.NewGreeterService()

    // create endpoints for the service
    // Wraps the service in Goa’s transport-agnostic endpoints
	endpoints := gengreeter.NewEndpoints(svc)

    // register service with gRPC server to handle incoming requests
    genpub.RegisterGreeterServer(srv, genserver.New(endpoints, nil))

    // enable server reflection for debugging tools
    // allowing tools like grpcurl to discover the service methods dynamically
    reflection.Register(srv)

    // handle graceful shutdown
    // listens for interrupt signal(ctrl+c) or termination requests
    // ensures in-flight requests complete before shutting down
    // prevents connection drops and data loss
    go func(){
        signalChannel := make(chan os.Signal, 1)
        signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)
        <- signalChannel
        log.Println("Shutting down gRPC server...")
        srv.GracefulStop()
    }()

    // start the gRPC server
    log.Println("Starting gRPC server on :8090")
    if err := srv.Serve(listener); err != nil {
        log.Fatalf("failed to serve gRPC server: %v", err)
    }

}
