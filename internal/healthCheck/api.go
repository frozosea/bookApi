package healthCheck

import "github.com/gin-gonic/gin"

type HealthCheck struct {
}

func (s *HealthCheck) HealthCheck(ctx *gin.Context) {
	ctx.JSON(200, gin.H{`success`: true})
}
