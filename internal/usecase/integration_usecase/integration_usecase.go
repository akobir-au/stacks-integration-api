package integration_usecase

import (
	"context"

        corev1 "k8s.io/api/core/v1"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"github.com/akobir-au/stacks_integration_api/internal/entity"
)

// GetDeployment getting a single deployment by name and namespace
func (usecase *IntegrationUseCase) GetDeployment(ctx context.Context, name, namespace string) (*appsv1.Deployment, error) {
        deployment, err := usecase.client.AppsV1().Deployments(namespace).Get(ctx, name, metav1.GetOptions{})

	if err != nil {
		return nil, entity.CreateError(entity.ErrNotFound.Error(), err.Error())
	}

	return deployment, nil
}

// CreatePod creating a pod
func (usecase *IntegrationUseCase) CreatePod(ctx context.Context, pod *corev1.Pod) (*corev1.Pod, error) {
	pod, err := usecase.client.CoreV1().Pods(pod.Namespace).Create(ctx, pod, metav1.CreateOptions{})

	if err != nil {
		return nil, entity.CreateError(entity.ErrBadRequest.Error(), err.Error())
	}

	return pod, nil
}
