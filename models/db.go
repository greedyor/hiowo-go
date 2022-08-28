package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"webkodes.com/admin/config"
)

func DB() *gorm.DB {

	// 获取配置
	Config := config.LoadMysqlConfig()

	// 配置 数据库
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s",
		Config.MySQL.Username,
		Config.MySQL.Password,
		Config.MySQL.Host,
		Config.MySQL.Port,
		Config.MySQL.Database,
		Config.MySQL.Charset)

	// 连接数据库
	MysqlDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   Config.MySQL.Prefix,
			SingularTable: true,
		},
	})
	if err != nil {
		fmt.Println("failed to connect database:", err)
	}

	return MysqlDB
}
