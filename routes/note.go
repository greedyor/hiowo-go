package routes

import (
	"github.com/gin-gonic/gin"
	"webkodes.com/admin/controllers"
)

// IndexRouter created auth router
func NoteRouter(r *gin.Engine) *gin.Engine {
	r.GET("/note/detail", controllers.NoteDetail)
	r.GET("/note", controllers.Note)

	return r
}
