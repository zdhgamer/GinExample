package routers

import (
	"github.com/gin-gonic/gin"
	"../setting"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "../docs"
	"../upload"
	"net/http"
	"../middleware"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.ServerSetting.RunMode)

	initStaticRouter(r)

	initSwaggerRouter(r)

	initUploadRouter(r)

	initAuthRouter(r)

	initArticleRouter(r)

	initRatRouter(r)

	return r
}

func initStaticRouter(r *gin.Engine) {
	//路径映射
	r.StaticFS("src/runtime/upload/images", http.Dir(upload.GetImageFullPath()))
}

func initSwaggerRouter(r *gin.Engine) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func initUploadRouter(r *gin.Engine) {
	apiUpload := r.Group("api/upload")
	{
		apiUpload.POST("uploadImage", UploadImage)
	}
}

func initAuthRouter(r *gin.Engine) {
	apiAuth := r.Group("/api/auth")
	{
		apiAuth.GET("/auth", GetAuth)
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
