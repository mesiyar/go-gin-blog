package routers

import (
	"gin-blog/src/gin-blog/middleware/jwt"
	"gin-blog/src/gin-blog/pkg/setting"
	"gin-blog/src/gin-blog/routers/api"
	v1 "gin-blog/src/gin-blog/routers/api/v1"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	r.GET("/auth", api.GetAuth)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiv1 := r.Group("/api/v1")

	apiv1.Use(jwt.JWT())
	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		// region 文章
		/*获取文章列表*/
		apiv1.GET("/articles", v1.GetArticles)
		apiv1.GET("/article/:id", v1.GetArticle)
		//新建文章
		apiv1.POST("/article", v1.AddArticle)
		//更新文章
		apiv1.PUT("/article/:id", v1.EditArticle)
		//删除指定文章
		apiv1.DELETE("/article/:id", v1.DeleteArticle)
		// endregion 文章
	}

	return r
}
