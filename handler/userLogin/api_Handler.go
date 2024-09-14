package userLogin

import (
	"github.com/gin-gonic/gin"
	"github.com/tech/cloud/message"
	"github.com/tech/core/persistence/user"
	"github.com/tech/core/service/login"
	login2 "github.com/tech/handler/userLogin/login"
	"github.com/tech/middleware"
)

func SetupRoutes(router *gin.RouterGroup) {

	router.Use(middleware.DBConnectionWithEnvMiddleware)

	messages := message.NewMessage()
	userPersist := *user.NewUser()

	loginSvc := login.NewLoginDetails(messages, userPersist)

	otp := login2.UserOtp(loginSvc)
	token := login2.GenerateToken(loginSvc)
	router.POST("/otp", otp)
	router.POST("/token", token)

	details := router.Group("/details")
	details.Use(middleware.DBConnectionWithEnvMiddleware)
	details.Use(middleware.CustomerAuthMiddleware(user.YpUserDetails()))

	userDetails := login2.UserDetails(loginSvc)
	details.POST("", userDetails)
}
