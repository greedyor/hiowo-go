package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"hiowo.com/models"
)

// Index
func Index(c *gin.Context) {

	// 默认页数
	page := 1
	// 每页数据数量
	pageSize := 10
	// 查询偏移量
	offset := (page - 1) * pageSize

	var article []models.Article

	models.DB.Where("type = 1").Offset(offset).Limit(pageSize).Find(&article)

	c.HTML(http.StatusOK, "index.html", RenderData(H{
		"articleList": article,
	}))

}
