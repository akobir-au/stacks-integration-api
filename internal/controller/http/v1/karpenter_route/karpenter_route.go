package karpenter_route

import (
	"github.com/gin-gonic/gin"
	"github.com/akobir-au/stacks_integration_api/internal/entity"
	"github.com/akobir-au/stacks_integration_api/internal/entity/intfaces"
	"github.com/akobir-au/stacks_integration_api/pkg/logger"
	"net/http"

        corev1 "k8s.io/api/core/v1"
        metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type KarpenterRoute struct {
	u intfaces.IntegrationUsecase
	l logger.Interface
}

func NewKarpenterRoute(handler *gin.RouterGroup, t intfaces.IntegrationUsecase, l logger.Interface) {
	r := &KarpenterRoute{t, l}

	h := handler.Group("/karpenter")
	{
		h.GET("/scaleup", r.scaleUp)

	}
}

// @Summary     Test Karpenter Scale Up
// @Description Creates a workload that will trigger Karpenter to scale up
// @ID          Karpenter Scale Up
// @Produce     json
// @Router      /karpenter/scaleup/ [get]
func (route *KarpenterRoute) scaleUp(ctx *gin.Context) {
	req := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: "nginx",
			Namespace: "default",
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:  "nginx",
					Image: "nginx",
				},
			},
		},
	}

	pod, err := route.u.CreatePod(ctx, req)
	if err != nil {
		route.l.Error(err, "http - v1 - error creating pod")
		ctx.JSON(entity.GetStatusCode(err), entity.ErrorCodeResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, pod)
}
