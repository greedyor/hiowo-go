package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"webkodes.com/admin/models"
)

// Index
func Index(c *gin.Context) {

	page := 1
	pageSize := 10
	offset := (page - 1) * pageSize

	var article []models.Article

	models.DB.Where("type = 1").Offset(offset).Limit(pageSize).Find(&article)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"articleList": article,
	})

}
