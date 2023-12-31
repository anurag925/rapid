// Code generated by "enumer -type=AccountType -json -transform=snake -trimprefix=AccountType"; DO NOT EDIT.

package models

import (
	"encoding/json"
	"fmt"
	"strings"
)

const _AccountTypeName = "customerretaileradmin"

var _AccountTypeIndex = [...]uint8{0, 8, 16, 21}

const _AccountTypeLowerName = "customerretaileradmin"

func (i AccountType) String() string {
	if i < 0 || i >= AccountType(len(_AccountTypeIndex)-1) {
		return fmt.Sprintf("AccountType(%d)", i)
	}
	return _AccountTypeName[_AccountTypeIndex[i]:_AccountTypeIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _AccountTypeNoOp() {
	var x [1]struct{}
	_ = x[AccountTypeCustomer-(0)]
	_ = x[AccountTypeRetailer-(1)]
	_ = x[AccountTypeAdmin-(2)]
}

var _AccountTypeValues = []AccountType{AccountTypeCustomer, AccountTypeRetailer, AccountTypeAdmin}

var _AccountTypeNameToValueMap = map[string]AccountType{
	_AccountTypeName[0:8]:        AccountTypeCustomer,
	_AccountTypeLowerName[0:8]:   AccountTypeCustomer,
	_AccountTypeName[8:16]:       AccountTypeRetailer,
	_AccountTypeLowerName[8:16]:  AccountTypeRetailer,
	_AccountTypeName[16:21]:      AccountTypeAdmin,
	_AccountTypeLowerName[16:21]: AccountTypeAdmin,
}

var _AccountTypeNames = []string{
	_AccountTypeName[0:8],
	_AccountTypeName[8:16],
	_AccountTypeName[16:21],
}

// AccountTypeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func AccountTypeString(s string) (AccountType, error) {
	if val, ok := _AccountTypeNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _AccountTypeNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to AccountType values", s)
}

// AccountTypeValues returns all values of the enum
func AccountTypeValues() []AccountType {
	return _AccountTypeValues
}

// AccountTypeStrings returns a slice of all String values of the enum
func AccountTypeStrings() []string {
	strs := make([]string, len(_AccountTypeNames))
	copy(strs, _AccountTypeNames)
	return strs
}

// IsAAccountType returns "true" if the value is listed in the enum definition. "false" otherwise
func (i AccountType) IsAAccountType() bool {
	for _, v := range _AccountTypeValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for AccountType
func (i AccountType) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for AccountType
func (i *AccountType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("AccountType should be a string, got %s", data)
	}

	var err error
	*i, err = AccountTypeString(s)
	return err
}
