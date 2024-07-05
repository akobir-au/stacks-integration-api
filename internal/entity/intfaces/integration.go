package intfaces

import (
	"context"

	appsv1 "k8s.io/api/apps/v1"
)

type IntegrationUsecase interface {
	GetDeployment(ctx context.Context, name, namespace string) (*appsv1.Deployment, error)
}
