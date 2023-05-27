package task

import (
	"time"

	"github.com/coffemanfp/todo/utils"
)

type Task struct {
	ID          int       `json:"id,omitempty"`
	ListID      int       `json:"list_id,omitempty"`
	Title       *string   `json:"title,omitempty"`
	Description *string   `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	CreatedBy   int       `json:"created_by,omitempty"`
}

func New(taskR Task) (task Task, err error) {
	err = validateListID(taskR.ListID)
	if err != nil {
		return
	}
	err = validateTitle(taskR.Title)
	if err != nil {
		return
	}
	task = populatePointerValues(taskR)
	task.CreatedAt = time.Now()
	return
}

func Update(taskR Task) (task Task, err error) {
	err = validateTitle(taskR.Title)
	if err != nil {
		return
	}
	task = populatePointerValues(taskR)
	return
}

func populatePointerValues(taskR Task) (task Task) {
	task = taskR
	if taskR.Title != nil {
		task.Title = new(string)
		*task.Title = utils.RemoveSpaceAndConvertSpecialChars(*taskR.Title)
	}
	if taskR.Description != nil {
		task.Description = new(string)
		*task.Description = utils.RemoveSpaceAndConvertSpecialChars(*taskR.Description)
	}
	return
}
