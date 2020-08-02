package providers

import (
	"github.com/go-masonry/bdatadog"
	"github.com/go-masonry/mortar/constructors"
	"github.com/go-masonry/mortar/interfaces/monitor"
	"go.uber.org/fx"
)

func MetricsOption() fx.Option {
	return fx.Options(
		fx.Provide(monitorBuilder),

		// custom context extractors for metrics should be added here

		fx.Provide(constructors.DefaultMonitor),
	)
}

func monitorBuilder() monitor.Builder {
	return bdatadog.Builder()
}
