package handler

import (
	"MyNewProject/NewBackEnd/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.RouterGroup) {
	router.Use(middleware.DBConnectionMiddleware)
}
