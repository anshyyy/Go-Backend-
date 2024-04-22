package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitalizeUserRoutes(r *gin.RouterGroup) {
	userGroup := r.Group("/user")
	{
		userGroup.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Success",
				"data":    "yo!!",
				"success": true,
			})
		})
	}
}
