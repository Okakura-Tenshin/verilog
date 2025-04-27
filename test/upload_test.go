package test

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"tll/controllers/file"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestUploadStudentFilesHandler(t *testing.T) {
	// 创建一个临时文件
	tempFile, err := os.CreateTemp("", "testfile-*.txt")
	assert.NoError(t, err)
	defer os.Remove(tempFile.Name())

	// 写入内容
	_, err = tempFile.WriteString("test content")
	assert.NoError(t, err)
	tempFile.Close()

	// 构造 multipart/form-data 请求
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("files", filepath.Base(tempFile.Name()))
	assert.NoError(t, err)

	fileData, err := os.ReadFile(tempFile.Name())
	assert.NoError(t, err)

	part.Write(fileData)
	writer.Close()

	// 创建路由
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// mock 中间件设置上下文值
	router.POST("/course/:course_name/:experiment_name/UploadFiles", func(c *gin.Context) {
		// 手动设置 context 中的 username（模拟登录学生）
		c.Set("username", "student1")

		// 设置路径参数到 context（模拟你传进来的course和experiment）
		c.Set("course_name", c.Param("course_name"))
		c.Set("experiment_name", c.Param("experiment_name"))

		// 调用上传逻辑
		file.UploadStudentFilesHandler(c)
	})

	// 发送请求
	req := httptest.NewRequest(http.MethodPost, "127.0.0.1:11451/api/course/233/666/UploadFiles", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// 检查返回结果
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "上传成功")
}

