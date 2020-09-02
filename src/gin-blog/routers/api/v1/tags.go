package v1

import (
	"gin-blog/src/gin-blog/models"
	"gin-blog/src/gin-blog/pkg/e"
	"gin-blog/src/gin-blog/pkg/setting"
	"gin-blog/src/gin-blog/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/unknwon/com"
)

// 获取tags
func GetTags(c *gin.Context) {
	name := c.Query("name")
	maps := make(map[string]interface{})
	data := make(map[string]interface{})
	if name != "" {
		maps["name"] = name
	}

	var state int = -1

	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	code := e.SUCCESS

	data["lists"] = models.GetTags(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// 添加文章标签
func AddTag(c *gin.Context) {

}

//修改文章标签
func EditTag(c *gin.Context) {
}

//删除文章标签
func DeleteTag(c *gin.Context) {
}