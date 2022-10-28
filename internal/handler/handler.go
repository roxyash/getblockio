package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/roxyash/getblock/internal/service"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	_ "github.com/roxyash/getblock/docs"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	sales := router.Group("/getblockio")
	{
		sales.POST("/:id", h.checkBalance)
	}

	return router
}
