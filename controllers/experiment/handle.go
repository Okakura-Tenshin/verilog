package experiment

import (
	"net/http"
	"os"
	"path/filepath"
	"github.com/gin-gonic/gin"
	"tll/config"
	"tll/models"
	"tll/global"
)
//创建一个实验
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
	CourseName := c.GetString("course_name")
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

func GetExperimentResult(c *gin.Context){
	experimentName := c.GetString("experiment_name")
	courseName := c.GetString("course_name")
	records,err:=models.GetExperimentResult(global.Db,experimentName,courseName)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusOK,gin.H{"records":records})
}