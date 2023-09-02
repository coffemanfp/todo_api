package task

import "time"

type Category struct {
	ID        int       `json:"id,omitempty"`
	CreatedBy int       `json:"account_id,omitempty"`
	Name      *string   `json:"name,omitempty"`
	Color     *string   `json:"color,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

var categoryColors = map[string]bool{
	"red":    true,
	"blue":   true,
	"yellow": true,
	"white":  true,
	"black":  true,
	"orange": true,
	"purple": true,
}

func NewCategory(c Category) (category Category, err error) {
	err = validateCreator(c.CreatedBy)
	if err != nil {
		return
	}
	err = validateTitle(c.Name)
	if err != nil {
		return
	}
	err = validateColor(c.Color)
	if err != nil {
		return
	}

	category.CreatedBy = c.CreatedBy
	category.Name = c.Name
	category.Color = c.Color
	c.CreatedAt = time.Now()
	return
}

func UpdateCategory(c Category) (category Category, err error) {
	if c.Name != nil {
		err = validateTitle(c.Name)
		if err != nil {
			return
		}
	}
	if c.Color != nil {
		err = validateColor(c.Color)
		if err != nil {
			return
		}
	}

	category.Name = c.Name
	category.Color = c.Color
	return
}
