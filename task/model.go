package task

import (
	"time"

	"github.com/coffemanfp/todo/utils"
)

type Task struct {
	ID             int        `json:"id,omitempty"`
	ListID         *int       `json:"list_id,omitempty"`
	Title          *string    `json:"title,omitempty"`
	Description    *string    `json:"description,omitempty"`
	IsAddedToMyDay *bool      `json:"is_added_to_my_day,omitempty"`
	IsImportant    *bool      `json:"is_important,omitempty"`
	Reminder       *time.Time `json:"reminder,omitempty"`
	DueDate        *time.Time `json:"due_date,omitempty"`
	Repeat         *string    `json:"repeat,omitempty"`
	CreatedAt      time.Time  `json:"created_at,omitempty"`
	CreatedBy      int        `json:"created_by,omitempty"`
}

func New(taskR Task) (task Task, err error) {
	err = validateCreator(taskR.CreatedBy)
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
	if taskR.Repeat != nil {
		task.Repeat = new(string)
		*task.Repeat = utils.RemoveSpaceAndConvertSpecialChars(*taskR.Repeat)
	}
	return
}
