package routes

import (
	"example.com/todoServer/routes/v1"
	"github.com/gin-gonic/gin"
)

func InitalizeRoutes(r *gin.Engine) {
	v1Group := r.Group("/api/v1")
	{
		v1.InitalizeUserRoutes(v1Group)
	}
}
