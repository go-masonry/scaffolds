package main

import (
	"github.com/alecthomas/kong"
	"github.com/go-masonry/mortar/handlers"
	"github.com/go-masonry/scaffolds/draft/app/controllers"
	"github.com/go-masonry/scaffolds/draft/app/db"
	"github.com/go-masonry/scaffolds/draft/app/providers"
	"github.com/go-masonry/scaffolds/draft/app/services"
	"github.com/go-masonry/scaffolds/draft/app/validations"
	"go.uber.org/fx"
)

var CLI struct {
	Config struct {
		Path            string   `arg:"" required:"" help:"Path to config file." type:"existingfile"`
		AdditionalFiles []string `optional:"" help:"Additional configuration files to merge, comma separated" type:"existingfile"`
	} `cmd:"" help:"Path to config file."`
}

func main() {
	ctx := kong.Parse(&CLI, kong.UsageOnError())
	switch cmd := ctx.Command(); cmd {
	case "config <path>":
		app := createApplication(CLI.Config.Path, CLI.Config.AdditionalFiles)
		app.Run()
	default:
		ctx.Fatalf("unknown option %s", cmd)
	}
}

func createApplication(configFilePath string, additionalFiles []string) *fx.App {
	return fx.New(
		fx.NopLogger,
		// Defaults
		providers.Configuration(configFilePath, additionalFiles...),
		providers.JWTExtractor(),
		providers.Logger(),
		providers.HttpClientBuilder(), // custom http client with interceptors
		// optional providers
		optionalProviders(),
		// This application/service
		thisServiceProviders(),
		// Invoke everything
		providers.InvokeService(), // start
	)
}

func thisServiceProviders() fx.Option {
	return fx.Provide(
		services.CreateWorkshopService,
		services.CreateSubWorkshopService,
		controllers.CreateWorkshopController,
		controllers.CreateSubWorkshopController,
		db.CreateCarDB,
		validations.CreateWorkshopValidations,
		validations.CreateSubWorkshopValidations,
	)
}

func optionalProviders() fx.Option {
	return fx.Provide(
		handlers.InternalDebugHandlers,
		handlers.InternalProfileHandlerFuncs,
		handlers.SelfHandlers,
	)
}
