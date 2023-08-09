package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin/models"
	"go-gin/pkg/app"
	"go-gin/pkg/e"
	"log"
	"net/http"
)

func Save(c *gin.Context) {
	result := app.Gin{C: c}

	// 取json字符串
	/*
		json := make(map[string]interface{})
		c.BindJSON(&json)
		log.Printf("%v", &json)
	*/
	// 序列化结构体, 注意如无法绑定参数,确定是否
	user := models.User{}
	err := c.ShouldBind(&user)
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Printf("%v", &user)
	isSave := models.Save(&user)
	log.Printf("save user %v", isSave)
	/*
		c.JSON(http.StatusOK, gin.H{
			"code": e.SUCCESS,
			"data": isSave,
		})
	*/
	result.Response(http.StatusOK, e.SUCCESS, isSave)
}
func SelectPage(c *gin.Context) {
	users := models.SelectPage(0, 10, nil)
	c.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"data": users,
	})
}
