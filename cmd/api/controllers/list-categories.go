package controllers

import (
	"net/http"

	"github.com/Gabriel_devs/go-categories-msvc/internal/repositories"
	user_cases "github.com/Gabriel_devs/go-categories-msvc/internal/user-cases"
	"github.com/gin-gonic/gin"
)

func ListCategories(ctx *gin.Context, repository repositories.ICategoryRepository) {

	useCase := user_cases.NewListCategoriesUsecase(repository)
	categories, err := useCase.Execute()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"sucess": false,
				"error":  err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"sucess": true, "categories": categories})
}
