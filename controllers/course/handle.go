package course

import (
	"net/http"
	"tll/models"
	"github.com/gin-gonic/gin"
	"tll/global"
	"tll/config"
	"os"
	"path/filepath"
)

func CreateCourse(c *gin.Context) {
	var req struct {
		CourseName string `json:"course_name" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数不合法"})
		return
	}

	// 拿到传入的 username
	value, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法获取用户信息"})
		return
	}
	username := value.(string)
	course := models.Course{
		CourseName: req.CourseName,
		Teachers:   username,
	}
	//如果课程已存在，则返回错误
	if err:=models.GetCourseByName(global.Db,req.CourseName); err!=nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": "课程已存在"})
		return
	}
	// 使用username创建课程
	if err:=models.CreateCourse(global.Db,&course); err!=nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建课程失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "创建课程成功",
		"creator": username,
	})
}
//获取管理课程表
func GetTeacherCourses(c *gin.Context) {
	
	username := c.GetString("username")
	courses, err := models.GetCoursesByTeacher(global.Db, username); 
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询课程失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"courses": courses})
}

func GetStudent(c *gin.Context){
	CourseName:=c.GetString("course_name")

	students, err := models.GetStudentsByCourse(global.Db, CourseName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取学生失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"students": students})
}

func ApproveJoinCourse(c *gin.Context) {
	var req struct {
		StudentID  uint `json:"student_id" binding:"required"`
	}
	CourseName := c.Param("course_name") // 从 URL 路径中获取课程名 /:course_name/GetStudent
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	
	// 更新对应学生记录的状态
	if err := models.SQLApproveJoinCourse(global.Db,CourseName,req.StudentID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "批准失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "学生已成功加入课程"})
}

func GetPendingStudents(c *gin.Context) {
	CourseName:=c.GetString("course_name")
	var pending []models.StudentInfo
	pending,err:=models.GetPendingStudents(global.Db,CourseName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"pending": pending})
}
//创建实验
func CreateExperiment(c *gin.Context){
	var req struct{
		ExperimentName string `json:"experiment_name" binding:"required"`
		ExperimentDescription string `json:"experiment_description" binding:"required"`
		
	}
	if err:=c.ShouldBindJSON(&req);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	basePath:=config.AppConfig.FilePath.BasePath
	CourseName := c.Param("course_name")
	//实验路径，根目录＋课程名＋实验名

	// 拼接完整路径
	filePath := filepath.Join(basePath, CourseName, req.ExperimentName)

	// 创建目录（包括中间层级）
	if err := os.MkdirAll(filePath, 0755); err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusOK,gin.H{"message":"实验创建成功"})
}
