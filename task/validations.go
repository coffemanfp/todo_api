package task

import (
	"errors"
	"fmt"
	"strings"
)

func validateListID(listID int) (err error) {
	if listID <= 0 {
		err = fmt.Errorf("invalid list id or not provided: %d", listID)
	}
	return
}

func validateTitle(title *string) (err error) {
	if title == nil || (title != nil && *title == "") {
		err = errors.New("invalid title: a task title can't be empty")
	}
	return
}

func validateCreator(createdby int) (err error) {
	if createdby <= 0 {
		err = fmt.Errorf("invalid creator id or not provided: %d", createdby)
	}
	return
}

func validateColor(color *string) (err error) {
	if color == nil || (color != nil && *color == "") {
		err = errors.New("invalid color: the color can't be empty")
	}
	if _, ok := categoryColors[strings.ToLower(*color)]; !ok {
		err = fmt.Errorf("invalid color: invalid color name %s", *color)
	}
	return
}
