package middleware

import (
	"apiDemoProject/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginRequiredMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Request.Header.Get("X-Username")
		token := c.Request.Header.Get("Token")

		if token == "abc" || username != "" {
			c.Header("status-code", "0")
			c.Next()
			return
		}
		c.Header("status-code", "100013")
		c.JSON(http.StatusOK, model.Response{
			Code:    100013,
			Message: "login required",
		})
		c.Abort()
		return
	}
}
