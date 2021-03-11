package v1

import (
	"gin/models"
	"gin/pkg/e"
	"gin/pkg/log"
	"gin/pkg/setting"
	"gin/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/unknwon/com"
	)

func GetTags(c *gin.Context)  {
	name := c.Query("name")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name !=""{
		maps["name"] = name
	}
	log.Info(name,"获取参数信息")
	var state int =-1
	if arg := c.Query("state"); arg != ""{
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	code := e.SUCCESS

	data["lists"] = models.GetTags(util.GetPage(c),setting.PageSize,maps)
	data["total"] = models.GetTagTotal(maps)

	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":e.GetMsg(code),
		"data":data,
	})
}

func AddTag(c *gin.Context )  {
	name :=c.Query("name")
	state :=com.StrTo(c.DefaultQuery("state","0")).MustInt()
	createdBy := c.Query("created_by")

	valid := validation.Validation{}

	valid.Required(name,"name").Message("名称不能为空")
	valid.MaxSize(name,100,"name").Message("名称最长10字符")
	valid.Required(createdBy,"created_by").Message("创建人不能为空")
	valid.MaxSize(createdBy,10,"created_by").Message("创建人不能为空")
	valid.Range(state,0,1,"state").Message("状态只能是")

	code := e.INVALID_PARAMS
	if  ! valid.HasErrors() {
		if ! models.ExistTagByName(name) {
			code = e.SUCCESS
			models.AddTag(name,state,createdBy)
		} else {
			code = e.ERROR_EXIST_TAG
		}
	}

	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":e.GetMsg(code),
		"data":make(map[string]string),
	})
}

func EditTag(c *gin.Context)  {
	id := com.StrTo(c.Param("id")).MustInt()
	name :=c.Query("name")

	modifiedBy := c.Query("modified_by")

	valid := validation.Validation{}

	var state int = -1
	if arg :=c.Query("state");arg !=""{
		state = com.StrTo(arg).MustInt()
		valid.Range(state,0,1,"state").Message("状态只能是0或者1")
	}


	valid.Required(name,"name").Message("名称不能为空")
	valid.MaxSize(name,100,"name").Message("名称最长10字符")
	valid.Required(modifiedBy,"modified_by").Message("创建人不能为空")
	valid.Range(state,0,1,"state").Message("状态只能是")

	code := e.INVALID_PARAMS

	if ! valid.HasErrors(){
		code = e.SUCCESS
		if models.ExitTagByID(id){
			data := make(map[string]interface{})
			data["modified_by"] = modifiedBy
			if name != ""{
				data["name"] = name
			}
			if state != -1 {
				data["state"] = state
			}
			models.EditTag(id ,data)
		} else {
			code = e.ERROR_EXIST_TAG
		}
	}

	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":e.GetMsg(code),
		"data":make(map[string]string),
	})

}

func DeleteTag(c *gin.Context)  {
	 id := com.StrTo(c.Param("id")).MustInt()
	 valid := validation.Validation{}
	 valid.Min(id,1,"id").Message("Id 必须大于0")

	 code :=e.INVALID_PARAMS
	 if ! valid.HasErrors(){
	 	code = e.SUCCESS
	 	if models.ExitTagByID(id){
	 		models.DeleteTag(id)
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	 }

	 c.JSON(http.StatusOK,gin.H{
		 "code" : code,
		 "msg" : e.GetMsg(code),
		 "data" : make(map[string]string),
	 })
}