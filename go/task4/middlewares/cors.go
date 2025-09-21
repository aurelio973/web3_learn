package middlewares

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// 跨链中间件配置
func CORSMiddleware() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:   []string{"https://prod.com","http://localhost:8080"},
		AllowMethods:   []string{"GET","POST","PUT","PATCH"},
		AllowHeaders:   []string{"Origin","Content-Type","Authorization"},
		ExposeHeaders:  []string{"Content-Length"},
		AllowCredentials:true,
		MaxAge:          12*time.Hour,      
	})
}
