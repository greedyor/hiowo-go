package models

// Article model
type Article struct {
	ID         int    `json:"id" binding:"required"`
	Type       int    `json:"type" binding:"required"`
	Title      string `json:"title"  binding:"required"`
	CreateTime string `json:"create_time"`
}

// 自定义表名称
func (Article) TableName() string {
	return "t_article"
}
