package router

import (
	"gin_demo/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	//实例化引擎
	r := gin.Default()

	r.Static("/static","./static")// /static/css/app.8eeeaf31.css等资源文件 从 ./static下找 告诉gin框架模板文件引用的静态文件去哪里找
	r.LoadHTMLGlob("./templates/*")//从templates目录下找index.html  告诉gin框架去哪里找模板文件

	r.GET("/", controller.IndexHandler)

	//定义路由组
	v1Group := r.Group("/v1")
	{
		//查询列表
		v1Group.GET("/todo", controller.GetAllTodo)
		//创建
		v1Group.POST("/todo", controller.CreateTodo)
		//修改
		v1Group.PUT("/todo/:id", controller.UpdateTodo)
		//删除
		v1Group.DELETE("/todo/:id", controller.DeleteTodo)
	}
	return r
}
