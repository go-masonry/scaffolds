package mortar

import (
	"context"

	"github.com/go-masonry/bjaeger"
	"github.com/go-masonry/mortar/interfaces/cfg"
	"github.com/go-masonry/mortar/interfaces/log"
	"github.com/go-masonry/mortar/mortar"
	"github.com/go-masonry/mortar/providers"
	opentracing "github.com/opentracing/opentracing-go"
	"go.uber.org/fx"
)

func TracerFxOption() fx.Option {
	return fx.Options(
		fx.Provide(JaegerBuilder),
		providers.TracerGRPCClientInterceptorFxOption(),       // trace client span for gRPC clients
		providers.TracerRESTClientInterceptorFxOption(),       // trace client span for REST clients
		providers.GRPCTracingUnaryServerInterceptorFxOption(), // trace server span
		providers.GRPCGatewayMetadataTraceCarrierFxOption(),   // read it's documentation to understand better
	)
}

// This constructor assumes you have JAEGER environment variables set
//
// https://github.com/jaegertracing/jaeger-client-go#environment-variables
//
// Once built it will register Lifecycle hooks (connect on start, close on stop)
func JaegerBuilder(lc fx.Lifecycle, config cfg.Config, logger log.Logger) (opentracing.Tracer, error) {
	openTracer, err := bjaeger.Builder().
		SetServiceName(config.Get(mortar.Name).String()).
		AddOptions(bjaeger.BricksLoggerOption(logger)). // verbose logging,
		Build()
	if err != nil {
		return nil, err
	}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return openTracer.Connect(ctx)
		},
		OnStop: func(ctx context.Context) error {
			return openTracer.Close(ctx)
		},
	})
	return openTracer.Tracer(), nil
}
