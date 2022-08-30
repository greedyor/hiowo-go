package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"hiowo.com/config"
)

var (
	DB *gorm.DB
)

// 数据库连接方法
func Connect(c *config.Config) {

	m := c.MySQL

	// 配置 数据库
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s",
		m.Username,
		m.Password,
		m.Host,
		m.Port,
		m.Database,
		m.Charset)

	// 连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   m.Prefix,
			SingularTable: true,
		},
	})
	if err != nil {
		panic(fmt.Errorf("failed to connect database: %v", err))
	}

	DB = db
}
