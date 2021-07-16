package main

import (
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/seanlan/packages/gin_router"
	"github.com/seanlan/packages/logging"
	"go.uber.org/zap"
	"time"
)

type TestApiArgs struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password int    `form:"password" json:"password" binding:"required"`
	Age      int    `form:"age" json:"age"`
}

// TestApi 测试接口
func TestApi(c *gin.Context) {
	// 接口参数绑定
	var args TestApiArgs
	err := gin_router.RequestParser(&args, c)
	// 接口参数错误
	if err != nil {
		return
	}
	var a,b = 10,0
	zap.S().Info(a/b)
	gin_router.JsonExit(c,0,"SUCCESS", map[string]interface{}{"args": args})
}

func main() {
	logging.Setup(false,"test")
	gin_router.Setup("release")
	gin_router.Router.Use(
		ginzap.Ginzap(zap.L(), time.StampMilli, true),
		ginzap.RecoveryWithZap(zap.L(), true))
	apiGroupV1 := gin_router.Router.Group("/api/v1")
	apiGroupV1.POST("/test/test",TestApi)
	gin_router.Run(":8080")
}