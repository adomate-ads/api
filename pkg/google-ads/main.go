package google_ads

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

const (
	EndPoint = "googleads.googleapis.com:443"
)

type (
	ContextOption func(*context.Context)
)

var conn *grpc.ClientConn

func GetGRPCConnection() *grpc.ClientConn {
	if conn != nil {
		return conn
	}
	cred := grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, ""))
	conn, err := grpc.Dial(EndPoint, cred)
	if err != nil {
		panic(err)
	}
	return conn
}

func WithContext(k string, v string) ContextOption {
	return func(ctx *context.Context) {
		*ctx = metadata.AppendToOutgoingContext(*ctx, k, v)
	}
}

func SetContext(ctx context.Context, opts ...ContextOption) context.Context {
	for _, opt := range opts {
		opt(&ctx)
	}
	return ctx
}
