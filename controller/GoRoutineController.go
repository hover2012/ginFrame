package controller

import (
	"fmt"
	"gin/models"
	"gin/pkg/e"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

func ShowWaitGroup(c *gin.Context) {
	code := 200
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		time.Sleep(1*time.Second)
		fmt.Println("goroutine 1 finished")
		wg.Done()
	}()
	go func() {
		//time.Sleep(1*time.Second)
		fmt.Println("goroutine 2 finished")
		wg.Done()
	}()

	wg.Wait()
	data :="alll goroutine finished!"

	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":e.GetMsg(code),
		"data":data,
	})
}

func CreateLog(c *gin.Context)  {

	subTitle := "测试副标题"
	other :="other"
	des :="描述"
	year :="2021"
	area :="北京"
	tag :="测试"
	star := "4"
	comment := "5"
	quote :="点赞数量"
	start := time.Now()
	var wg sync.WaitGroup
	wg.Add(10010)
	for i := 0;i<10000;i++ {
		title := "测试"
		num := strconv.Itoa(i)
		title = title + num
		moviemodel := &models.MovieModel{
			Title: title,
			Subtitle: subTitle,
			Other: other,
			Description: des,
			Year: year,
			Area: area,
			Tag: tag,
			Star: star,
			Comment: comment,
			Quote: quote,
		}
		go models.AddMovie(moviemodel)
		wg.Done()


		//log.Printf( num + "条数据入库成功")
	}
	wg.Wait()
    elapsed := time.Now().Sub(start)
    log.Printf("耗时：",elapsed)
	c.JSON(http.StatusOK,gin.H{
		"code":200,
		"msg":e.GetMsg(200),
		"data":"10000数据导入成功耗时",
	})
	//log.Printf("i:%d movie:v%",i,movie)

}


