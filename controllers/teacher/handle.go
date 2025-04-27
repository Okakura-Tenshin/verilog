package teacher
import(
	"net/http"
	"tll/global"
	"tll/models"
	"github.com/gin-gonic/gin"
)

func ApproveJoinCourse(c *gin.Context) {
	var req struct {
		CourseName string `json:"course_name" binding:"required"`
		StudentID  string `json:"student_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	
	// 更新对应学生记录的状态
	err := global.Db.Model(&models.SelectCourse{}).
		Where("course_name = ? AND student_id = ?", req.CourseName, req.StudentID).
		Update("student_status", true).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "批准失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "学生已成功加入课程"})
}

func GetPendingStudents(c *gin.Context) {
	courseName := c.Query("course_name")
	if courseName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "缺少课程名"})
		return
	}

	var pending []models.SelectCourse
	err := global.Db.Preload("Student").
		Where("course_name = ? AND student_status = ?", courseName, false).
		Find(&pending).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"pending": pending})
}
