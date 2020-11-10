package middleware

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"time"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// 日志记录到文件
func LoggerToFile() gin.HandlerFunc {

	//日志文件
	fileName := viper.GetString("log.access_name")

	//写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

	if err != nil {
		if os.IsNotExist(err) {
			src, err = os.Create(fileName)
			if err != nil {
				fmt.Println("err", err)
			}
		} else {
			fmt.Println("err", err)
		}
	}

	//实例化
	logger := logrus.New()

	//设置输出
	logger.Out = src

	logger.Formatter = &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	}

	return func(c *gin.Context) {
		var bodyBytes []byte // 我们需要的body内容

		// 从原有Request.Body读取
		bodyBytes, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			println("failed get body")
		}

		// 新建缓冲区并替换原有Request.body
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		latencyTime := endTime.Sub(startTime)

		request := "-"
		response := "-"
		if c.Request.Method == "POST" {
			request = string(bodyBytes)
			response = blw.body.String()
		}

		logger.WithFields(logrus.Fields{
			"remote":       c.ClientIP(),
			"local":        c.Request.RemoteAddr,
			"method":       c.Request.Method,
			"path":         c.Request.RequestURI,
			"code":         c.Writer.Status(),
			"status":       c.Writer.Header().Get("status-code"),
			"req_size":     c.Request.ContentLength,
			"res_size":     c.Writer.Size(),
			"traffic":      c.Request.ContentLength + int64(c.Writer.Size()),
			"request_time": float64(endTime.UnixNano())/1e6 - float64(startTime.UnixNano())/1e6,
			"duration":     latencyTime,
			"referer":      c.Request.Referer(),
			"agent":        c.Request.UserAgent(),
			"forwoad":      "-",
			"request":      request,
			"response":     response,
			"uuid":         c.Request.Header.Get("X-Request-Id"),
			"app_id":       c.Request.Header.Get("appid"),
		}).Info()
	}
}
