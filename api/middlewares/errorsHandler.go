package middlewares

import (
	"log"

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
			// c.Errors.Last()
			for i, ginErr := range c.Errors {
				// log error
				log.Println(i, ginErr)
				// return last error obj pushed
			}
		}
		
	}
}