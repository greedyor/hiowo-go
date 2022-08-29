package models

// ArticleDetail model
type ArticleDetail struct {
	ID         int    `json:"id" binding:"required"`
	ArticleId  int    `json:"article_id" binding:"required"`
	Content    string `json:"content"  binding:"required"`
	CreateTime string `json:"create_time"`
}

// 自定义表名称
func (ArticleDetail) TableName() string {
	return "ob_article_detail"
}
