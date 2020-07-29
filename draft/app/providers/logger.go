package providers

import (
	"github.com/go-masonry/bricks/cfg"
	"github.com/go-masonry/bricks/log"
	"github.com/go-masonry/bricks/log/bzerolog"
	"github.com/go-masonry/mortar/constructors"
	"github.com/go-masonry/mortar/self"
	"go.uber.org/fx"
	"os"
)

const (
	framesToSkip = 0

	application = "app"
	hostname    = "host"
	gitCommit   = "git"
)

func Logger() fx.Option {
	return fx.Options(
		fx.Provide(loggerBuilder),
		fx.Provide(constructors.DefaultLogger),
	)
}

func loggerBuilder(config cfg.Config) log.Builder {
	appName := config.Get(self.Name).String() // empty string is just fine

	builder := bzerolog.
		Builder().
		AddStaticFields(selfStaticFields(appName)).
		// You can add explicit context extractors here or use the implicit fx.Group used by `go-masonry/mortar/constructors/logger.go`
		// AddContextExtractors().
		IncludeCallerAndSkipFrames(framesToSkip)
	if config.Get(self.LoggerWriterConsole).Bool() {
		builder = builder.SetWriter(bzerolog.ConsoleWriter(os.Stderr))
	}
	return builder
}

func selfStaticFields(name string) map[string]interface{} {
	output := make(map[string]interface{})
	info := self.GetBuildInformation()
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
