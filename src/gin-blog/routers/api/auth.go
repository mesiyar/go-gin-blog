package api

import (
	"gin-blog/src/gin-blog/models"
	"gin-blog/src/gin-blog/pkg/e"
	"gin-blog/src/gin-blog/pkg/logging"
	"gin-blog/src/gin-blog/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

type auth struct {
	Username string `valid:"Required; MaxSize(20)"`
	Password string `valid:"Required; MaxSize(32)"`
}

// @Summary 获取认证
// @Produce  json
// @Param username query string true "用户名"
// @Param password query string true "密码"
// @Success 200 {string} json "{"code":200,"data":{"token":""},"msg":"ok"}"
// @Router /auth [get]
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
			logging.Info(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func CreateAuth(c *gin.Context) {

	username := c.Query("username")
	password := c.Query("password")
	valid := validation.Validation{}
	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(a)
	data := "ok"
	code := e.INVALID_PARAMS
	if ok {
		isExist := models.CheckAuthName(username)
		if isExist {
			createAuth := models.CreateAuthAccount(username, password)
			if createAuth {
				code = e.SUCCESS
			} else {
				data = "创建失败"
				code = e.ERROR
			}
		} else {
			code = e.ERROR_AUTHNAME_EXIST
		}
	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
