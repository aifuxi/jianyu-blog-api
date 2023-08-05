package routers

import (
	"net/http"

	"github.com/aifuxi/jianyu-blog-api/middleware/jwt"
	"github.com/aifuxi/jianyu-blog-api/pkg/setting"
	"github.com/aifuxi/jianyu-blog-api/routers/api"
	v1 "github.com/aifuxi/jianyu-blog-api/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.ServerSetting.RunMode)

	r.GET("/auth", api.GetAuth)

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "pong!",
		})
	})

	r.Use(jwt.JWT())
	ap1v1 := r.Group("/api/v1")
	{
		// 获取标签列表
		ap1v1.GET("/tags", v1.GetTags)
		// 新建标签
		ap1v1.POST("/tags", v1.AddTag)
		// 更新标签
		ap1v1.PUT("/tags/:id", v1.EditTag)
		// 删除标签
		ap1v1.DELETE("/tags/:id", v1.DeleteTag)
	}

	{
		// 获取文章列表
		ap1v1.GET("/articles", v1.GetArticles)
		// 获取指定文章
		ap1v1.GET("/articles/:id", v1.GetArticle)
		// 新建文章
		ap1v1.POST("/articles", v1.AddArticle)
		// 更新文章
		ap1v1.PUT("/articles/:id", v1.EditArticle)
		// 删除文章
		ap1v1.DELETE("/articles/:id", v1.DeleteArticle)
	}

	return r
}
