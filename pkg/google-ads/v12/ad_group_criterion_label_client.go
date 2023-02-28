// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package googleads

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	servicespb "github.com/adomate-ads/api/pkg/google-ads/pb/v12/services"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

var newAdGroupCriterionLabelClientHook clientHook

// AdGroupCriterionLabelCallOptions contains the retry settings for each method of AdGroupCriterionLabelClient.
type AdGroupCriterionLabelCallOptions struct {
	MutateAdGroupCriterionLabels []gax.CallOption
}

func defaultAdGroupCriterionLabelGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("googleads.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("googleads.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://googleads.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultAdGroupCriterionLabelCallOptions() *AdGroupCriterionLabelCallOptions {
	return &AdGroupCriterionLabelCallOptions{
		MutateAdGroupCriterionLabels: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    5000 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
	}
}

// internalAdGroupCriterionLabelClient is an interface that defines the methods available from Google Ads API.
type internalAdGroupCriterionLabelClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	MutateAdGroupCriterionLabels(context.Context, *servicespb.MutateAdGroupCriterionLabelsRequest, ...gax.CallOption) (*servicespb.MutateAdGroupCriterionLabelsResponse, error)
}

// AdGroupCriterionLabelClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to manage labels on ad group criteria.
type AdGroupCriterionLabelClient struct {
	// The internal transport-dependent client.
	internalClient internalAdGroupCriterionLabelClient

	// The call options for this service.
	CallOptions *AdGroupCriterionLabelCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *AdGroupCriterionLabelClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *AdGroupCriterionLabelClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *AdGroupCriterionLabelClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// MutateAdGroupCriterionLabels creates and removes ad group criterion labels.
// Operation statuses are returned.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// DatabaseError (at )
// FieldError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RequestError (at )
func (c *AdGroupCriterionLabelClient) MutateAdGroupCriterionLabels(ctx context.Context, req *servicespb.MutateAdGroupCriterionLabelsRequest, opts ...gax.CallOption) (*servicespb.MutateAdGroupCriterionLabelsResponse, error) {
	return c.internalClient.MutateAdGroupCriterionLabels(ctx, req, opts...)
}

// adGroupCriterionLabelGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type adGroupCriterionLabelGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing AdGroupCriterionLabelClient
	CallOptions **AdGroupCriterionLabelCallOptions

	// The gRPC API client.
	adGroupCriterionLabelClient servicespb.AdGroupCriterionLabelServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewAdGroupCriterionLabelClient creates a new ad group criterion label service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to manage labels on ad group criteria.
func NewAdGroupCriterionLabelClient(ctx context.Context, opts ...option.ClientOption) (*AdGroupCriterionLabelClient, error) {
	clientOpts := defaultAdGroupCriterionLabelGRPCClientOptions()
	if newAdGroupCriterionLabelClientHook != nil {
		hookOpts, err := newAdGroupCriterionLabelClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	disableDeadlines, err := checkDisableDeadlines()
	if err != nil {
		return nil, err
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := AdGroupCriterionLabelClient{CallOptions: defaultAdGroupCriterionLabelCallOptions()}

	c := &adGroupCriterionLabelGRPCClient{
		connPool:                    connPool,
		disableDeadlines:            disableDeadlines,
		adGroupCriterionLabelClient: servicespb.NewAdGroupCriterionLabelServiceClient(connPool),
		CallOptions:                 &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *adGroupCriterionLabelGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *adGroupCriterionLabelGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *adGroupCriterionLabelGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *adGroupCriterionLabelGRPCClient) MutateAdGroupCriterionLabels(ctx context.Context, req *servicespb.MutateAdGroupCriterionLabelsRequest, opts ...gax.CallOption) (*servicespb.MutateAdGroupCriterionLabelsResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 14400000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).MutateAdGroupCriterionLabels[0:len((*c.CallOptions).MutateAdGroupCriterionLabels):len((*c.CallOptions).MutateAdGroupCriterionLabels)], opts...)
	var resp *servicespb.MutateAdGroupCriterionLabelsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.adGroupCriterionLabelClient.MutateAdGroupCriterionLabels(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
