package list

import (
	"time"

	"github.com/coffemanfp/todo/utils"
)

type List struct {
	ID                   int       `json:"id,omitempty"`
	Title                *string   `json:"title,omitempty"`
	Description          *string   `json:"description,omitempty"`
	BackgroundPictureURL *string   `json:"background_picture_url,omitempty"`
	CreatedBy            int       `json:"created_by,omitempty"`
	CreatedAt            time.Time `json:"created_at,omitempty"`
}

func New(listR List) (list List, err error) {
	err = validateTitle(listR.Title)
	if err != nil {
		return
	}
	list = populatePointerValues(listR)
	list = setDefaultValuesIfEmpty(list)
	list.CreatedAt = time.Now()
	return
}

func Update(listR List) (list List, err error) {
	err = validateTitle(listR.Title)
	if err != nil {
		return
	}
	list = populatePointerValues(listR)
	list = setDefaultValuesIfEmpty(list)
	return
}

func setDefaultValuesIfEmpty(listR List) (list List) {
	list = listR
	if listR.BackgroundPictureURL == nil || *listR.BackgroundPictureURL == "" {
		// Example data. This has to be modified
		list.BackgroundPictureURL = new(string)
		*list.BackgroundPictureURL = "https://picsum.photos/200/300"
	}
	return
}

func populatePointerValues(listR List) (list List) {
	list = listR
	if listR.Title != nil {
		list.Title = new(string)
		*list.Title = utils.RemoveSpaceAndConvertSpecialChars(*listR.Title)
	}
	if listR.Description != nil {
		list.Description = new(string)
		*list.Description = utils.RemoveSpaceAndConvertSpecialChars(*listR.Description)
	}
	return
}
