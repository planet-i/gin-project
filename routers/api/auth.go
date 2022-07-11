package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/planet-i/gin-project/models"
	"github.com/planet-i/gin-project/pkg/e"
	"github.com/planet-i/gin-project/pkg/util"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"` //;后必须有空格
	Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	valid := validation.Validation{}
	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)

	data := make(map[string]interface{})
	code := e.INVALID_PARAMS
	if ok {
		isExist := models.CheckAuth(username, password)

		if isExist {
			token, err := util.GenerateToken(username, password)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token
				code = e.SUCCESS
			}
		} else {
			code = e.ERROR_AUTH
		}

	} else {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}
	fmt.Println(e.GetMsg(code))
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
