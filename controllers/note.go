package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"hiowo.com/models"
)

// Note
func Note(c *gin.Context) {

	Index(c)

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
	result := models.DB.Select("title").Where("type = 1 and id=?", id).First(&article)

	if result.Error != nil {
		Error(c)
		return
	}

	// 查询详情
	var articleDetail models.ArticleDetail
	models.DB.Select("content").Where("article_id=?", id).First(&articleDetail)

	// 查询列表
	var articleList []models.Article
	models.DB.Where("type = 1").Offset(0).Limit(10).Find(&articleList)

	c.HTML(http.StatusOK, "detail.html", RenderData(H{
		"head_title":  article.Title,
		"title":       article.Title,
		"detail":      articleDetail.Content,
		"articleList": articleList,
	}))

}
