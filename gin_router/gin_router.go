package gin_router

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

var Router *gin.Engine

// Setup 初始化Router
func Setup(mode string) {
	//设置启动模式
	gin.SetMode(mode)
	Router = gin.New()
}

func Run(addr string) {
	Router.Run(addr)
}

type JsonResponse struct {
	Error        int         `json:"error,required"`
	ErrorMessage string      `json:"error_msg,required"`
	Data         interface{} `json:"data",omitempty`
}

// JsonExit 返回JSON数据并退出当前HTTP执行函数。
func JsonExit(c *gin.Context, err int, msg string, data ...interface{}) {
	var result = new(JsonResponse)
	result.Error = err
	result.ErrorMessage = msg
	if len(data) > 0 {
		result.Data = data[0]
	}
	c.JSON(http.StatusOK, result)
	c.Done()
}

// RequestParser 参数的绑定与解析
func RequestParser(args interface{}, c *gin.Context) error {
	var err error
	contentType := c.Request.Header.Get("Content-Type")
	switch contentType {
	case "application/json":
		err = c.BindJSON(args)
	case "application/x-www-form-urlencoded":
		err = c.MustBindWith(args, binding.Form)
	default:
		err = c.Bind(args)
	}
	if err != nil {
		// 返回错误信息
		JsonExit(c,400, err.Error())
	}
	return err
}