package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Error
func Error(c *gin.Context) {
	c.HTML(http.StatusOK, "error.html", H{})
}
