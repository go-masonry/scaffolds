package providers

import (
	"github.com/go-masonry/mortar/constructors"
	"go.uber.org/fx"
)

func JWTExtractor() fx.Option {
	return fx.Provide(constructors.DefaultJWTTokenExtractor)
}
