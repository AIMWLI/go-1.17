package timeout

import (
	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
	"go-gin/pkg/setting"
	"net/http"
	"time"
)

func timeoutResponse(context *gin.Context) {
	context.JSON(http.StatusGatewayTimeout, gin.H{
		"code": http.StatusGatewayTimeout,
		"msg":  "timeout",
	})
}

// TimeoutMiddleware
// @Description: 定义请求超时中间件
// @receiver nil
// @param nil
// @return gin.HandlerFunc
func TimeoutMiddleware() gin.HandlerFunc {
	withTimeoutOption := timeout.WithTimeout(setting.ServerSetting.Timeout * time.Second)
	withTimeoutResponse := timeout.WithResponse(timeoutResponse)
	withTimeoutHandler := timeout.WithHandler(func(context *gin.Context) {
		context.Next()
	})
	handlerFunc := timeout.New(withTimeoutOption, withTimeoutHandler, withTimeoutResponse)
	return handlerFunc
}
