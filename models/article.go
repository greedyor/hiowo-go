package models

// Article eporting model for joke
type Article struct {
	ID         int    `json:"id" binding:"required"`
	Type       int    `json:"type" binding:"required"`
	Title      string `json:"title"  binding:"required"`
	CreateTime string `json:"create_time"`
}

func (Article) TableName() string {
	//实现TableName接口，以达到结构体和表对应，如果不实现该接口，并未设置全局表名禁用复数，gorm会自动扩展表名为articles（结构体+s）
	return "ob_article"
}
