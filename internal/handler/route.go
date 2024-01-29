package handler

import "github.com/gin-gonic/gin"

func InitRoutes(h HandlerInterface) *gin.Engine {
	router := gin.Default()
	router.GET("/currency", h.GetCurrency)
	return router
}
