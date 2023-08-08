package routers

import (
	"github.com/gin-gonic/gin"
	"go-gin/middleware/handleLog"
	"go-gin/pkg/setting"
	"net/http"
)

func InitRouter() *gin.Engine {
	gin.SetMode(setting.ServerSetting.RunMode)
	r := gin.Default()
	//r.Use(timeout.TimeoutMiddleware())
	r.Use(handleLog.HandleLog())
	r.GET("ping", func(context *gin.Context) {
		//模拟响应超时
		//time.Sleep(time.Second * 2)
		context.JSON(http.StatusOK, gin.H{
			"msg": "pong",
		})
	})
	r.POST("/log", func(context *gin.Context) {
		//模拟计算响应事件
		//time.Sleep(time.Second * 2)
		context.JSON(http.StatusOK, gin.H{
			"msg": "handle log from console",
		})
	})
	return r
}
