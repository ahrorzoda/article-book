package cors

import "github.com/gin-gonic/gin"

func CorsMiddleware(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Next()
}
