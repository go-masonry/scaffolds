package providers

import (
	"github.com/go-masonry/bjaeger"
	"github.com/go-masonry/bzerolog"
	"github.com/go-masonry/mortar/constructors"
	"github.com/go-masonry/mortar/interfaces/cfg"
	"github.com/go-masonry/mortar/interfaces/log"
	"github.com/go-masonry/mortar/mortar"
	"go.uber.org/fx"
	"os"
)

const (
	framesToSkip = 0

	application = "app"
	hostname    = "host"
	gitCommit   = "git"
)

func LoggerOption() fx.Option {
	return fx.Options(
		fx.Provide(loggerBuilder),
		fx.Provide(constructors.DefaultLogger),
	)
}

func loggerBuilder(config cfg.Config) log.Builder {
	appName := config.Get(mortar.Name).String() // empty string is just fine

	builder := bzerolog.
		Builder().
		AddStaticFields(selfStaticFields(appName)).
		// You can add explicit context extractors here or use the implicit fx.Group used by `go-masonry/mortar/constructors/logger.go`
		AddContextExtractors(bjaeger.TraceInfoExtractorFromContext).
		IncludeCallerAndSkipFrames(framesToSkip)
	if config.Get(mortar.LoggerWriterConsole).Bool() {
		builder = builder.SetWriter(bzerolog.ConsoleWriter(os.Stderr))
	}
	return builder
}

func selfStaticFields(name string) map[string]interface{} {
	output := make(map[string]interface{})
	info := mortar.GetBuildInformation()
	if len(name) > 0 {
		output[application] = name
	}
	if len(info.Hostname) > 0 {
		output[hostname] = info.Hostname
	}
	if len(info.GitCommit) > 0 {
		output[gitCommit] = info.GitCommit
	}
	return output
}
