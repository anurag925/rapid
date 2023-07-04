//go:generate go run github.com/dmarkham/enumer -type=AccountType -json -transform=snake -trimprefix=AccountType
//go:generate go run github.com/dmarkham/enumer -type=AccountStatus -json -transform=snake -trimprefix=AccountStatus
package models

import (
	"encoding/json"
	"time"

	"github.com/uptrace/bun"
	"gopkg.in/guregu/null.v4"
)

type AccountType int8

const (
	AccountTypeCustomer AccountType = iota
	AccountTypeRetailer
	AccountTypeAdmin
)

type AccountStatus int8

const (
	AccountStatusCreated AccountStatus = iota
	AccountStatusOnboarding
	AccountStatusActive
	AccountStatusBlocked
)

type Account struct {
	bun.BaseModel
	ID           int64         `bun:"id,pk,autoincrement"`
	CreatedAt    time.Time     `bun:",nullzero,notnull,default:current_timestamp" json:"-"`
	UpdatedAt    time.Time     `bun:",nullzero,notnull,default:current_timestamp" json:"-"`
	Type         AccountType   `bun:",notnull" json:"type"`
	Status       AccountStatus `bun:",notnull" json:"status"`
	FirstName    null.String   `json:"first_name"`
	LastName     null.String   `json:"last_name"`
	Email        string        `bun:",notnull,unique" json:"email"`
	MobileNumber string        `bun:",notnull,unique" json:"mobile_number"`
	Password     null.String   `json:"password"`
}

func (a Account) MarshalJSON() ([]byte, error) {
	type PasswordMaskedAccount Account
	passwordMaskedAccount := PasswordMaskedAccount(a)
	if passwordMaskedAccount.Password.Valid {
		passwordMaskedAccount.Password.SetValid("xxxxxxxxxxxxxxxxx")
	}
	return json.Marshal(passwordMaskedAccount)
}
