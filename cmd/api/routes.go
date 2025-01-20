package main

import (
	"github.com/Gabriel_devs/go-categories-msvc/cmd/api/controllers"
	"github.com/Gabriel_devs/go-categories-msvc/internal/repositories"
	"github.com/gin-gonic/gin"
)

func CategoryRoutes(r *gin.Engine) {
	categoryRoutes := r.Group("/categories")

	inMemoryCategoryRepository := repositories.NewInMemoryCategoryRepository()
	categoryRoutes.POST("/", func(ctx *gin.Context) {
		controllers.Createcategory(ctx, inMemoryCategoryRepository)
	})
}
