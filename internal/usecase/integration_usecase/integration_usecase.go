package integration_usecase

import (
	"context"

	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"github.com/akobir-au/stacks_integration_api/internal/entity"
)

// GetDeployment getting a single deployment by name and namespace
func (usecase *IntegrationUseCase) GetDeployment(ctx context.Context, name, namespace string) (*appsv1.Deployment, error) {
        deployment, err := usecase.client.AppsV1().Deployments(namespace).Get(ctx, name, metav1.GetOptions{})
        //deployment, err := usecase.client.AppsV1().Deployments(namespace).Get(context.TODO(), name, metav1.GetOptions{})

	if err != nil {
		return nil, entity.CreateError(entity.ErrNotFound.Error(), err.Error())
	}

	return deployment, nil
}
