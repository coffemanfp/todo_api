package account

import (
	"time"

	"github.com/coffemanfp/todo/utils"
)

// Account is the representation of the common account data
type Account struct {
	Name      string    `json:"name,omitempty"`
	LastName  string    `json:"last_name,omitempty"`
	Nickname  string    `json:"nickname,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// New initializes a new account based on the basic data provided from the account passed as param.
//
//	@param accountR Account: Basic data of the account to build.
//	@return account Account: Account builded
//	@return err error: error in the validation of the based account.
func New(accountR Account) (account Account, err error) {
	err = ValidateNickname(accountR.Nickname)
	if err != nil {
		return
	}
	account = accountR
	account.Password, err = utils.HashPassword(account.Password)
	return
}
