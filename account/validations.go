package account

import (
	"net/http"
	"regexp"

	"github.com/coffemanfp/todo/errors"
)

func ValidateCredentials(account Account) (err error) {
	if account.Email != "" {
		err = ValidateEmail(account.Email)
	} else if account.Nickname != "" {
		err = ValidateNickname(account.Nickname)
	}
	return
}

var nicknameRegex = regexp.MustCompile(`^[a-z0-9_-]{3,32}$`)

// ValidateNickname validate the nickname with a regular expression.
//
//	@param nickname string: nickname to validate.
//	 @return err error: don't match the regex with the string provided.
func ValidateNickname(nickname string) (err error) {
	if !nicknameRegex.MatchString(nickname) {
		err = errors.NewHTTPError(http.StatusBadRequest, "invalid nickname: invalid nickname format of %s", nickname)
	}
	return
}

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

// ValidateEmail validate the email with a regular expression.
//
//	@param email string: email to validate.
//	 @return err error: don't match the regex with the string provided.
func ValidateEmail(email string) (err error) {
	if !emailRegex.MatchString(email) {
		err = errors.NewHTTPError(http.StatusBadRequest, "invalid email: invalid email format of %s", email)
	}
	return
}
