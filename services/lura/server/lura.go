package server

import (
	"context"
	"net/http"
	"os"

	"github.com/luraproject/lura/v2/config"
	"github.com/luraproject/lura/v2/logging"
	"github.com/luraproject/lura/v2/proxy"
	"github.com/luraproject/lura/v2/router/gin"
	server "github.com/luraproject/lura/v2/transport/http/server/plugin"
)

type LuraInstance struct {
}

func NewLuraInstance() *LuraInstance {
	return &LuraInstance{}
}

func (l *LuraInstance) Spawn(port int) (context.CancelFunc, error) {
	spawnCtx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		configFile := "config.json"

		parser := config.NewParser()
		serviceConfig, _ := parser.Parse(configFile)

		logger, _ := logging.NewLogger("logLevel", os.Stdout, "[LURA]")

		pluginLoader := pluginLoader{}
		pluginLoader.Load(serviceConfig.Plugin.Folder, serviceConfig.Plugin.Pattern, logger)

		routerFactory := gin.DefaultFactory(proxy.DefaultFactory(logger), logger)

		routerFactory.New().Run(serviceConfig)

	}(spawnCtx)

	return cancel, nil
}

func (l LuraInstance) Stop() error {
	return nil
}

type RunServer func(context.Context, config.ServiceConfig, http.Handler) error

func newRunServer(l logging.Logger, next gin.RunServerFunc) RunServer {
	return RunServer(server.RunServer(next))
}
