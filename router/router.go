package router

import (
	"log"
	"tll/controllers/admin"
	"tll/controllers/auth"
	"tll/controllers/course"
    "tll/controllers/student"
	"tll/controllers/experiment"
	"tll/controllers/file"
	"tll/global"
	"tll/models"
    "tll/middlewares"
	"github.com/gin-gonic/gin"
)
type Register func(r *gin.RouterGroup)

var register = []Register{
	auth.AuthRouterGroup,
	admin.AdminRouterGroup,
    course.CourseRouterGroup,
    student.StudentGroupRouter,
    experiment.ExperimentGroupRouter,
    file.FileRouterGroup,
}

func InitRouterGroup()(r *gin.Engine){
	r=gin.Default()
	apiGroup:=r.Group("/api")
	apiGroup.Use(middlewares.LogFullRequest())
	for _,register:=range register{
		register(apiGroup)
	}
    if err := models.InitUser(global.Db); err != nil {
        log.Fatalf("初始化用户表失败: %v", err)
    }
    if err := models.InitCourse(global.Db); err != nil {
        log.Fatalf("初始化课程表失败: %v", err)
    }
    if err := models.InitExperiment(global.Db); err != nil {
        log.Fatalf("初始化实验表失败: %v", err)
    }
    if err := models.InitSubmitRecord(global.Db); err != nil {
        log.Fatalf("初始化提交记录表失败: %v", err)
    }

    if err := models.CreateAdminIfNotExists(global.Db); err != nil {
        log.Fatalf("初始化管理员失败: %v", err)
    }
    
    return r
}


/*
func AuthRouterGroup(r *gin.RouterGroup) {
    auth := r.Group("/auth")
    {
        auth.POST("/login", Login)
        auth.POST("/register", auth.Register)
    }

    teacher := r.Group("/teacher")
    teacher.Use(middlewares.Authorize(0, 1)) // 管理员和教师都可以访问
    {
        teacher.GET("/course", controllers.GetCourseList)
    }

    admin := r.Group("/admin")
    admin.Use(middlewares.Authorize(0)) // 只有管理员可以
    {
        admin.GET("/users", controllers.AdminUserList)
    }
}
*/
