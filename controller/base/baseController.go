package base

import (
	"apiDemoProject/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/shima-park/agollo"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"strings"
)

// @Summary 健康检查
// @Tags 基础接口
// @Produce  plain
// @Success 200 {string} string 成功后返回值
// @Router /api/v6/data/health/check/get [get]
func HealthCheck(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}

// @Summary 测试配置更新
// @Tags 基础接口
// @Produce  plain
// @Success 200 {string} string 成功后返回值
// @Router /api/v6/data/config/test/get [post]
func TestAgollo(c *gin.Context) {
	c.String(http.StatusOK, agollo.Get("timeout"))
}

// @Summary Redis操作
// @Tags 基础接口
// @Accept  json
// @Produce  json
// @Param command body model.RedisCommandList true "操作命令列表"
// @Success 200 {object} model.Response 成功后返回值
// @Router /api/v6/data/redis_command/post [post]
func RedisCommand(c *gin.Context) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", viper.GetString("redis.host"), viper.GetInt("redis.port")),
		Password: viper.GetString("redis.password"),
		DB:       0,
	})

	pong, err := redisClient.Ping().Result()
	if err != nil {
		logrus.Panic("Connecting redis error")
	} else {
		logrus.Printf("Redis response is %s", pong)
	}

	var redisCommandList model.RedisCommandList
	if err := c.ShouldBindJSON(&redisCommandList); err != nil {
		data := make(map[string]interface{})
		data["result"] = "fail"
		data["error"] = err.Error()
		c.JSON(http.StatusOK, model.Response{
			Code:    0,
			Message: "success",
			Data:    data,
		})
		return
	}

	for _, redisCommand := range redisCommandList.CommandList {
		method := redisCommand.Method
		if method != "" {
			method = strings.ToUpper(method)
		}
		switch method {
		case "HMSET":
			redisClient.HMSet(redisCommand.Name, redisCommand.ArgList)
		case "SET":
			redisClient.Set(redisCommand.Name, redisCommand.Args, 0)
		case "DEL":
			redisClient.Del(redisCommand.Name)
		}
	}

	data := make(map[string]interface{})
	data["result"] = "OK"
	c.JSON(http.StatusOK, model.Response{
		Code:    0,
		Message: "success",
		Data:    data,
	})
}
