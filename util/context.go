package util

import (
	"context"
	"github.com/gin-gonic/gin"
)

// SetContext ... set all the context from gin context
func SetContext(c *gin.Context) context.Context {
	user := c.Keys["customer"]
	ctx := context.WithValue(context.Background(), "user", user)
	return ctx
}
