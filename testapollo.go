
package main

import (
	"fmt"
	"github.com/shima-park/agollo"
	"os"
)

func main() {
	// 通过默认根目录下的app.properties初始化agollo
	//err := agollo.InitWithDefaultConfigFile(
	//	agollo.WithLogger(agollo.NewLogger(agollo.LoggerWriter(os.Stdout))), // 打印日志信息
	//	agollo.PreloadNamespaces("application"),                          // 预先加载的namespace列表，如果是通过配置启动，会在app.properties配置的基础上追加
	//	agollo.AutoFetchOnCacheMiss(),                                       // 在配置未找到时，去apollo的带缓存的获取配置接口，获取配置
	//	agollo.FailTolerantOnBackupExists(),                                 // 在连接apollo失败时，如果在配置的目录下存在.agollo备份配置，会读取备份在服务器无法连接的情况下
	//)


	err := agollo.Init(
		"129.211.158.165:8080",
		"golang_test",
		agollo.WithLogger(agollo.NewLogger(agollo.LoggerWriter(os.Stdout))),
		agollo.AutoFetchOnCacheMiss(),
		agollo.FailTolerantOnBackupExists(),
	)
	if err != nil {
		panic(err)
	}

	// 获取默认配置中cluster=default namespace=application key=Name的值
	fmt.Println("timeout:", agollo.Get("timeout"))

	// 获取namespace下的所有配置项
	fmt.Println("Configuration of the namespace:", agollo.GetNameSpace("application"))

	fmt.Println(agollo.Get("timeout", agollo.WithDefault("10"), agollo.WithNamespace("golang_test_dev")))

	// 如果想监听并同步服务器配置变化，启动apollo长轮训
	// 返回一个期间发生错误的error channel,按照需要去处理
	errorCh := agollo.Start()

	// 监听apollo配置更改事件
	// 返回namespace和其变化前后的配置,以及可能出现的error
	watchCh := agollo.Watch()

	stop := make(chan bool)
	watchNamespace := "application"
	watchNSCh := agollo.WatchNamespace(watchNamespace, stop)

	go func() {
		for {
			select {
			case err := <-errorCh:
				fmt.Println("Error:", err)
			case resp := <-watchCh:
				fmt.Println("Watch Apollo:", resp)
				fmt.Println("timeout:", agollo.Get("timeout"))
			case resp := <-watchNSCh:
				fmt.Println("Watch Namespace", watchNamespace, resp)
			}
		}
	}()

	select {}

	agollo.Stop()
}