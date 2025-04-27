package middlewares

import(
	"github.com/gin-gonic/gin"
	"net/http"
	"tll/utils"
)
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Check if the user is authenticated
		// If not, return an error response
		// If yes, continue to the next handler
		token:=ctx.GetHeader("Authorization")
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			ctx.Abort()
			return
		}
		username,err:= utils.ParseJWT(token)

		if err!=nil{
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			ctx.Abort()
			return
		}
		ctx.Set("username",username)
		ctx.Next()
	}
}

func Authorize(allowedRoles ...int) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.GetHeader("Authorization")

		if tokenStr == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "empty token,Unauthorized"})
			c.Abort()
			return
		}

		username, err := utils.ParseJWT(tokenStr)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()+"Unauthorized"})
			c.Abort()
			return
		}

		userRole, err := utils.GetUserRoleFromJWT(tokenStr)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "lower role,Unauthorized"})
			c.Abort()
			return
		}

		// 
		c.Set("username", username)
		c.Set("role", userRole)

		// 检查用户角色是否在允许的角色列表中
		for _, r := range allowedRoles {
			if userRole == r {
				c.Next()
				return
			}
		}

		// 如果角色不匹配，返回403 Forbidden
		c.JSON(http.StatusForbidden, gin.H{"error": "access denied"})
		c.Abort()
	}
}

func ExtractParamName() gin.HandlerFunc {
	return func(c *gin.Context) {
		courseName := c.Param("course_name")
		if courseName != "" {
			c.Set("course_name", courseName)
		}
		experimentName := c.Param("experiment_name")
		if experimentName != "" {
			c.Set("experiment_name", experimentName)
		}
		c.Next()
	}
}