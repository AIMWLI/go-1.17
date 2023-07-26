package main

import (
	"context"
	gin "github.com/gin-gonic/gin"
	"go-gin/middleware/handleLog"
	"go-gin/pkg/setting"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
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
	srv := &http.Server{
		Addr:    ":" + strconv.Itoa(setting.ServerSetting.HttpPort),
		Handler: r,
	}
	//err := r.Run(":" + strconv.Itoa(setting.ServerSetting.HttpPort))

	// https://gin-gonic.com/docs/examples/graceful-restart-or-stop/
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")
}
