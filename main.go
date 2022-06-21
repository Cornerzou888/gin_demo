package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"//导入mysql驱动
	"net/http"
)
//代办 model
type Todo struct {
	ID uint `json:"id"`
	Title string `json:"title" gorm:"unique_index;not null"`
	Status bool `json:"status"`
	IgnoreMe int `gorm:"-"` // 忽略本字段
}

//初始化
var (
	DB *gorm.DB
)

//连接数据库
func initMysql() (err error){
	DB, err = gorm.Open("mysql","root:@(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	return err
	//return DB.DB().Ping() //ping不通的话 sql: database is closed
}
func main()  {
	//创建数据库 本地test库
	//连接数据库
	err := initMysql()
	if err != nil {
		panic(err)
	}
	//延时关闭数据库连接
	defer DB.Close()
	//表结构迁移
	DB.AutoMigrate(&Todo{})
	//实例化引擎
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
			var todos []Todo
			err = DB.Debug().Find(&todos).Error
			if err != nil {
				c.JSON(http.StatusOK,gin.H{
					"error" : err.Error(),
				})
			}
			c.JSON(http.StatusOK,todos)
		})
		//查询某一个
		v1Group.GET("/todo/:id", func(c *gin.Context) {
			c.JSON(http.StatusOK,gin.H{
				"msg" : "info",
			})
		})
		//创建
		v1Group.POST("/todo", func(c *gin.Context) {
			var todo Todo
			_ = c.BindJSON(&todo) //绑定参数
			//.ShouldBind()强大的功能，它能够基于请求自动提取JSON、form表单和QueryString类型的数据，并把值绑定到指定的结构体对象
			//_ = c.ShouldBind(&todo)
			err = DB.Debug().Create(&todo).Error
			if err != nil{
				c.JSON(http.StatusOK,gin.H{
					"error" : err.Error(),
				})
			}
			//todo 如何判断 是否数据库里已经有
			c.JSON(http.StatusOK,gin.H{
				"msg" : "success",
				"data": todo,
			})
			//c.JSON(http.StatusOK,todo)
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
