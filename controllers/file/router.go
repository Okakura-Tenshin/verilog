package file
import(
	"github.com/gin-gonic/gin"
	"tll/middlewares"
)
func FileRouterGroup(r *gin.RouterGroup){
	file:=r.Group("/course/:course_name/:experiment_name")
	file.Use(middlewares.ExtractParamName()).Use(middlewares.Authorize(0,1,2))
	{
		file.POST("/UploadFiles",UploadStudentFilesHandler)
		file.POST("/GetResult",GetResult)
	}
}
