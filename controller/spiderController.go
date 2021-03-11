package controller

import (
	"gin/pkg/e"
	"gin/pkg/log"
	"gin/pkg/spider"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

var (
	BaseUrl = "https://movie.douban.com/top250"
)

//获取切片数据
func GetDouban(c *gin.Context)  {
	name := c.Query("name")
	log.Info(name,"获取参数信息")
	code := e.SUCCESS
	var movies [][]spider.DoubanMove

	pages := spider.GetPages(BaseUrl)
	for _,page := range pages{
		url := strings.Join([]string{BaseUrl,page.Url},"")
		move :=  spider.GetMovies(url)
		//fmt.Println("type:", reflect.TypeOf(movies))
		movies = append(movies ,move)
	}

	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":e.GetMsg(code),
		"data":movies,
	})

}
