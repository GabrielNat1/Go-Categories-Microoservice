package user_cases

import (
	"github.com/Gabriel_devs/go-categories-msvc/internal/entities"
	"github.com/Gabriel_devs/go-categories-msvc/internal/repositories"
)

type createCategoryUseCase struct {
	repository repositories.ICategoryRepository
}

func NewcreateCategoryUseCase(repository repositories.ICategoryRepository) createCategoryUseCase {
	return createCategoryUseCase{repository: repository}
}

func (u *createCategoryUseCase) Execute(name string) error {
	category, err := entities.NewCategory(name)
	if err != nil {
		return err
	}

	err = u.repository.Save(category)
	if err != nil {
		return err
	}

	return nil
}
