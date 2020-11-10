package main

import (
	"apiDemoProject/controller/base"
	"apiDemoProject/controller/member"
	"apiDemoProject/docs"
	"apiDemoProject/middleware"
	"apiDemoProject/model"
	"apiDemoProject/util"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"

	"net/http"
	"os"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "c", "", "config file path")
	flag.Parse()

	var config model.Config
	if configPath != "" {
		viper.SetConfigFile(configPath)
	} else {
		viper.SetConfigName("api-demo")
		viper.AddConfigPath(".")
	}

	viper.SetDefault("server.host", "0.0.0.0")
	viper.SetDefault("server.port", 8080)
	viper.SetDefault("redis.host", "127.0.0.1")
	viper.SetDefault("redis.port", 6379)
	viper.SetDefault("log.level", "info")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	_ = viper.Unmarshal(&config)

	_ = SetConfig(logrus.StandardLogger(), config)
}

func SetConfig(logger *logrus.Logger, conf model.Config) (err error) {
	src := os.Stdout
	if conf.Log.CustomLogName != "" {
		//写入文件
		src, err = os.OpenFile(conf.Log.CustomLogName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

		if err != nil {
			if os.IsNotExist(err) {
				src, err = os.Create(conf.Log.CustomLogName)
				if err != nil {
					fmt.Println("err", err)
				}
			} else {
				fmt.Println("err", err)
			}
		}
	} else {
		src = os.Stdout
	}

	logger.Out = src

	logger.Formatter = &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	}

	var lvl = logrus.DebugLevel
	if lvl, err = logrus.ParseLevel(conf.Log.Level); err != nil {
		return
	} else {
		logger.Level = lvl
	}

	return
}

func main() {
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Title = "ServiceMesh Go Api Demo"
	docs.SwaggerInfo.Host = "localhost:8080"
	//
	//err := agollo.Init(
	//	viper.GetString("agollo.config_server_url"),
	//	viper.GetString("agollo.appid"),
	//	agollo.WithLogger(agollo.NewLogger(agollo.LoggerWriter(os.Stdout))),
	//	agollo.PreloadNamespaces("application"),
	//	agollo.AutoFetchOnCacheMiss(),
	//	agollo.FailTolerantOnBackupExists(),
	//)
	//if err != nil {
	//	logrus.Panic(err)
	//}

	//errorCh := agollo.Start()
	//go func() {
	//	for {
	//		select {
	//		case err := <-errorCh:
	//			logrus.Error(err)
	//		}
	//	}
	//}()

	gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()
	router := gin.New()

	router.Use(middleware.TraceHandelMiddleWare())
	router.Use(middleware.LoggerToFile())

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome Go Api Demo Server")
	})

	baseRouter := router.Group("/")
	{
		baseRouter.GET("/api/v6/data/health/check/get", base.HealthCheck)

		baseRouter.POST("/api/v6/data/redis_command/post", base.RedisCommand)

		//baseRouter.POST("/api/v6/data/config/test/get", base.TestAgollo)

		baseRouter.POST("/api/v6/data/test_call_java", func(c *gin.Context) {
			url := `http://java-rpc-sample:8080/internal/today.do`
			resp := util.RequestUtil(url, map[string]interface{}{"uid": "1"}, c)
			c.String(http.StatusOK, resp.String())
		})

	}

	memberRouter := router.Group("/api/member")
	{
		memberRouter.POST("/info/get", middleware.LoginRequiredMiddleWare(), member.GetInfo)

		memberRouter.POST("/upload/post", middleware.LoginRequiredMiddleWare(), member.Upload)

		memberRouter.POST("/list/get", middleware.LoginRequiredMiddleWare(), member.GetAll)

		memberRouter.POST("/update/post", middleware.LoginRequiredMiddleWare(), member.Update)

		memberRouter.POST("/delete/post", middleware.LoginRequiredMiddleWare(), member.Delete)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	_ = router.Run(fmt.Sprintf("%s:%d", viper.GetString("server.host"), viper.GetInt("server.port")))
}
