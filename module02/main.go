package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"strings"
)

/*
编写一个 HTTP 服务器，大家视个人不同情况决定完成到哪个环节，但尽量把 1 都做完：

1. 接收客户端 request，并将 request 中带的 header 写入 response header
2. 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
3. Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
4. 当访问 localhost/healthz 时，应返回 200
*/
var logger = log.Default()

func main() {
	port := os.Getenv("port")
	if port == "" {
		port = "8080"
	}
	logger := log.Default()
	logger.Printf("启动，端口:%s", port)
	engine := gin.Default()
	engine.Use(getLogInfo())
	engine.GET("/healthz", func(context *gin.Context) {
		context.JSON(200, map[string]string{"code": "200"})
	})
	engine.GET("sayHello", func(context *gin.Context) {
		context.JSON(200, map[string]string{"say": "hello"})
	})
	engine.Run(":" + port)
}

func getLogInfo() gin.HandlerFunc {
	return func(context *gin.Context) {
		logger.Print("中间件开始执行了")
		respHeader := context.Writer.Header()
		// 请求头写入响应头
		for key, val := range context.Request.Header {
			respHeader[key] = val
		}

		// 环境变量写入响应头
		envs := os.Environ()
		for index := range envs {
			s := envs[index]
			log.Println(s)
			split := strings.Split(s, "=")
			respHeader[split[0]] = []string{split[1]}
		}
		ip := context.RemoteIP()
		context.Next()
		status := context.Writer.Status()

		logger.Printf("请求客户端ip:%v, http状态码:%v\n", ip, status)
	}
}
