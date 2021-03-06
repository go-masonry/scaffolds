package mortar

import (
	"context"

	serverInt "github.com/go-masonry/mortar/interfaces/http/server"
	"github.com/go-masonry/mortar/providers/groups"
	workshop "{{cookiecutter.project_dir}}/{{cookiecutter.app_name}}/api"
	"{{cookiecutter.project_dir}}/{{cookiecutter.app_name}}/app/controllers"
	"{{cookiecutter.project_dir}}/{{cookiecutter.app_name}}/app/data"
	"{{cookiecutter.project_dir}}/{{cookiecutter.app_name}}/app/services"
	"{{cookiecutter.project_dir}}/{{cookiecutter.app_name}}/app/validations"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

type tutorialServiceDeps struct {
	fx.In

	// API Implementations
	Workshop    workshop.WorkshopServer
	SubWorkshop workshop.SubWorkshopServer
}

func TutorialAPIsAndOtherDependenciesFxOption() fx.Option {
	return fx.Options(
		// GRPC Service APIs registration
		fx.Provide(fx.Annotated{
			Group:  groups.GRPCServerAPIs,
			Target: tutorialGRPCServiceAPIs,
		}),
		// GRPC Gateway Generated Handlers registration
		fx.Provide(fx.Annotated{
			Group:  groups.GRPCGatewayGeneratedHandlers + ",flatten", // "flatten" does this [][]serverInt.GRPCGatewayGeneratedHandlers -> []serverInt.GRPCGatewayGeneratedHandlers
			Target: tutorialGRPCGatewayHandlers,
		}),
		// All other tutorial dependencies
		tutorialDependencies(),
	)
}

func tutorialGRPCServiceAPIs(deps tutorialServiceDeps) serverInt.GRPCServerAPI {
	return func(srv *grpc.Server) {
		workshop.RegisterWorkshopServer(srv, deps.Workshop)
		workshop.RegisterSubWorkshopServer(srv, deps.SubWorkshop)
		// Any additional gRPC Implementations should be called here
	}
}

func tutorialGRPCGatewayHandlers() []serverInt.GRPCGatewayGeneratedHandlers {
	return []serverInt.GRPCGatewayGeneratedHandlers{
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

func tutorialDependencies() fx.Option {
	return fx.Provide(
		services.CreateWorkshopService,
		services.CreateSubWorkshopService,
		controllers.CreateWorkshopController,
		controllers.CreateSubWorkshopController,
		data.CreateCarDB,
		validations.CreateWorkshopValidations,
		validations.CreateSubWorkshopValidations,
	)
}
