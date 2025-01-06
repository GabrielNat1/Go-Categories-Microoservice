package user_cases

import (
	"log"

	"github.com/Gabriel_devs/go-categories-msvc/internal/entities"
)

type createCategoryUseCase struct {
}

func NewcreateCategoryUseCase() createCategoryUseCase {
	return createCategoryUseCase{}
}

// Corrigido para retornar um error
func (u *createCategoryUseCase) Execute(name string) error {
	category, err := entities.NewCategory(name)
	if err != nil {
		return err
	}

	log.Println(category)
	return nil
}
