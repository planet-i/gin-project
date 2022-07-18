package util

import (
	"github.com/gin-gonic/gin"
	"github.com/planet-i/gin-project/pkg/setting"
	"github.com/unknwon/com"
)

// 获取分页页码
func GetPage(c *gin.Context) (result int) {
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * setting.AppSetting.PageSize
	}
	return
}
