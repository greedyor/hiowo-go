package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// About
func About(c *gin.Context) {
	c.HTML(http.StatusOK, "about.html", gin.H{})
}
