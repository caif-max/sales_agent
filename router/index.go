package router

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"sales_agent/common/config"
	"sales_agent/common/log"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// 启动http服务
	httpLog := config.GetConf("log.dir") + "/http.log"
	f, _ := os.Create(httpLog)
	defer f.Close()

	r := gin.Default()
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r.Use(logRequestMiddleware())
	route := r.Group("/api")

	route.GET("/healthcheck", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "consul health check ok .")
	})
	// route.POST("/sendMessageToUser", handler.SendMessageToUser)

	return r
}

// 添加中间件函数来记录请求信息
func logRequestMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取客户端IP
		clientIP := c.ClientIP()
		// 获取请求URL
		requestURL := c.Request.URL.String()

		// 读取原始请求数据
		rawData, _ := c.GetRawData()
		// 由于读取了raw data，需要重新设置body供后续中间件使用
		c.Request.Body = io.NopCloser(bytes.NewBuffer(rawData))

		// 获取所有请求参数
		params := c.Request.URL.Query()

		// 记录请求信息
		log.GetLogger().Infof("Client IP: %s, URL: %s, Params: %v, Raw Data: %s",
			clientIP, requestURL, params, string(rawData))

		c.Next()
	}
}
