package main

import (
	"github.com/Gabriel_devs/go-categories-msvc/cmd/api/controllers"
	"github.com/gin-gonic/gin"
)

func CategoryRoutes(r *gin.Engine) {
	categoryRoutes := r.Group("/categories")

	categoryRoutes.POST("/", controllers.Createcategory)
}
