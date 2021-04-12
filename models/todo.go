package models

import (
	"bubble/dao"
)

type Toto struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func GetAllTodo() (todoList []*Toto, err error) {
	if err := dao.DB.Find(&todoList).Error; err != nil {
		return nil, err
	}
	return
}

func CreateTodo(todo *Toto) (err error) {
	err = dao.DB.Create(&todo).Error
	return
}

func UpdateTodo(todo *Toto) (err error) {
	err = dao.DB.Save(&todo).Error
	return
}

func DoDel(id string) (err error) {
	err = dao.DB.Where("id=?", id).Delete(Toto{}).Error
	return
}

func GetTodo(id string) (todo *Toto, err error) {
	todo = new(Toto)
	if err = dao.DB.Debug().Where("id=?", id).First(todo).Error; err != nil {
		return nil, err
	}
	return
}
