package routes

import (
	"github.com/gin-gonic/gin"
	"hiowo.com/controllers"
)

// SarscovRouter created auth router
func SarscovRouter(r *gin.Engine) *gin.Engine {

	r.GET("/sarscov", controllers.Sarscov)

	return r
}
