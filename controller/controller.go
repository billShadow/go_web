package controller

import (
	"bubble/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func TodoListHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"error": "222"})

	// todoList, err := models.GetAllTodo()
	// if err != nil {
	// 	c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	// } else {
	// 	c.JSON(http.StatusOK, todoList)
	// }

}

func AddTodoHandler(c *gin.Context) {
	// 1. 从请求中把数据拿出来
	var todo models.Toto
	c.BindJSON(&todo)
	// 2. 存入数据库
	err := models.CreateTodo(&todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func DoUpdateTodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效id"})
		return
	}
	todo, err := models.GetTodo(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	c.BindJSON(&todo)
	if err = models.UpdateTodo(todo); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todo)
}

func DoDel(c *gin.Context) {
	id, _ := c.Params.Get("id")
	if err := models.DoDel(id); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "删除成功"})
}
