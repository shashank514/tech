package login

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tech/core/domain"
	"github.com/tech/core/service/login/driver"
	"github.com/tech/util"
	"net/http"
)

func UserOtp(svc driver.LoginService) gin.HandlerFunc {
	var exeCtx context.Context
	return func(c *gin.Context) {

		requestBody := &domain.OtpRequest{}
		if err := c.Bind(requestBody); err != nil {
			fmt.Println(err)
			response := domain.Response{Code: "404", Msg: "invalid request body"}
			c.JSON(http.StatusOK, response)
			return
		}

		response := domain.Response{}

		switch requestBody.Action {
		case "generate":
			response = svc.GenerateOtpForUser(exeCtx, requestBody)
		case "submit":
			response = svc.SubmitOtpForUser(exeCtx, requestBody)
		default:
			response.Code = "600"
		}

		c.JSON(http.StatusOK, response)
	}
}

func GenerateToken(svc driver.LoginService) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := &domain.NewToken{}
		if err := c.Bind(requestBody); err != nil {
			fmt.Println(err)
			response := domain.Response{Code: "404", Msg: "invalid request body"}
			c.JSON(http.StatusOK, response)
			return
		}

		response := svc.GetNewToken(requestBody)
		c.JSON(http.StatusOK, response)
	}
}

func UserDetails(svc driver.LoginService) gin.HandlerFunc {
	functionName := "UserDetails"
	var exeCtx context.Context
	return func(c *gin.Context) {
		exeCtx = util.SetContext(c)
		types := string(c.Query("type"))

		user := c.Keys["customer"]
		if user == nil {
			fmt.Println(functionName, "invalid credentials")
			response := domain.Response{Code: "459", Msg: "Session has expired"}
			c.JSON(http.StatusUnauthorized, response)
			return
		}

		loggedUser := user.(*domain.User)

		response := domain.Response{}
		switch types {
		case "personalDetails":
			requestBody := &domain.PersonalDetails{}
			if err := c.Bind(requestBody); err != nil {
				fmt.Println(err)
				response := domain.Response{Code: "404", Msg: "invalid request body"}
				c.JSON(http.StatusOK, response)
				return
			}
			response = svc.GetUserPersonalDetails(exeCtx, loggedUser, requestBody)
		case "address":
			requestBody := &domain.AddressDetails{}
			if err := c.Bind(requestBody); err != nil {
				fmt.Println(err)
				response := domain.Response{Code: "404", Msg: "invalid request body"}
				c.JSON(http.StatusOK, response)
				return
			}
			response = svc.GetUserAddressDetails(exeCtx, loggedUser, requestBody)
		case "status":
			response = svc.GetUserDetails(exeCtx, loggedUser)
		default:
			response.Code = "400"
			response.Msg = "invalid request body"
		}

		c.JSON(http.StatusOK, response)
	}
}
