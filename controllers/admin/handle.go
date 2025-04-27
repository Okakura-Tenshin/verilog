package admin
import (
	"net/http"
	"tll/global"
	"tll/models"

	"github.com/gin-gonic/gin"
)
//第一个功能，根据传入名字，给与教师权限1
//后续考虑更改，根据唯一性的ID查找，因为ID是唯一的，名字不一定
func GiveTeacherrole(c *gin.Context){
	var req struct{
		Username string `json:"username" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//根据名字找出人，然后赋权1
	if err := models.GiveTeacherRole(global.Db, req.Username); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "用户不存在或更新失败"})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "user role updated to teacher"})
}