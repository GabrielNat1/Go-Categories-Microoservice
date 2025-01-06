package entities

import (
	"fmt"
	"time"
)

type Category struct {
	ID       uint      `json:"id"`
	Name     string    `json:"name"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}

func NewCategory(name string) (*Category, error) {
	category := &Category{
		Name:     name,
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}

	err := category.IsValid()

	if err != nil {
		return nil, err
	}

	return category, nil
}

func (c *Category) IsValid() error {
	if len(c.Name) < 5 {
		return fmt.Errorf("name must be at least 5 characters. got %d", len(c.Name))
	}
	return nil
}
