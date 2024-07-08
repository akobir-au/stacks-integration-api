package intfaces

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	appsv1 "k8s.io/api/apps/v1"
)

type IntegrationUsecase interface {
	GetDeployment(ctx context.Context, name, namespace string) (*appsv1.Deployment, error)
	CreatePod(ctx context.Context, pod *corev1.Pod) (*corev1.Pod, error)
}
