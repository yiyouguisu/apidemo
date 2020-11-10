package util

import (
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

func RequestUtil(url string, data map[string]interface{}, c *gin.Context) *resty.Response {
	var headerNames = [8]string{
		"X-Request-Id",

		"X-B3-TraceId",
		"X-B3-SpanId",
		"X-B3-ParentSpanId",
		"X-B3-Sampled",
		"X-B3-Flags",

		"appid",
		"X-Username",
	}
	var headers = map[string]string{}
	for i := 0; i < len(headerNames); i++ {
		if c.GetHeader(headerNames[i]) != "" {
			headers[headerNames[i]] = c.GetHeader(headerNames[i])
		}
	}
	client := resty.New()
	resp, err := client.R().
		SetHeaders(headers).
		SetBody(data).
		Post(url)
	if err != nil {
		println(err)
	}
	return resp
}
