package handleLog

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin/util"
	"log"
	"time"
)

// HandleLog
// @Description: Get request parameters and response results and print them
// @receiver nil
// @param nil
// @return nil
func HandleLog() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 记录请求开始时间
		start := time.Now()
		// 打印请求信息
		rawData, err := ctx.GetRawData()
		if err != nil {
			log.Fatalf("[WARN] Request err: %v\n", err)
		}
		// todo 多次读取请求 Body 的问题
		log.Printf("[INFO] Request: %s %s \n%s\n", ctx.Request.Method, ctx.Request.RequestURI, rawData)
		// 继续执行
		//ctx.Next()
		// 记录响应时间
		end := time.Now()
		stopWatch := end.Sub(start)
		//rewrite response body in middleware
		//https://github.com/gin-gonic/gin/issues/3384
		/*
			cw := &copyWriter{buf: &bytes.Buffer{}, ResponseWriter: ctx.Writer}
			ctx.Writer = cw
			ctx.Next()
			// read data in buf
			// do something for data
			// write to ResponseWriter
		*/
		// 打印响应体
		res := util.LogResponseBody(ctx)
		fmt.Printf("[INFO] Response: %s %s %s (%v)\n", ctx.Request.Method, ctx.Request.RequestURI, res, stopWatch)
	}
}
