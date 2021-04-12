package routes

import (
	"bubble/controller"

	"github.com/gin-gonic/gin"
)

func InitRoute() *gin.Engine {
	r := gin.Default()
	r.Static("/static", "static")
	r.LoadHTMLGlob("templates/*")
	r.GET("/", controller.IndexHandler)

	v1Group := r.Group("v1")
	{
		//待办事项
		v1Group.GET("/todo", controller.TodoListHandler)
		//添加
		v1Group.POST("/todo", controller.AddTodoHandler)

		//修改
		v1Group.PUT("/todo/:id", controller.DoUpdateTodo)
		//删除
		v1Group.DELETE("/todo/:id", controller.DoDel)
	}
	return r
}
