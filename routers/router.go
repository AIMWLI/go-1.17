package routers

import (
	"github.com/gin-gonic/gin"
	"go-gin/middleware/handleLog"
	"go-gin/middleware/jwt"
	"go-gin/pkg/setting"
	"go-gin/routers/api"
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
	r.GET(
		"test",
		//app.Response(http.StatusOK, e.SUCCESS, "Response"),
	)
	r.POST("/log", func(context *gin.Context) {
		//模拟计算响应事件
		//time.Sleep(time.Second * 2)
		context.JSON(http.StatusOK, gin.H{
			"msg": "handle log from console",
		})
	})
	// 注册路由
	r.GET("/auth", api.GetAuth)
	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		apiv1.POST("/save", api.Save)
		apiv1.GET("/select", api.SelectPage)
	}

	return r
}
