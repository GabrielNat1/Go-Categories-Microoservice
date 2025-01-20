package use_cases

import (
	"github.com/Gabriel_devs/go-categories-msvc/internal/entities"
	"github.com/Gabriel_devs/go-categories-msvc/internal/repositories"
)

type ListCategoriesUsecase struct {
	repository repositories.ICategoryRepository
}

func NewListCategoriesUsecase(repository repositories.ICategoryRepository) ListCategoriesUsecase {
	return ListCategoriesUsecase{repository: repository}
}

func (u *ListCategoriesUsecase) Execute() ([]*entities.Category, error) {
	categories, err := u.repository.List()
	if err != nil {
		return nil, err
	}

	return categories, nil
}
