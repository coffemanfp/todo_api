package auth

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword uses the bcrypt algorithm to encrypt the password provided.
//
//	@param password string: password to encrypt.
//	@return $1 string: password encrypted.
//	@return $2 error: bcrypt encryptation error.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		err = fmt.Errorf("failed to generate password: %s", err)
	}
	return string(bytes), err
}
