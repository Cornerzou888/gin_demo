package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //导入mysql驱动
)
//初始化
var (
	DB *gorm.DB
)

//连接数据库
func InitMysql(env string) (err error){
	if env == "release" {
		DB, err = gorm.Open("mysql","root:123456@(localhost:13306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	}else{
		DB, err = gorm.Open("mysql","root:@(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	}
	return err
	//return DB.DB().Ping() //ping不通的话 sql: database is closed
}


