package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/planet-i/gin-project/api"
	v1 "github.com/planet-i/gin-project/api/v1"
	_ "github.com/planet-i/gin-project/docs"
	"github.com/planet-i/gin-project/middleware/jwt"
	"github.com/planet-i/gin-project/pkg/export"
	"github.com/planet-i/gin-project/pkg/qrcode"
	"github.com/planet-i/gin-project/pkg/setting"
	"github.com/planet-i/gin-project/pkg/upload"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	gin.SetMode(setting.ServerSetting.RunMode)

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	r.StaticFS("/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))

	r.POST("/auth", api.GetAuth)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/upload", api.UploadImage)

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
		// 标签导出到Excel文件
		r.POST("/tags/export", v1.ExportTag)
		//导入标签
		r.POST("/tags/import", v1.ImportTag)

		// 获取指定文章
		apiv1.GET("/articles/:id", v1.AddrArticle)
		// 获取文章列表
		apiv1.GET("/articles", v1.GetArticles)
		// 添加文章
		apiv1.POST("/articles", v1.AddrArticle)
		// 修改指定文章
		apiv1.PUT("/articles/:id", v1.EditArticle)
		// 删除指定文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
		// 生成文章海报
		apiv1.POST("/articles/poster/generate", v1.GenerateArticlePoster)
	}

	return r
}
