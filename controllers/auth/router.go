package auth

import (
	"github.com/gin-gonic/gin"
)

func AuthRouterGroup(r *gin.RouterGroup){
	auth:=r.Group("/account")
	{
		auth.POST("/login",Login)
		auth.POST("/register",Register)
	}
}
