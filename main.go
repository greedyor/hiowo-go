package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"hiowo.com/config"
	"hiowo.com/models"
	"hiowo.com/routes"
)

func main() {

	// 日志记录
	logfile, err := os.Create("./log/gin_http_" + strconv.FormatInt(time.Now().Unix(), 10) + ".log")
	if err != nil {
		fmt.Println("Could not create log file")
	}
	gin.SetMode(gin.DebugMode)
	gin.DefaultWriter = io.MultiWriter(logfile)

	// GIN 初始化
	r := gin.Default()

	// 设置公共目录
	r.Static("/public/static", "./public/static")

	// 设置视图HTML资源目录
	r.LoadHTMLGlob("views/*")

	c := config.LoadMysqlConfig()

	// mysql 初始化
	models.Connect(c)

	// 设置路由
	r = routes.IndexRouter(r)
	r = routes.NoteRouter(r)
	r = routes.SarscovRouter(r)
	r.Run(":8080")

}
