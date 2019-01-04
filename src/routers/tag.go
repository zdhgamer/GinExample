package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/Unknwon/com"
	"../e"
	"../models"
	"../util"
	"../setting"
	"net/http"
	"github.com/astaxie/beego/validation"
)

func GetTags(c *gin.Context)  {
	name := c.Query("name")
	maps := make(map[string]interface{})
	data := make(map[string]interface{})
	if name != "" {
		maps["name"] = name
	}
	state := -1
	if arg := c.Query("state"); arg!="" {
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

func AddTag(c *gin.Context)  {
	name := c.Query("name")
	state := com.StrTo(c.DefaultQuery("state","0")).MustInt()
	createdBy := c.Query("created_by")

	valid := validation.Validation{}
	valid.Required(name,"name").Message("名字不能为空")
	valid.MaxSize(name,100,"name").Message("名称最长为100字符")
	valid.Required(createdBy,"created_by").Message("创建人不能为空")
	valid.Range(state,0,0,"state").Message("状态只能是0或者1")

	code := e.INVALID_PARAMS

	if ! valid.HasErrors() {
		if !models.ExistTagByName(name) {
			code = e.SUCCESS
			if re := models.AddTag(name,state,createdBy);re{
				code = e.ERROR
			}
		}else {
			code = e.ERROR_EXIST_TAG
		}
	}

	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":e.GetMsg(code),
		"data":make(map[string]interface{}),
	})
}

//修改文章标签
func EditTag(c *gin.Context) {
}

//删除文章标签
func DeleteTag(c *gin.Context) {
}