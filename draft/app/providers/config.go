package providers

import (
	"github.com/go-masonry/bricks/cfg"
	"github.com/go-masonry/bricks/cfg/bviper"
	"go.uber.org/fx"
)

func Configuration(configFilePath string, additionalFilePaths ...string) fx.Option {
	return fx.Provide(func() (cfg.Config, error) {
		builder := bviper.Builder().SetConfigFile(configFilePath)
		for _, extraFile := range additionalFilePaths {
			builder = builder.AddExtraConfigFile(extraFile)
		}
		return builder.Build()
	})
}
