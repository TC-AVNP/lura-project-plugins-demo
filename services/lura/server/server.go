package server

import (
	"os"

	"github.com/luraproject/lura/v2/config"
	"github.com/luraproject/lura/v2/logging"
	"github.com/luraproject/lura/v2/proxy"
	"github.com/luraproject/lura/v2/router/gin"
)

type LuraInstance struct {
}

func NewLuraInstance() *LuraInstance {
	return &LuraInstance{}
}

func (l *LuraInstance) Start() error {
	go func() {
		configFile := "config.json"

		parser := config.NewParser()
		serviceConfig, _ := parser.Parse(configFile)

		logger, _ := logging.NewLogger("logLevel", os.Stdout, "[LURA]")

		pluginLoader := pluginLoader{}
		pluginLoader.Load(serviceConfig.Plugin.Folder, serviceConfig.Plugin.Pattern, logger)

		routerFactory := gin.DefaultFactory(proxy.DefaultFactory(logger), logger)

		routerFactory.New().Run(serviceConfig)

	}()

	return nil
}
