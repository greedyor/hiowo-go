package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

// Sarscov
func Sarscov(c *gin.Context) {
	response, err := http.Get("http://trip.zhongan.com/api/epidemic/queryCurrentData")
	if err != nil || response.StatusCode != http.StatusOK {
		c.Status(http.StatusServiceUnavailable)
	}
	bytes, err := ioutil.ReadAll(response.Body)
	response.Body.Close()

	if err != nil {
		fmt.Println("ioutil.ReadAll err=", err)
	}

	json := url.QueryEscape(string(bytes))

	c.HTML(http.StatusOK, "sarscov.html", RenderData(H{
		"json": json,
	}))
}
