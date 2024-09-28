package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"github.com/tech/core/domain"
	"github.com/tech/core/persistence/user"
	"github.com/tech/util"
	"net/http"
	"strings"
	"time"
)

func CustomerAuthMiddleware(userPersistence user.YpUser) gin.HandlerFunc {
	functionName := "CustomerAuthMiddleware"
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")

		response := domain.Response{Code: "459", Msg: "Session has expired"}

		parts := strings.Split(authHeader, "Bearer ")
		if len(parts) != 2 {
			fmt.Println(functionName, "invalid header")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		getData, err := util.DecodeToken(parts[1])
		if err != nil {
			fmt.Println(functionName, "token validation failed ", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		auth := cast.ToString(getData["auth"])

		userExpert, err := CheckTokenExpire(cast.ToString(getData["userExpert"]))
		if err != nil || userExpert {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenExpert, err := CheckTokenExpire(cast.ToString(getData["tokenExpert"]))
		if err != nil || tokenExpert {
			fmt.Println(functionName, "token validation failed ", err)
			c.AbortWithStatusJSON(http.StatusOK, response)
			return
		}

		// get user Details using auth
		userDetails, err := userPersistence.GetYpUserByAuth(auth)
		if err != nil {
			fmt.Println(functionName, err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		c.Set("customer", userDetails)
		c.Next()

	}
}

func CheckTokenExpire(date string) (isExpired bool, err error) {
	isExpired = true
	// Define the format of the time string
	timeLayout := "2006-01-02 15:04:05.999999 -0700 MST"

	// Parse the target time from the string
	targetTime, err := time.Parse(timeLayout, date)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return
	}

	// Get the current time
	currentTime := time.Now()

	// Check if the current time is before the target time
	if currentTime.Before(targetTime) {
		isExpired = false
	}
	return isExpired, nil
}
