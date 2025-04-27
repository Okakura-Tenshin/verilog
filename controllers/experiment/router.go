package experiment

import (
	"tll/middlewares"

	"github.com/gin-gonic/gin"
)
func ExperimentGroupRouter(c *gin.RouterGroup){
	experiment:=c.Group("/course/:course_name/:experiment_name")
	experiment.Use(middlewares.Authorize(0,1)).Use(middlewares.ExtractParamName())
	{
		experiment.GET("/result",GetExperimentResult)
	}
	
}