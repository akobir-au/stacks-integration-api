package karpenter_route

import (
	"github.com/gin-gonic/gin"
	"github.com/akobir-au/stacks_integration_api/internal/entity"
	"github.com/akobir-au/stacks_integration_api/internal/entity/intfaces"
	"github.com/akobir-au/stacks_integration_api/pkg/logger"
	"net/http"
)

type KarpenterRoute struct {
	u intfaces.IntegrationUsecase
	l logger.Interface
}

func NewKarpenterRoute(handler *gin.RouterGroup, t intfaces.IntegrationUsecase, l logger.Interface) {
	r := &KarpenterRoute{t, l}

	h := handler.Group("/karpenter")
	{
		h.GET("/", r.deployment)

	}
}

// @Summary     Fetch karperter deployment
// @Description Show details of the karperter deployment
// @ID          Karpenter Deployment
// @Produce     json
// @Router      /karpenter/ [get]
func (route *KarpenterRoute) deployment(ctx *gin.Context) {
	deployment, err := route.u.GetDeployment(ctx, "karpenter", "karpenter")
	if err != nil {
		route.l.Error(err, "http - v1 - getting single blog")
		ctx.JSON(entity.GetStatusCode(err), entity.ErrorCodeResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, deployment)
}
