package file

import (
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"tll/config"
	"tll/global"
	"tll/models"
	"github.com/gin-gonic/gin"
	"bytes"
	"encoding/json"
)

func UploadStudentFilesHandler(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "上传表单解析失败"})
		return
	}

	files := form.File["files"]
	if len(files) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "没有上传文件"})
		return
	}

	// 获取上下文的studentname
	studentName := c.GetString("username")
	studentID, err := models.FindUserIDByUsername(global.Db, studentName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if studentID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未认证学生"})
		return
	}

	// 文件存储目录：/根目录/课程名/实验名/学生ID/
	courseName := c.GetString("course_name")
	experimentName := c.GetString("experiment_name")
	basePath := filepath.Join(config.AppConfig.FilePath.BasePath, courseName, experimentName, strconv.FormatUint(uint64(studentID), 10))
	if err := os.MkdirAll(basePath, os.ModePerm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建学生目录失败"})
		return
	}

	uploaded := []string{}
	for _, file := range files {
		dst := basePath + "/" + file.Filename
		if err := c.SaveUploadedFile(file, dst); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "保存文件失败", "file": file.Filename})
			return
		}
		uploaded = append(uploaded, file.Filename)
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "上传成功",
		"files":    uploaded,
		"student":  studentID,
		"savePath": basePath,
	})
}

func GetResult(c *gin.Context) {
	var req struct {
		Pushfile string `json:"pushfile" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	courseName := c.GetString("course_name")

	experimentName := c.GetString("experiment_name")

	studentName := c.GetString("username")
	studentID, err := models.FindUserIDByUsername(global.Db, studentName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	//config.AppConfig.FilePath.BasePath有问题，不是绝对路径
	localPath,_:=os.Getwd()
	basePath := filepath.Join(localPath, config.AppConfig.FilePath.BasePath, courseName, experimentName,)
	verilogPath := filepath.Join(basePath, strconv.FormatUint(uint64(studentID), 10),req.Pushfile)
	standardPath := filepath.Join(basePath, "standard.vcd")
	demoPath := filepath.Join(basePath, "demo.v")
	
	payload := map[string]string{
		"received":  verilogPath,
		"testbench": demoPath,
		"standard":  standardPath,
	}
	
	jsonValue, _ := json.Marshal(payload)
	resp, err := http.Post("http://127.0.0.1:5000/run_test", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	var result struct {
		Match int `json:"match"`
	}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()+"don't have this file"})
		return
	}
	if err:=models.AddSubmitRecord(global.Db, studentID, studentName, experimentName, courseName, result.Match);err!=nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"match": result.Match})
}
