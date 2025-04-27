package student

import (
	"tll/middlewares"

	"github.com/gin-gonic/gin"
)
func StudentGroupRouter(c *gin.RouterGroup){
	student:=c.Group("/student")
	student.Use(middlewares.Authorize(0,1,2))
	{
		student.POST("/JoinCourse",JoinCourse)
	}
}
