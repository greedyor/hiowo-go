package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"webkodes.com/admin/models"
)

// Note
func Note(c *gin.Context) {

	page := 1
	pageSize := 10
	offset := (page - 1) * pageSize

	// 查询
	var article []models.Article
	models.DB().Where("type = 1").Offset(offset).Limit(pageSize).Find(&article)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"articleList": article,
	})
}

// NoteDetail
func NoteDetail(c *gin.Context) {

	// 请求参数判断
	id := c.DefaultQuery("id", "")
	var getDetail struct {
		ID int16 `form:"id" json:"username" binding:"required"` //必填
	}

	err := c.ShouldBindWith(&getDetail, binding.Query)

	if err != nil {
		Error(c)
		return
	}

	// 查询文章信息
	var article models.Article
	result := models.DB().Select("title").Where("type = 1 and id=?", id).First(&article)

	if result.Error != nil {
		Error(c)
		return
	}

	// 查询详情
	var articleDetail models.ArticleDetail
	models.DB().Select("content").Where("article_id=?", id).First(&articleDetail)

	// 查询列表
	var articleList []models.Article
	models.DB().Where("type = 1").Offset(0).Limit(10).Find(&articleList)

	c.HTML(http.StatusOK, "detail.html", gin.H{
		"title":       article.Title,
		"detail":      articleDetail.Content,
		"articleList": articleList,
	})

}
