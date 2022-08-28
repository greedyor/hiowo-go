package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	MySQL struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		Database string `mapstructure:"database"`
		Charset  string `mapstructure:"charset"`
		Prefix   string `mapstructure:"prefix"`
	}
}

func LoadMysqlConfig() *Config {

	v := viper.New()
	v.SetConfigType("yaml")     // 指定文件格式
	v.SetConfigName("config")   // 文件名不用加后缀
	v.AddConfigPath("./config") // 路径

	err := v.ReadInConfig() // 查找并读取配置文件
	if err != nil {         // 处理读取配置文件的错误
		panic(fmt.Errorf("Fatal error config file: %s ", err))
	}

	var Conf = new(Config)

	if err := v.Unmarshal(&Conf); err != nil {
		panic(fmt.Errorf("Unmarshal: %v", err))
	}

	return Conf
}
