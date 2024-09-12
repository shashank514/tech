package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/tech/cloud/message"
	"github.com/tech/core/persistence/user"
	"github.com/tech/core/service/login"
	login2 "github.com/tech/handler/login"
	"github.com/tech/middleware"
)

func SetupRoutes(router *gin.RouterGroup) {

	router.Use(middleware.DBConnectionWithEnvMiddleware)

	messages := message.NewMessage()
	userPersist := *user.NewUser()

	loginSvc := login.NewLoginDetails(messages, userPersist)

	otp := login2.UserOtp(loginSvc)

	subRouter := router.Group("/login")
	{
		subRouter.POST("/otp", otp)
	}
}
