package routers

import (
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/vancholen/chatgpt-web/config"
	"github.com/vancholen/chatgpt-web/controllers"
	"github.com/vancholen/chatgpt-web/logger"
	middlewares "github.com/vancholen/chatgpt-web/middleware"
)

var router *gin.Engine
var once sync.Once
var chatController = controllers.NewChatController()

func SetUpRoute() {
	once.Do(func() {
		router = gin.Default()
		RegisterWebRoutes(router)
	})
}

// RegisterWebRoutes 注册路由
func RegisterWebRoutes(router *gin.Engine) {
	router.Use(middlewares.Cors())
	// router.Use(middlewares.TlsHandler())
	router.GET("/", chatController.Index)
	router.POST("/completion", chatController.Completion)
}

// initTemplate 初始化HTML模板加载路径
func initTemplateDir() {
	router.LoadHTMLGlob("resources/view/*")
}

// initStaticServer 初始化静态文件处理
func initStaticServer() {
	router.StaticFS("/static", http.Dir("static"))
	router.StaticFile("logo192.png", "static/logo192.png")
	router.StaticFile("logo512.png", "static/logo512.png")
}

func StartWebServer() {
	// 注册启动所需各类参数
	SetUpRoute()
	initTemplateDir()
	initStaticServer()
	// 启动服务
	port := config.LoadConfig().Port
	portString := strconv.Itoa(port)
	err := router.Run(":" + portString)
	if err != nil {
		logger.Danger("run webserver error %s", err)
		return
	}
}
