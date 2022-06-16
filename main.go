package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)
//代办 model
type Todo struct {
	ID uint `json:"id"`
	Title string `json:"title" gorm:"unique_index;not null"`
	Status bool `json:"status"`
	IgnoreMe int `gorm:"-"` // 忽略本字段
}

func main()  {
	//创建数据库 本地 test库
	//连接数据库
	db, err := gorm.Open("mysql","root:@(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("错误了")
		panic(err)
	}
	//延时关闭
	defer db.Close()
	//表结构迁移
	db.AutoMigrate(&Todo{})

	r := gin.Default()

	r.Static("/static","./static")// /static/css/app.8eeeaf31.css等资源文件 从 ./static下找 告诉gin框架模板文件引用的静态文件去哪里找
	r.LoadHTMLGlob("./templates/*")//从templates目录下找index.html  告诉gin框架去哪里找模板文件

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK,"index.html",nil)
	})

	//定义路由组
	v1Group := r.Group("/v1")
	{
		//查询列表
		v1Group.GET("/todo", func(c *gin.Context) {
			c.JSON(http.StatusOK,gin.H{
				"msg" : "list",
			})
		})
		//查询某一个
		v1Group.GET("/todo/:id", func(c *gin.Context) {
			c.JSON(http.StatusOK,gin.H{
				"msg" : "info",
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
