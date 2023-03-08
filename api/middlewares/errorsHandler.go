package middlewares

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JSONAppErrorReporter() gin.HandlerFunc {
	return jsonAppErrorReporterT(gin.ErrorTypeAny)
}
// TODO: handle error 
func jsonAppErrorReporterT(errType gin.ErrorType) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if (len(c.Errors) != 0) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"success": false, "error": "username already sign up"})
			return 
		}
		fmt.Println("SUCCESS")
	}
}