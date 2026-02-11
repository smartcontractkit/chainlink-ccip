package jd

import (
	"context"

	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// authTokenInterceptor returns a gRPC unary client interceptor that injects
// an OAuth2 Bearer token into the outgoing metadata of every request.
// The token is obtained (and auto-refreshed) from the given TokenSource.
func authTokenInterceptor(source oauth2.TokenSource) grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply any,
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		token, err := source.Token()
		if err != nil {
			return err
		}
		ctx = metadata.AppendToOutgoingContext(ctx, "authorization", "Bearer "+token.AccessToken)
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}
