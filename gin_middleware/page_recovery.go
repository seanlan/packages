package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"runtime"
)

// 页面错误恢复
func Page500Middleware(c *gin.Context) {
	defer func() {
		var err error
		if e := recover(); e != nil {
			switch e := e.(type) {
			case error:
				err = e
			default:
				err = fmt.Errorf("%v", e)
			}
			var buf []string
			// 只有在Debug模式下输出stack信息
			if gin.Mode() == gin.DebugMode {
				stack := make([]uintptr, 50)
				length := runtime.Callers(5, stack[:])
				frames := runtime.CallersFrames(stack[:length])
				for {
					frame, more := frames.Next()
					buf = append(buf, fmt.Sprintf("%s:%d %s\n", frame.File, frame.Line, frame.Function))
					if !more {
						break
					}
				}
			}
			c.HTML(500, "500.html", gin.H{
				"title": "Error",
				"err":   err,
				"stack": buf,
			})
		}
	}()
	c.Next()
}
