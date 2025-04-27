package course

import (
	"tll/middlewares"

	"github.com/gin-gonic/gin"
)

func CourseRouterGroup(r *gin.RouterGroup){
	course:=r.Group("/course")
	teacher:=course.Group("/")
	teacher.Use(middlewares.Authorize(0,1)).Use(middlewares.ExtractParamName())
	{
		course.POST("/CreateCourse",CreateCourse)

		course.GET("/GetTeacherCourses",GetTeacherCourses)

		course.GET("/:course_name/GetStudent",GetStudent)

		course.GET("/:course_name/GetPendingStudents",GetPendingStudents)

		course.POST("/:course_name/ApproveJoinCourse",ApproveJoinCourse)
		
		course.POST("/:course_name/CreateExperiment",CreateExperiment)
	}
	all:=course.Group("/")
	all.Use(middlewares.Authorize(0,1,2)).Use(middlewares.ExtractParamName())
	{
		
	}	
}

