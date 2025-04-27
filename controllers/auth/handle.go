package auth

import (
	"net/http"
	"tll/global"
	"tll/models"
	"tll/utils"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hash := utils.HashPasswd(req.Password)
	
	user := models.User{
		Username: req.Username,
		Password: hash,
		Role: 2,
	}

	// 检查用户名是否已存在
	var existingUser models.User
	if err := global.Db.Where("username = ?", user.Username).First(&existingUser).Error; err == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "username already exists"})
		return
	}

	// 创建用户
	if err := models.CreateUser(global.Db, &user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 生成JWT令牌

	token, err := utils.GenerateJWT(user.Username,2)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id":       user.ID,
		"token":    token,
		"username": user.Username,
		"role":     user.Role,
	})
}

func Login(ctx *gin.Context) {
	var input struct {
		Name     string `json:"username"`
		Password string `json:"password"`
	}
	if err := ctx.ShouldBind(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var user models.User
	if err := global.Db.Where("username = ?", input.Name).First(&user).Error; err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid username "})
		return
	}

	if !utils.Cheakpasswd(input.Password, user.Password) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid  password"})
		return
	}

	token, err := utils.GenerateJWT(user.Username,user.Role)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}


	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
