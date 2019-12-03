package controller

import (
	"net/http"

	"../db"
	"../entity"
	"github.com/gin-gonic/gin"
)

type Controller struct{}

// Create ... レコードの生成
func (ctlr Controller) Create(c *gin.Context) {
	task := entity.Task{}
	db := db.GetDB()
	err := c.BindJSON(&task)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	db.Create(&task)
	if db.NewRecord(task) == false {
		c.JSON(http.StatusOK, task)
	}
}

// ReadAll ... 全タスクの取得
func (ctlr Controller) ReadAll(c *gin.Context) {
	var task []entity.Task
	db := db.GetDB()
	db.Find(&task)
	c.JSON(200, task)
}

// ReadByID ... 一件のタスクの取得
func (ctlr Controller) ReadByID(c *gin.Context) {
	var task []entity.Task
	db := db.GetDB()
	id := c.Params.ByName("id")
	db.First(&task, id)
	c.JSON(200, task)
}

// UpdateByID ... レコードの更新
func (ctlr Controller) UpdateByID(c *gin.Context) {
	task := entity.Task{}
	db := db.GetDB()
	id := c.Params.ByName("id")
	err := c.BindJSON(&task)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	db.Where("id=?", id).Update(task)
	db.Save(&task)
}

// DeleteByID ... レコードの削除
func (ctlr Controller) DeleteByID(c *gin.Context) {
	var task []entity.Task
	db := db.GetDB()
	id := c.Params.ByName("id")
	db.Where("id=?", id).Delete(task)
}
