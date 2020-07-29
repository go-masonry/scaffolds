package providers

import (
	"github.com/go-masonry/bricks/monitor"
	"github.com/go-masonry/bricks/monitor/bdatadog"
	"github.com/go-masonry/mortar/constructors"
	"go.uber.org/fx"
)

func Metrics() fx.Option {
	return fx.Options(
		fx.Provide(monitorBuilder),

		// custom context extractors for metrics should be added here

		fx.Provide(constructors.DefaultMonitor),
	)
}

func monitorBuilder() monitor.Builder {
	return bdatadog.Builder()
}
