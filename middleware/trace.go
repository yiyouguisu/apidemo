package middleware

import (
	"github.com/gin-gonic/gin"
)

func TraceHandelMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		headers := []string{
			"X-Request-Id",

			"X-B3-TraceId",
			"X-B3-SpanId",
			"X-B3-ParentSpanId",
			"X-B3-Sampled",
			"X-B3-Flags",
			"X-Ot-Span-Context",

			"appid",
			"X-Username",
		}
		for _, header := range headers {
			if c.Request.Header.Get(header) != "" {
				c.Set(header, c.Request.Header.Get(header))
			}
		}
	}
}
