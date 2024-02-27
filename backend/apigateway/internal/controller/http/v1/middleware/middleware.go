package middleware

import (
	"gateway/pkg/jwt"
	"strings"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// import (
// 	configs "gateway/internal/config"

// 	"github.com/gin-gonic/gin"
// )

// type Middleware struct {
// 	app *gin.Engine
// 	config *configs.Config
// }
// func SetUpMiddleware(
// 	app *gin.Engine,
// 	config *configs.Config,
// 	)	{
// 		middleware := &Middleware{
// 			app: app,
// 			config: config,
// 		}
// 		middleware.gin
// 	}

func Middleware(next gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		accessToken:=c.GetHeader("Authorization")
		if accessToken==""{
			c.JSON(401,gin.H{
				"message":"Unauthorized",
			})
			return
		}
		if !strings.HasPrefix(accessToken, "Bearer ") {
			c.JSON(401,gin.H{
				"message":"Unauthorized",
			})
		}

		// Extract the token from the header
		token := strings.TrimPrefix(accessToken, "Bearer ")

		claims, err := jwt.VerifyToken(token)
		if err != nil {
			c.JSON(401,gin.H{
				"error":err.Error(),
			})			
		}
		id, ok := claims["id"].(string)
		if !ok {
			c.JSON(401,gin.H{
				"message":"Unauthorized",
			})
		}
		role, ok := claims["role"].(string)
		if !ok {
			c.JSON(401,gin.H{
				"message":"Unauthorized",
			})
		}
		session.Set("id", id)
		session.Set("role", role)
		session.Save()
		next(c)
	}

}