package routers

import (
	v1 "blog/internal/routers/v1"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	tag := v1.NewTag()
	article := v1.NewArticle()
	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/tags", tag.Get)
		apiv1.DELETE("/tags/:id", tag.Delete)
		apiv1.PUT("/tags/:id", tag.Update)
		apiv1.PATCH("/tags/:id/state", tag.Update)
		apiv1.GET("/tags", tag.List)

		apiv1.POST("/article", article.Create)
		apiv1.DELETE("/article/:id", article.Delete)
		apiv1.PUT("/article/:id", article.Update)
		apiv1.PATCH("/article/:id/state", article.Update)
		apiv1.GET("/articles/:id", article.Get)
		apiv1.GET("/articles", article.List)
	}

	return r
}
