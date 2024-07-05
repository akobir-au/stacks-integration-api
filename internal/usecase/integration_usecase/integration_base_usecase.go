package integration_usecase

import (
	"github.com/akobir-au/stacks_integration_api/config"
	"github.com/akobir-au/stacks_integration_api/internal/entity/intfaces"
        "k8s.io/client-go/kubernetes"
)

type IntegrationUseCase struct {
	client *kubernetes.Clientset
	config *config.Config
}

func NewIntegrationUseCase(client *kubernetes.Clientset, config *config.Config) intfaces.IntegrationUsecase {
	return &IntegrationUseCase{
		client: client,
		config: config,
	}
}
