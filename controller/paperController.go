package controller

import (
	"encoding/json"
	"fmt"
	"gin/models"
	"gin/pkg/e"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
)

var  (
	PaperBaseUrl = "http://wx.xuexiguanjia.org"
	Page = 0
)

func GetPaperDate(c *gin.Context)  {
	xueke := c.Query("xueke")
	xueduan := c.Query("xueduan")
	//code := e.SUCCESS
	num := 0
	for  {
		pageStr := strconv.Itoa(Page)
		url :=  PaperBaseUrl + "/api/products?page="+pageStr+"&limit=100&store_subject="+ strings.Replace(xueke, " ", "", -1)+"&store_grade="+ strings.Replace(xueduan, " ", "", -1)+"&store_type=&store_year="
		result := request("GET",url)
		//	log.Printf("++++++++",result["data"])
		data := result["data"]
		lend := len(data.([]interface{}))
		if lend == 0  {
			Page = 0
			break
		}
		addData(data.([]interface{}),xueke,xueduan)
		num +=  lend
		Page++
	}

	//fmt.Println(len(data.([]interface{})))


	//for _,v := range data.([]interface{}) {
	//	log.Printf("++++++++",v)
	//	fmt.Println(reflect.TypeOf(v))
	//	log.Printf("********",v.(map[string]interface{})["id"])
	//}
	//fmt.Println(reflect.TypeOf(data))

	//c.JSON(http.StatusOK,result)
	c.JSON(http.StatusOK,gin.H{
		"code":200,
		"msg":e.GetMsg(200),
		"data":num,
	})

}

func addData(data []interface{},xk string,xd string)  {
	var paperData models.Paper
	for _,v := range data {
		flag := v.(map[string]interface{})
		id, _ := flag["id"].(float64)
		store_name,_ := flag["store_name"].(string)
		word_answer,ok := flag["word_answer"].(string)

		if !ok {
			fmt.Println("参数非法")
			continue
		}
		str := strings.Split(word_answer,".")
		store_type,ok := flag["store_type"].(string)
		//store_name,_ := flag["store_name"]
		//store_name,_ := flag["store_name"]
		//fmt.Println(int(id))
		paperData.ID = int(id)
		paperData.ExmId = int(id)
		paperData.Name =  store_name
		paperData.DetailUrl = PaperBaseUrl + "/home/file_detail/?id=" + strconv.Itoa(paperData.ExmId)
		paperData.DocUrl = PaperBaseUrl + word_answer
		paperData.PdfUrl = PaperBaseUrl + str[0] + ".pdf"
		paperData.StoreSubject = store_type
		paperData.Xueke = xk
		paperData.Xueduan = xd

		//log.Printf(" data:v%",paperData)
		models.AddPaper(&paperData)
	}
}

//func assertData(d interface{}) (interface{},error)  {
//	switch d.(type) {
//	case float64:
//		r,_ := d.(float64)
//		return r,nil
//	case string:
//		r,_:= d.(string)
//		return r,nil
//	default:
//
//	}
//}

func request(method ,url string) map[string]interface{}  {
	client := &http.Client{}
	reqest,err := http.NewRequest(method,url,nil)
	reqest.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	reqest.Header.Set("Accept-Charset", "GBK,utf-8;q=0.7,*;q=0.3")
	reqest.Header.Set("Accept-Encoding", "gzip,deflate,sdch")
	reqest.Header.Set("Accept-Language", "zh-CN,zh;q=0.8")
	reqest.Header.Set("Cache-Control", "max-age=0")
	reqest.Header.Set("Connection", "keep-alive")
	//reqest.Header.Set("Content-Type", "multipart/form-data")
	if err != nil{
		log.Print("请求错误",err)
	}
	//处理返回结果
	response,_ := client.Do(reqest)
	//获取body信息
	body,_ := ioutil.ReadAll(response.Body)
	result := make(map[string]interface{},0)
	//把json 转换成map
	err = json.Unmarshal(body,&result)

	if err != nil {
		log.Printf("json 转换错误", err)
	}

	status := response.StatusCode

	if status !=200 {
		log.Printf("请求错误地址",status)
	}

	return result
}

func UpdateFile(c *gin.Context)  {
	xueke := c.Query("xueke")
	xueduan := c.Query("xueduan")
	maps := make(map[string]interface{})

	maps["xueke"] = xueke
	maps["xueduan"] = xueduan

	data := models.GetPaper(maps)
	fmt.Println(reflect.TypeOf(data))
	go downLoad(data)
	code := e.SUCCESS
	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":e.GetMsg(code),
		"data":data,
	})
}

func downLoad(data []models.Paper)  {
	baseDir := "/Users/wanfei/Desktop/"
	for _,v := range data{
		dir := baseDir + v.Xueke+ "/"+ v.Xueduan
		os.MkdirAll(dir,os.ModePerm)
		doDownLoad(v.DocUrl,dir,v.Name,"doc")
		//go  doDownLoad(v.PdfUrl,dir,v.Name,"pdf")
	}
}

func doDownLoad(url string,dir string, name string,ext string)(bool) {
	log.Printf("文件链接"+url)
	res,err := http.Get(url)
	if err != nil{
		log.Printf("文件获取失败[1]", err)
		return false
	}
	f,err := os.Create(dir +"/" +name+"."+ext)
	if err != nil{
		log.Printf("文件创建失败[2]",err)
		return false
	}

	defer f.Close()
 	 _,err =	io.Copy(f,res.Body)
	if err != nil{
		log.Printf("文件创建失败[3]",err)
		return false
	}
	return true
}


/**
* @desc 修改数据
**/
func UpdatePaperData(c *gin.Context)  {
	maps := make(map[string]interface{})
	maps["doc_url"] = PaperBaseUrl
	baseDir := "/Users/wanfei/Desktop/"

	datas := models.GetPaper(maps)
	for _,v := range datas{
		id := strconv.Itoa(v.ID)
		url := PaperBaseUrl + "/api/product/detail/" + id
		fmt.Println(url)
		result :=request("GET",url)
		storInfos := result["data"].(map[string]interface{})
		storInfo  := storInfos["storeInfo"].(map[string]interface{})
		pdf_paper :=  storInfo["pdf_paper"]
		name := storInfo["store_name"].(string)
		word_paper := PaperBaseUrl +  storInfo["word_paper"].(string)
		//doDownLoad(v.DocUrl,dir,v.Name,"doc")
		dir := baseDir + v.Xueke+ "/"+ v.Xueduan
		os.MkdirAll(dir,os.ModePerm)
	 go 	doDownLoad(word_paper,dir,name,"doc")

		updateData := make(map[string]interface{})
		updateData["doc_url"] =word_paper
		updateData["pdf_url"] =pdf_paper

		models.UpdatePaper(v.ID,updateData)

		//fmt.Println(word_paper)
		fmt.Println("succ" + id)
		//break;
	}

	c.JSON(http.StatusOK,gin.H{
		"code":200,
		"msg":e.GetMsg(200),
		"data":"",
	})
}
