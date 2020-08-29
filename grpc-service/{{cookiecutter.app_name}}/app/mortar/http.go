package mortar

import (
	"github.com/go-masonry/mortar/providers"
	"go.uber.org/fx"
)

// HTTPClientFxOptions adds miscellaneous middleware for Clients
func HTTPClientFxOptions() fx.Option {
	return fx.Options(
		providers.HTTPClientBuildersFxOption(),               // client builders
		providers.CopyGRPCHeadersClientInterceptorFxOption(), // copy selected gRPC headers from metadata.Incoming to metadata.Outgoing
	)
}

// HTTPServerFxOptions adds miscellaneous middleware for gRPC server
func HTTPServerFxOptions() fx.Option {
	return fx.Options(
		providers.HTTPServerBuilderFxOption(),     // Web Server Builder
		providers.LoggerGRPCInterceptorFxOption(), // output incoming/outgoing gRPC requests to log
	)
}

// InternalHTTPHandlersFxOptions these will help you to debug/profile or understand the internals of your service
func InternalHTTPHandlersFxOptions() fx.Option {
	return fx.Options(
		providers.InternalDebugHandlersFxOption(),
		providers.InternalProfileHandlerFunctionsFxOption(),
		providers.InternalSelfHandlersFxOption(),
	)
}
