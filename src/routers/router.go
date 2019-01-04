package routers

import (
	"github.com/gin-gonic/gin"
	"../setting"
	"net/http"
)

func InitRouter() *gin.Engine  {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)

	initArticleRouter(r)

	initRatRouter(r)

	return r
}

func initArticleRouter(r *gin.Engine)  {
	apiArticle := r.Group("/api/articles")
	{
		apiArticle.GET("/getOne",GetArticle)
	}
}

func initRatRouter(r *gin.Engine)  {
	apiTag := r.Group("/api/Tags")
	{
		apiTag.GET("/getList",GetTags)
		apiTag.GET("/AddOne",AddTag)
		apiTag.GET("/updateOne",EditTag)
		apiTag.GET("/deleteOne",DeleteTag)
	}
}

// func1: 处理最基本的GET
func func1(c *gin.Context) {
	// 回复一个200OK,在client的http-get的resp的body中获取数据
	c.JSON(http.StatusOK, gin.H{
		"message":"test1",
	})
}

// func2: 处理最基本的POST
func func2(c *gin.Context) {
	// 回复一个200 OK, 在client的http-post的resp的body中获取数据
	c.String(http.StatusOK, "test2 OK")
}
