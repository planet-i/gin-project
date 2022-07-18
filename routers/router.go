package routers

import (
	_ "github.com/EDDYCJY/go-gin-example/docs"
	"github.com/gin-gonic/gin"
	"github.com/planet-i/gin-project/middleware/jwt"
	"github.com/planet-i/gin-project/pkg/setting"
	"github.com/planet-i/gin-project/routers/api"
	v1 "github.com/planet-i/gin-project/routers/api/v1"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	gin.SetMode(setting.ServerSetting.RunMode)

	r.GET("/auth", api.GetAuth)
	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		// 获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		// 添加标签
		apiv1.POST("/tags", v1.AddTag)
		// 修改指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		// 删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		// 获取指定文章
		apiv1.GET("/articles/:id", v1.AddrArticle)
		// 获取文章列表
		apiv1.GET("/articles", v1.GetAritcles)
		// 添加文章
		apiv1.POST("/articles", v1.AddrArticle)
		// 修改指定文章
		apiv1.PUT("/articles/:id", v1.EditArticle)
		// 删除指定文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
	}

	return r
}
