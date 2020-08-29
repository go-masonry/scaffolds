package mortar

import (
	"os"

	"github.com/go-masonry/bjaeger"
	"github.com/go-masonry/bzerolog"
	"github.com/go-masonry/mortar/interfaces/cfg"
	"github.com/go-masonry/mortar/interfaces/log"
	"github.com/go-masonry/mortar/mortar"
	"github.com/go-masonry/mortar/providers"
	"go.uber.org/fx"
)

// LoggerFxOption add logger
func LoggerFxOption() fx.Option {
	return fx.Options(
		fx.Provide(ZeroLogBuilder),                             // Zerolog brick
		providers.LoggerFxOption(),                             // Mortar logger wrapper with middleware support
		providers.LoggerGRPCIncomingContextExtractorFxOption(), // add different fields from Context into log entry
		bjaeger.TraceInfoContextExtractorFxOption(),            // Explicit logger context extractor to add tracing info into log entry
	)
}

// ZeroLogBuilder create Mortar Logger builder using zerolog
func ZeroLogBuilder(config cfg.Config) log.Builder {
	builder := bzerolog.Builder().IncludeCaller()
	if config.Get(mortar.LoggerWriterConsole).Bool() { // if not true JSON output will be used, better suitable for Elastic/AWS
		builder = builder.SetWriter(bzerolog.ConsoleWriter(os.Stderr))
	}
	return builder
}
