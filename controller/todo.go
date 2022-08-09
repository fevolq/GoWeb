package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"go-web/models/todo"
)

func TodoAdd(c *gin.Context) {
	var code int

	title := c.DefaultPostForm("title", "test")
	memo := c.DefaultPostForm("memo", "备注")

	todoObj := new(todo.Todo)
	_, err := todoObj.Add(title, memo)
	if err == nil {
		code = 20000
	} else {
		code = 20001
	}
	c.JSON(http.StatusOK, gin.H{"code": code})
}

func TodoEdit(c *gin.Context) {
	var code int
	updateInfo := map[string]interface{}{}

	tdID, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"meg": "id类型异常"})
		return
	}
	title, ok := c.GetPostForm("title")
	if ok {
		updateInfo["title"] = title
	}
	memo, ok := c.GetPostForm("memo")
	if ok {
		updateInfo["memo"] = memo
	}
	status, ok := c.GetPostForm("status")
	if ok {
		updateInfo["status"] = status
	}

	td := new(todo.Todo)
	err = td.Init(tdID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"meg": "未查找到该记录"})
		return
	}

	err = td.Edit(updateInfo)

	if err != nil {
		code = 20001
	} else {
		code = 20000
	}
	c.JSON(http.StatusOK, gin.H{"code": code})
}
