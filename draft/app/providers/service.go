package providers

import (
	"context"
	"github.com/go-masonry/mortar/constructors"
	"github.com/go-masonry/mortar/constructors/partial"
	"github.com/go-masonry/mortar/http/server"
	"github.com/go-masonry/mortar/interfaces/cfg"
	serverInt "github.com/go-masonry/mortar/interfaces/http/server"
	"github.com/go-masonry/mortar/interfaces/log"
	"github.com/go-masonry/mortar/mortar"
	workshop "github.com/go-masonry/scaffolds/draft/api"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	opentracing "github.com/opentracing/opentracing-go"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

type httpServiceDeps struct {
	fx.In

	Logger log.Logger
	Config cfg.Config
	Tracer opentracing.Tracer `optional:"true"`
	// API Implementations
	Workshop    workshop.WorkshopServer
	SubWorkshop workshop.SubWorkshopServer
	Builder     serverInt.GRPCWebServiceBuilder
}

func InvokeServiceOption() fx.Option {
	return fx.Options(
		fx.Provide(partial.HttpServerBuilder),
		fx.Provide(serviceSetup),
		// This one starts the server and tells fx to build it's dependency graph
		fx.Invoke(constructors.Service),
	)
}

func serviceSetup(deps httpServiceDeps) (serverInt.WebService, error) {
	builder := deps.Builder.
		SetLogger(deps.Logger.Debug).
		RegisterGRPCAPIs(deps.gRPCServerAPIs) // setup grpc api
	builder = deps.configureGRPCGateway(builder) // setup rest api over grpc
	return builder.Build()
}

func (deps httpServiceDeps) configureGRPCGateway(builder serverInt.GRPCWebServiceBuilder) serverInt.GRPCWebServiceBuilder {
	if externalRESTPort := deps.Config.Get(mortar.ServerRESTExternalPort); externalRESTPort.IsSet() {
		return builder.
			AddRESTServerConfiguration().
			ListenOn(":" + externalRESTPort.String()).
			AddGRPCGatewayHandlers(deps.gRPCGatewayHandlers()...).
			AddGRPCGatewayOptions(server.MetadataTraceCarrierOption(deps.Tracer)).
			// TODO Add custom header matchers
			BuildRESTPart()
	}
	return builder
}

func (deps httpServiceDeps) gRPCServerAPIs(srv *grpc.Server) {
	workshop.RegisterWorkshopServer(srv, deps.Workshop)
	workshop.RegisterSubWorkshopServer(srv, deps.SubWorkshop)
	// Any additional gRPC Implementations should be called here
}

func (deps httpServiceDeps) gRPCGatewayHandlers() []func(mux *runtime.ServeMux, endpoint string) error {
	return []func(mux *runtime.ServeMux, endpoint string) error{
		// Register workshop REST API
		func(mux *runtime.ServeMux, endpoint string) error {
			return workshop.RegisterWorkshopHandlerFromEndpoint(context.Background(), mux, endpoint, []grpc.DialOption{grpc.WithInsecure()})
		},
		// Register sub workshop REST API
		func(mux *runtime.ServeMux, endpoint string) error {
			return workshop.RegisterSubWorkshopHandlerFromEndpoint(context.Background(), mux, endpoint, []grpc.DialOption{grpc.WithInsecure()})
		},
		// Any additional gRPC gateway registrations should be called here
	}
}
