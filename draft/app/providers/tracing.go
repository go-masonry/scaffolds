package providers

import (
	"context"
	"github.com/go-masonry/bjaeger"
	"github.com/go-masonry/mortar/interfaces/cfg"
	"github.com/go-masonry/mortar/interfaces/log"
	"github.com/go-masonry/mortar/middleware/interceptors/trace"
	"github.com/go-masonry/mortar/mortar"
	opentracing "github.com/opentracing/opentracing-go"
	"go.uber.org/fx"
)

func TracerOption() fx.Option {
	return fx.Options(
		fx.Provide(tracerBuilder),
		trace.GRPCTracingUnaryInterceptorOption(),
		trace.TracerGRPCClientInterceptorOption(),
		trace.TracerRESTClientInterceptorOption(),
	)
}

// This constructor assumes you have JAEGER environment variables set
// https://github.com/jaegertracing/jaeger-client-go#environment-variables
func tracerBuilder(lc fx.Lifecycle, config cfg.Config, logger log.Logger) (opentracing.Tracer, error) {
	openTracer, err := bjaeger.Builder().
		SetServiceName(config.Get(mortar.Name).String()).
		AddOptions(bjaeger.BricksLoggerOption(logger)).
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
