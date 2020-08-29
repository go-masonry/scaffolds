package mortar

import (
	"github.com/go-masonry/mortar/providers/groups"
	serverInt "github.com/go-masonry/mortar/interfaces/http/server"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

type {{cookiecutter.app_name|lower}}ServiceDeps struct {
	fx.In

	// API Interfaces
	// Workshop    workshop.WorkshopServer
	// SubWorkshop workshop.SubWorkshopServer
}

func {{cookiecutter.app_name|capitalize}}APIsAndOtherDependenciesFxOption() fx.Option {
	return fx.Options(
		// GRPC Service APIs registration
		fx.Provide(fx.Annotated{
			Group:  groups.GRPCServerAPIs,
			Target: {{cookiecutter.app_name|lower}}GRPCServiceAPIs,
		}),
		// GRPC Gateway Generated Handlers registration
		fx.Provide(fx.Annotated{
			Group:  groups.GRPCGatewayGeneratedHandlers + ",flatten", // "flatten" does this [][]serverInt.GRPCGatewayGeneratedHandlers -> []serverInt.GRPCGatewayGeneratedHandlers
			Target: {{cookiecutter.app_name|lower}}GRPCGatewayHandlers,
		}),
		// All other {{cookiecutter.app_name}} dependencies
		{{cookiecutter.app_name|lower}}Dependencies(),
	)
}

func {{cookiecutter.app_name|lower}}GRPCServiceAPIs(deps {{cookiecutter.app_name|lower}}ServiceDeps) serverInt.GRPCServerAPI {
	return func(srv *grpc.Server) {
		panic("add your grpc server API implementation")
		// workshop.RegisterWorkshopServer(srv, deps.Workshop)
		// workshop.RegisterSubWorkshopServer(srv, deps.SubWorkshop)
		//
		// Any additional gRPC Implementations should be called here
	}
}

func {{cookiecutter.app_name|lower}}GRPCGatewayHandlers() []serverInt.GRPCGatewayGeneratedHandlers {
	return []serverInt.GRPCGatewayGeneratedHandlers{
		// Register your grpc-gateway REST API
		func(mux *runtime.ServeMux, endpoint string) error {
			panic("add your grpc-gateway handlers or remove this function")
			// return workshop.RegisterWorkshopHandlerFromEndpoint(context.Background(), mux, endpoint, []grpc.DialOption{grpc.WithInsecure()})
		},
		// Register additional grpc-gateway REST API
		func(mux *runtime.ServeMux, endpoint string) error {
			panic("add your grpc-gateway handlers or remove this function")
			// return workshop.RegisterSubWorkshopHandlerFromEndpoint(context.Background(), mux, endpoint, []grpc.DialOption{grpc.WithInsecure()})
		},
		// Any additional gRPC gateway registrations should be called here
	}
}

func {{cookiecutter.app_name}}Dependencies() fx.Option {
	return fx.Provide(
		// your constructors should be here
	)
}
