package routers

import (
	"net/http"

	"github.com/aifuxi/jianyu-blog-api/pkg/setting"
	v1 "github.com/aifuxi/jianyu-blog-api/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	r.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "new router struct",
		})
	})

	ap1v1 := r.Group("/api/v1")
	{
		// 获取标签列表
		ap1v1.GET("/tags", v1.GetTags)
		// 新建标签列表
		ap1v1.POST("/tags", v1.AddTag)
		// 更新标签列表
		ap1v1.PUT("/tags", v1.EditTag)
		// 删除标签列表
		ap1v1.DELETE("/tags", v1.DeleteTag)
	}

	return r
}
