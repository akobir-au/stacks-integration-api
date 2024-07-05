package k8sclient

import (
	"log"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
        "github.com/akobir-au/stacks_integration_api/config"
)

func New(cfg *config.Config) (*kubernetes.Clientset, error) {
	// Create in-cluster or external config
	config, err := rest.InClusterConfig()

        if len(cfg.App.KubeConfig) != 0 {
		config, err = clientcmd.BuildConfigFromFlags("", cfg.App.KubeConfig)
	}

	if err != nil {
		log.Fatal("Failed to create cluster config", err)
	}

	// Create clientset
	clientset, err := kubernetes.NewForConfig(config)

	if err != nil {
		log.Fatal("Failed to create clientset", err)
	}

	return clientset, nil
}
