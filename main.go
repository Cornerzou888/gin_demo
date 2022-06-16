package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main()  {

	r := gin.Default()

	r.Static("/static","./static")// /static/css/app.8eeeaf31.css等资源文件 从 ./static下找 告诉gin框架模板文件引用的静态文件去哪里找
	r.LoadHTMLGlob("./templates/*")//从templates目录下找index.html  告诉gin框架去哪里找模板文件

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK,"index.html",nil)
	})

	//定义路由组
	v1Group := r.Group("/v1")
	{
		//查询
		v1Group.GET("/todo", func(c *gin.Context) {
			c.JSON(http.StatusOK,gin.H{
				"msg" : "list",
			})
		})
		//创建
		v1Group.POST("/todo", func(c *gin.Context) {
			c.JSON(http.StatusOK,gin.H{
				"msg" : "create",
			})
		})

		//修改
		v1Group.PUT("/todo/:id", func(c *gin.Context) {
			c.JSON(http.StatusOK,gin.H{
				"msg" : "update",
			})
		})
		//删除
		v1Group.DELETE("/todo/:id", func(c *gin.Context) {
			c.JSON(http.StatusOK,gin.H{
				"msg" : "del",
			})
		})

	}



	r.Run(":9090")

}
