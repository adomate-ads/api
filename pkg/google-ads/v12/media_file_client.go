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

var newMediaFileClientHook clientHook

// MediaFileCallOptions contains the retry settings for each method of MediaFileClient.
type MediaFileCallOptions struct {
	MutateMediaFiles []gax.CallOption
}

func defaultMediaFileGRPCClientOptions() []option.ClientOption {
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

func defaultMediaFileCallOptions() *MediaFileCallOptions {
	return &MediaFileCallOptions{
		MutateMediaFiles: []gax.CallOption{
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

// internalMediaFileClient is an interface that defines the methods available from Google Ads API.
type internalMediaFileClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	MutateMediaFiles(context.Context, *servicespb.MutateMediaFilesRequest, ...gax.CallOption) (*servicespb.MutateMediaFilesResponse, error)
}

// MediaFileClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to manage media files.
type MediaFileClient struct {
	// The internal transport-dependent client.
	internalClient internalMediaFileClient

	// The call options for this service.
	CallOptions *MediaFileCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *MediaFileClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *MediaFileClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *MediaFileClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// MutateMediaFiles creates media files. Operation statuses are returned.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// DatabaseError (at )
// DistinctError (at )
// FieldError (at )
// HeaderError (at )
// IdError (at )
// ImageError (at )
// InternalError (at )
// MediaBundleError (at )
// MediaFileError (at )
// NewResourceCreationError (at )
// NotEmptyError (at )
// NullError (at )
// OperatorError (at )
// QuotaError (at )
// RangeError (at )
// RequestError (at )
// SizeLimitError (at )
// StringFormatError (at )
// StringLengthError (at )
func (c *MediaFileClient) MutateMediaFiles(ctx context.Context, req *servicespb.MutateMediaFilesRequest, opts ...gax.CallOption) (*servicespb.MutateMediaFilesResponse, error) {
	return c.internalClient.MutateMediaFiles(ctx, req, opts...)
}

// mediaFileGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type mediaFileGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing MediaFileClient
	CallOptions **MediaFileCallOptions

	// The gRPC API client.
	mediaFileClient servicespb.MediaFileServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewMediaFileClient creates a new media file service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to manage media files.
func NewMediaFileClient(ctx context.Context, opts ...option.ClientOption) (*MediaFileClient, error) {
	clientOpts := defaultMediaFileGRPCClientOptions()
	if newMediaFileClientHook != nil {
		hookOpts, err := newMediaFileClientHook(ctx, clientHookParams{})
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
	client := MediaFileClient{CallOptions: defaultMediaFileCallOptions()}

	c := &mediaFileGRPCClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		mediaFileClient:  servicespb.NewMediaFileServiceClient(connPool),
		CallOptions:      &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *mediaFileGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *mediaFileGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *mediaFileGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *mediaFileGRPCClient) MutateMediaFiles(ctx context.Context, req *servicespb.MutateMediaFilesRequest, opts ...gax.CallOption) (*servicespb.MutateMediaFilesResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 14400000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).MutateMediaFiles[0:len((*c.CallOptions).MutateMediaFiles):len((*c.CallOptions).MutateMediaFiles)], opts...)
	var resp *servicespb.MutateMediaFilesResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.mediaFileClient.MutateMediaFiles(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
