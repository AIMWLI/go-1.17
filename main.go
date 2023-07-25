package main

import (
	gin "github.com/gin-gonic/gin"
	"go-gin/middleware/handleLog"
	"go-gin/pkg/setting"
	"log"
	"net/http"
	"strconv"
)

func init() {
	setting.Setup()

}

func main() {

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
		context.JSON(http.StatusOK, gin.H{
			"msg": "handle log from console",
		})
	})
	err := r.Run(":" + strconv.Itoa(setting.ServerSetting.HttpPort))
	if err != nil {
		log.Fatalf("main#Run err: %v", err)
		return
	}
}
