package routers

import (
	"gin/controller"
	v1 "gin/controller/v1"
	"gin/pkg/setting"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/tags", v1.GetTags)
		apiv1.POST("/tags", v1.AddTag)
		apiv1.PUT("/tags/:id", v1.EditTag)
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
	}
	 spiders := r.Group("/spiders")
	{
		spiders.GET("/douban",controller.GetDouban)
		spiders.GET("/index",controller.ShowIndex)
		spiders.GET("/getJson",controller.ShowSecureJson)
	}
	gr := r.Group("/routine")
	{
		gr.GET("/showWaitGroup",controller.ShowWaitGroup)
		gr.GET("/createLog",controller.CreateLog)
	}
	paper := r.Group("/paper")
	{
		paper.GET("/get",controller.GetPaperDate)
		paper.GET("/getFile",controller.UpdateFile)
	}
	space := r.Group("/space",gin.BasicAuth(gin.Accounts{
		"wang":"bar",
		"liu":"liu",
	}))
	{
		space.GET("/index",controller.ShowSpaceList)
	}
	r.LoadHTMLGlob("resources/html/*")
	return r
}