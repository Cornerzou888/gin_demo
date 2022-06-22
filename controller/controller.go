package controller

import (
	"gin_demo/model"
	"github.com/gin-gonic/gin"
	"net/http"
)
//首页
func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK,"index.html",nil)
}

//列表
func GetAllTodo(c *gin.Context) {
	todos, err := model.GetAllTodo()
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"error" : err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK,todos)
}
//创建
func CreateTodo(c *gin.Context) {
	var todo model.Todo
	_ = c.BindJSON(&todo) //绑定参数
	//.ShouldBind()强大的功能，它能够基于请求自动提取JSON、form表单和QueryString类型的数据，并把值绑定到指定的结构体对象
	//_ = c.ShouldBind(&todo)
	err := model.CreateATodo(&todo)
	if err != nil{
		c.JSON(http.StatusOK,gin.H{
			"error" : err.Error(),
		})
		return
	}
	//todo 如何判断 是否数据库里已经有
	c.JSON(http.StatusOK,gin.H{
		"msg" : "success",
		"data": todo,
	})
	//c.JSON(http.StatusOK,todo)
}

//修改
func UpdateTodo(c *gin.Context) {
	//id := c.Param("id")
	id,ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK,gin.H{
			"error":"无效的id",
		})
		return
	}
	todo,err := model.GetTodoById(id)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"error":err.Error(),
		})
		return
	}
	//fmt.Println(todo)
	_ = c.BindJSON(&todo)//将json参数绑定到todo 再保存到数据库
	//fmt.Println(todo)
	err = model.UpdateATodo(&todo)
	if err != nil{
		c.JSON(http.StatusOK,gin.H{
			"error":err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK,todo)
}
//删除
func DeleteTodo(c *gin.Context) {
	id,ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK,gin.H{
			"error":"无效的id",
		})
		return
	}
	//硬删除
	err := model.DeleteATodo(id)
	if err != nil{
		c.JSON(http.StatusOK,gin.H{
			"error":err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		id : "deleted",
	})
}