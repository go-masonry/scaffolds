package mortar

import (
	"github.com/go-masonry/bprometheus"
	"github.com/go-masonry/mortar/interfaces/cfg"
	"github.com/go-masonry/mortar/interfaces/monitor"
	mortarProject "github.com/go-masonry/mortar/mortar"
	"github.com/go-masonry/mortar/providers"
	"go.uber.org/fx"
)

// MonitoringFxOption registers prometheus
func MonitoringFxOption() fx.Option {
	return fx.Options(
		providers.MonitorFxOption(),                     // Mortar Metrics interface
		providers.MonitorGRPCInterceptorFxOption(),      // measure every gRPC call duration
		bprometheus.PrometheusInternalHandlerFxOption(), // expose ":[internal port]/metrics" for Prometheus service
		fx.Provide(PrometheusBuilder),                   // Create monitor.Builder using Prometheus brick
	)
}

// PrometheusBuilder returns a monitor.Builder that is implemented by Prometheus
func PrometheusBuilder(cfg cfg.Config) monitor.Builder {
	name := cfg.Get(mortarProject.Name).String()
	return bprometheus.Builder().SetNamespace(name)
}
