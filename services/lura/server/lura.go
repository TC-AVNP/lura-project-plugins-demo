package server

import (
	"context"
	"os"

	gonicgin "github.com/gin-gonic/gin"
	"github.com/luraproject/lura/config"
	"github.com/luraproject/lura/logging"
	"github.com/luraproject/lura/proxy"
	"github.com/luraproject/lura/router"
	"github.com/luraproject/lura/router/gin"
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

		serviceConfig.Port = port
		serviceConfig.Debug = true

		luraLogger, _ := logging.NewLogger(gonicgin.DebugMode, os.Stdout, "[LURA]")
		proxyFactory := proxy.DefaultFactory(luraLogger)

		routerFactory := gin.NewFactory(gin.Config{
			Engine:         gonicgin.Default(),
			Middlewares:    []gonicgin.HandlerFunc{},
			ProxyFactory:   proxyFactory,
			HandlerFactory: gin.EndpointHandler,
			Logger:         luraLogger,
			RunServer:      router.RunServer,
		}).NewWithContext(spawnCtx)

		routerFactory.Run(serviceConfig)
	}(spawnCtx)

	return cancel, nil
}

func (l LuraInstance) Stop() error {
	return nil
}
