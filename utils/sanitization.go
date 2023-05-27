package utils

import (
	"html"
	"strings"
)

func RemoveSpaceAndConvertSpecialChars(s string) string {
	return html.EscapeString(strings.TrimSpace(s))
}
