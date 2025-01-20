package controllers

import (
	"net/http"

	"github.com/Gabriel_devs/go-categories-msvc/internal/repositories"
	user_cases "github.com/Gabriel_devs/go-categories-msvc/internal/user-cases"
	"github.com/gin-gonic/gin"
)

type createCategoryInput struct {
	Name string `json:"name" binding:"required"`
}

func Createcategory(ctx *gin.Context, repository repositories.ICategoryRepository) {
	var body createCategoryInput

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"sucess": false,
				"error":  err.Error()})
		return
	}

	useCase := user_cases.NewcreateCategoryUseCase(repository)
	err := useCase.Execute(body.Name)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"sucess": false,
				"error":  err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"sucess": true})
}
