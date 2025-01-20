package repositories

import "github.com/Gabriel_devs/go-categories-msvc/internal/entities"

type ICategoryRepository interface {
	Save(category *entities.Category) error
	List() ([]*entities.Category, error)
}
