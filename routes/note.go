package routes

import (
	"github.com/gin-gonic/gin"
	"hiowo.com/controllers"
)

// NoteRouter created auth router
func NoteRouter(r *gin.Engine) *gin.Engine {
	r.GET("/note/detail", controllers.NoteDetail)
	r.GET("/note", controllers.Note)

	return r
}
