package main

import (
	"gin_demo/dao"
	"gin_demo/model"
	"gin_demo/router"
)

func main()  {
	//创建数据库 本地test库
	//连接数据库
	err := dao.InitMysql("")
	if err != nil {
		panic(err)
	}
	//延时关闭数据库连接
	defer dao.DB.Close()
	//绑定模型
	model.InitModel()
	//注册路由
	r := router.SetupRouter()
	r.Run(":9090")

}
