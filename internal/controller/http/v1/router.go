package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/akobir-au/stacks_integration_api/internal/controller/http/v1/karpenter_route"
	"github.com/akobir-au/stacks_integration_api/internal/entity/intfaces"
	"github.com/akobir-au/stacks_integration_api/pkg/logger"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	// Swagger docs -.
	_ "github.com/akobir-au/stacks_integration_api/docs"
)

// NewRouter
// Swagger spec:
// @title       Stacks Integration API
// @description API performing integration tests against Stacks 
// @version     1.0
// @host        localhost:8080
// @BasePath    /api/v1
// @securityDefinitions.basic BasicAuth
// @in header
// @name Authorization
func NewRouter(handler *gin.Engine, l logger.Interface, u intfaces.IntegrationUsecase) {
	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	// Swagger ui router group with basic auth
	doc := handler.Group("/swagger", gin.BasicAuth(gin.Accounts{
		"admin": "admin",
	}))

	// Creating a swaggo instance
	swaggerHandler := ginSwagger.WrapHandler(swaggerFiles.Handler)

	doc.GET("/*any", swaggerHandler)

	// K8s health probe
	handler.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Stacks Integration API is healthy.")
	})

	// Handle page not found
	handler.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "The requested page is not found."})
	})

	// Prometheus metrics
	handler.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Routers
	g := handler.Group("/api/v1")

	{
		karpenter_route.NewKarpenterRoute(g, u, l)
	}
}
