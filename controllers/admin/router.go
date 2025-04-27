package admin

import (
	"tll/middlewares"
	"github.com/gin-gonic/gin"
)

func AdminRouterGroup(r *gin.RouterGroup) {
	admin := r.Group("/admin")
	admin.Use(middlewares.Authorize(0))
	{
		admin.POST("/Role", GiveTeacherrole)
	}
}


