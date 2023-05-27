package list

import "errors"

func validateTitle(title *string) (err error) {
	if title == nil || (title != nil && *title == "") {
		err = errors.New("invalid title: a task title can't be empty")
	}
	return
}
