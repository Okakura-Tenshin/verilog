package student

import (
	"net/http"
	"tll/global"
	"tll/models"
	"github.com/gin-gonic/gin"
)


//申请加课
func JoinCourse(c *gin.Context) {
	var req struct {
		CourseName string `json:"course_name" binding:"required"`
	}

	usernameRaw, ok := c.Get("username")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未获得学生名"})
		return
	}
	username := usernameRaw.(string)

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 查询当前学生的ID
	var student models.User
	if err := global.Db.Where("username = ?", username).First(&student).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "学生信息未找到"})
		return
	}
	
	if err:=models.JoinCourse(global.Db,req.CourseName,student.ID);err!=nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": "申请加入课程失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "已发送加课申请，等待老师批准"})
}

