// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: name.proto

/*
Package namepb is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package namepb

import (
	"context"
	"errors"
	"io"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Suppress "imported and not used" errors
var (
	_ codes.Code
	_ io.Reader
	_ status.Status
	_ = errors.New
	_ = runtime.String
	_ = utilities.NewDoubleArray
	_ = metadata.Join
)

func request_NameService_HealthCheck_0(ctx context.Context, marshaler runtime.Marshaler, client NameServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var (
		protoReq HealthCheckRequest
		metadata runtime.ServerMetadata
	)
	msg, err := client.HealthCheck(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err
}

func local_request_NameService_HealthCheck_0(ctx context.Context, marshaler runtime.Marshaler, server NameServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var (
		protoReq HealthCheckRequest
		metadata runtime.ServerMetadata
	)
	msg, err := server.HealthCheck(ctx, &protoReq)
	return msg, metadata, err
}

// RegisterNameServiceHandlerServer registers the http handlers for service NameService to "mux".
// UnaryRPC     :call NameServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterNameServiceHandlerFromEndpoint instead.
// GRPC interceptors will not work for this type of registration. To use interceptors, you must use the "runtime.WithMiddlewares" option in the "runtime.NewServeMux" call.
func RegisterNameServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server NameServiceServer) error {
	mux.Handle(http.MethodGet, pattern_NameService_HealthCheck_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		annotatedContext, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/namepb.NameService/HealthCheck", runtime.WithHTTPPathPattern("/api/v1/name/healthcheck"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_NameService_HealthCheck_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_NameService_HealthCheck_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})

	return nil
}

// RegisterNameServiceHandlerFromEndpoint is same as RegisterNameServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterNameServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.NewClient(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Errorf("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Errorf("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()
	return RegisterNameServiceHandler(ctx, mux, conn)
}

// RegisterNameServiceHandler registers the http handlers for service NameService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterNameServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterNameServiceHandlerClient(ctx, mux, NewNameServiceClient(conn))
}

// RegisterNameServiceHandlerClient registers the http handlers for service NameService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "NameServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "NameServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "NameServiceClient" to call the correct interceptors. This client ignores the HTTP middlewares.
func RegisterNameServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client NameServiceClient) error {
	mux.Handle(http.MethodGet, pattern_NameService_HealthCheck_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		annotatedContext, err := runtime.AnnotateContext(ctx, mux, req, "/namepb.NameService/HealthCheck", runtime.WithHTTPPathPattern("/api/v1/name/healthcheck"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_NameService_HealthCheck_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_NameService_HealthCheck_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	return nil
}

var (
	pattern_NameService_HealthCheck_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"api", "v1", "name", "healthcheck"}, ""))
)

var (
	forward_NameService_HealthCheck_0 = runtime.ForwardResponseMessage
)
