package routers

import (
	"github.com/gin-gonic/gin"
	"../setting"
	"../middleware"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)

	initAuthRouter(r)

	initArticleRouter(r)

	initRatRouter(r)

	return r
}

func initAuthRouter(r *gin.Engine) {
	apiAuth := r.Group("/api/auth")
	{
		apiAuth.GET("/auth",GetAuth)
	}
}

func initArticleRouter(r *gin.Engine) {
	apiArticle := r.Group("/api/articles")
	r.Use(middleware.JWT())
	{
		apiArticle.GET("/getOne/:id", GetArticle)
		apiArticle.GET("/getList", GetArticles)
		apiArticle.GET("/addOne", AddArticle)
		apiArticle.GET("/updateOne/:id", EditArticle)
		apiArticle.GET("/deleteOne/:id", DeleteArticle)
	}
}

func initRatRouter(r *gin.Engine) {
	apiTag := r.Group("/api/tags")
	r.Use(middleware.JWT())
	{
		apiTag.GET("/getList", GetTags)
		apiTag.GET("/addOne", AddTag)
		apiTag.GET("/updateOne/:id", EditTag)
		apiTag.GET("/deleteOne/:id", DeleteTag)
	}
}