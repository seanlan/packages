package main

import (
	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-contrib/requestid"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/seanlan/packages/config"
	"github.com/seanlan/packages/gin_router"
	"github.com/seanlan/packages/gredis"
	"github.com/seanlan/packages/logging"
	"go.uber.org/zap"
	"time"
)

type TestApiReq struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password int    `form:"password" json:"password" binding:"required"`
	Age      int    `form:"age" json:"age,omitempty"`
}

// TestApi 测试接口
func TestApi(c *gin.Context) {
	// 接口参数绑定
	var args TestApiReq
	err := gin_router.RequestParser(&args, c)
	// 接口参数错误
	if err != nil {
		return
	}
	//var a,b = 10,0
	//zap.S().Info(a/b)
	gin_router.JsonExit(
		c,0,"SUCCESS",
		map[string]interface{}{
			"args": args,
		})
}

func main() {
	config.Setup("example/conf.d/conf.yaml")
	gredis.Setup(config.GetString("redis"))
	logging.Setup(false,"test")
	store := persistence.NewRedisCacheWithPool(gredis.RedisPool, time.Minute * 10)
	gin_router.Setup("release")
	gin_router.Router.Use(
		ginzap.Ginzap(zap.L(), time.StampMilli, true),
		ginzap.RecoveryWithZap(zap.L(), true),
		requestid.New())
	apiGroupV1 := gin_router.Router.Group("/api/v1")
	apiGroupV1.POST("/test/test",cache.CachePage(store, time.Minute, TestApi))
	gin_router.Run(":8080")
}