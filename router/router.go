package router

import (
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

//初始化Router
func Setup(mode string) {
	//设置启动模式
	gin.SetMode(mode)
	Router = gin.Default()
	Router.LoadHTMLGlob("templates/*.html")
	Router.Static("/static", "static")
}

func Run(addr string) {
	Router.Run(addr)
}
