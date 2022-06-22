package model

import (
	"gin_demo/dao"
)

//代办 model
type Todo struct {
	ID uint `json:"id"`
	Title string `json:"title" gorm:"unique_index;not null"`
	Status bool `json:"status"`
}

//增删改查

func CreateATodo(todo *Todo) (err error) {
	err = dao.DB.Debug().Create(todo).Error
	return
}

func GetAllTodo() (todos []*Todo, err error){
	err = dao.DB.Debug().Find(&todos).Error
	//return todos,err
	return
}

func GetTodoById(id string) (todo Todo,err error){
	err = dao.DB.Debug().Where("id=?",id).First(&todo).Error
	return todo,err
}

func UpdateATodo(todo *Todo) (err error) {
	err = dao.DB.Debug().Save(todo).Error
	return
}

func DeleteATodo(id string) (err error) {
	err = dao.DB.Where("id=?",id).Delete(&Todo{}).Error
	return
}

func InitModel(){
	//表结构迁移
	dao.DB.AutoMigrate(&Todo{})
}


