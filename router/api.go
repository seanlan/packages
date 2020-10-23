package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

//API 成功返回数据结构
type SuccessResult struct {
	Error        int         `json:"error,required"`
	ErrorMessage string      `json:"error_msg,required"`
	Data         interface{} `json:"data"`
}

//API 错误返回数据结构
type ErrorResult struct {
	Error        int    `json:"error,required"`
	ErrorMessage string `json:"error_msg,required"`
}

//API 成功返回数据
func SuccessReturn(c *gin.Context, res interface{}) {
	result := SuccessResult{Error: 0, ErrorMessage: "SUCCESS", Data: res}
	c.JSON(http.StatusOK, result)
}

func ErrorReturn(c *gin.Context, errCode int, errMsg string) {
	c.JSON(http.StatusOK, ErrorResult{errCode, errMsg})
}

// 参数的绑定与解析
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
		c.AbortWithStatusJSON(400,
			ErrorResult{400, err.Error()},
		)
	}
	return err
}
