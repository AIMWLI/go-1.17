package util

import (
	"github.com/gin-gonic/gin"
	"go-gin/pkg/setting"
	"strconv"
)

func GetPage(c *gin.Context) int {
	result := 0
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		return 0
	}
	size, err := strconv.Atoi(setting.AppSetting.PageSize)
	if err != nil {
		return 0
	}
	if page > 0 {
		result = (page - 1) * size
	}
	return result
}
