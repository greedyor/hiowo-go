package routes

import (
	"github.com/gin-gonic/gin"
	"webkodes.com/admin/controllers"
)

// IndexRouter created auth router
func SarscovRouter(r *gin.Engine) *gin.Engine {

	r.GET("/sarscov", controllers.Sarscov)

	return r
}
