package api

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go-gin/models"
	"go-gin/pkg/e"
	"go-gin/util"
	"log"
	"net/http"
)

type Auth struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	auth := Auth{Username: username, Password: password}
	err := validator.New().Struct(auth)
	if err != nil {
		log.Printf("[WARN] GetAuth.validate err: %v", err)
	}
	data := make(map[string]interface{})
	code := e.INVALID_PARAMS
	isExist := models.CheckAuth(username, password)
	if isExist {
		//token, err := util.GenerateToken(username, password)
		token, err := util.GenerateToken(username, "") //不在JWT的payload或header中放置敏感信息
		if err != nil {
			code = e.ERROR_AUTH_TOKEN
		} else {
			data["token"] = token
			code = e.SUCCESS
		}

	} else {
		code = e.ERROR_AUTH_TOKEN
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})

}
