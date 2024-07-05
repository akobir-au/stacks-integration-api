package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/akobir-au/stacks_integration_api/config"
	"github.com/akobir-au/stacks_integration_api/internal/controller/http/v1"
	"github.com/akobir-au/stacks_integration_api/internal/usecase/integration_usecase"
	"github.com/akobir-au/stacks_integration_api/pkg/httpserver"
	"github.com/akobir-au/stacks_integration_api/pkg/logger"
	"github.com/akobir-au/stacks_integration_api/pkg/k8sclient"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	// HTTP Server
	handler := gin.Default()

	// K8s client
	k8sClient, err := k8sclient.New(cfg)

	if err != nil {
		fmt.Errorf("failed to create k8s client: %w", err)
	}

	integrationUsecase := integration_usecase.NewIntegrationUseCase(k8sClient, cfg)

	v1.NewRouter(handler, l, integrationUsecase)

	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	// Interrupt handling
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app:Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app:Run - httpServer.Notify: %w", err))
	}

	// Shutdown http server
	err = httpServer.Shutdown()

	if err != nil {
		l.Error(fmt.Errorf("app:Run - httpServer.Shutdown: %w", err))
	}
}
