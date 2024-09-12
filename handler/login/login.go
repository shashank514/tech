package login

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tech/core/domain"
	"github.com/tech/core/service/login/driver"
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
