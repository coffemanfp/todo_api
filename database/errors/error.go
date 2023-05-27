package errors

import "fmt"

type Error struct {
	Type    string
	prefix  string
	content string
}

func (e Error) Error() string {
	return fmt.Sprintf("%s: %s", e.prefix, e.content)
}

func NewError(t, prefix, content string) Error {
	return Error{
		t, prefix, content,
	}
}
