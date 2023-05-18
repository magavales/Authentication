package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.POST("/auth", h.Auth)
		api.GET("/get-token", h.GetToken)
	}

	return router
}
