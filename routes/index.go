package routes

import (
	"github.com/gin-gonic/gin"
	"hiowo.com/controllers"
)

// IndexRouter created auth router
func IndexRouter(r *gin.Engine) *gin.Engine {

	r.GET("/", controllers.Index)
	r.GET("/index", controllers.Index)
	r.GET("/about", controllers.About)
	r.GET("/download", controllers.Index)

	return r
}
