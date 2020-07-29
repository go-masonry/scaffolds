package providers

import (
	"github.com/go-masonry/mortar/constructors/partial"
	"go.uber.org/fx"
)

func HttpClientBuilder() fx.Option {
	return fx.Provide(partial.HttpClientBuilder)
}
