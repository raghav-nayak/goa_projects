// Code generated by goa v3.21.1, DO NOT EDIT.
//
// greeter gRPC client encoders and decoders
//
// Command:
// $ goa gen grpcgreeter/design

package client

import (
	"context"
	greeter "grpcgreeter/gen/greeter"
	greeterpb "grpcgreeter/gen/grpc/greeter/pb"

	goagrpc "goa.design/goa/v3/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// BuildSayHelloFunc builds the remote method to invoke for "greeter" service
// "SayHello" endpoint.
func BuildSayHelloFunc(grpccli greeterpb.GreeterClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb any, opts ...grpc.CallOption) (any, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.SayHello(ctx, reqpb.(*greeterpb.SayHelloRequest), opts...)
		}
		return grpccli.SayHello(ctx, &greeterpb.SayHelloRequest{}, opts...)
	}
}

// EncodeSayHelloRequest encodes requests sent to greeter SayHello endpoint.
func EncodeSayHelloRequest(ctx context.Context, v any, md *metadata.MD) (any, error) {
	payload, ok := v.(*greeter.SayHelloPayload)
	if !ok {
		return nil, goagrpc.ErrInvalidType("greeter", "SayHello", "*greeter.SayHelloPayload", v)
	}
	return NewProtoSayHelloRequest(payload), nil
}

// DecodeSayHelloResponse decodes responses from the greeter SayHello endpoint.
func DecodeSayHelloResponse(ctx context.Context, v any, hdr, trlr metadata.MD) (any, error) {
	message, ok := v.(*greeterpb.SayHelloResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("greeter", "SayHello", "*greeterpb.SayHelloResponse", v)
	}
	res := NewSayHelloResult(message)
	return res, nil
}
