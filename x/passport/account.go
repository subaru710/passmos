package passport

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/wire"
	"github.com/cosmos/cosmos-sdk/x/auth"
)

var _ auth.Account = (*PassportAccount)(nil)

// PassportAccount is a custom extension for this application. It is an example of
// extending auth.BaseAccount with custom fields. It is compatible with the
// stock auth.AccountMapper, since auth.AccountMapper uses the flexible go-amino
// library.
type PassportAccount struct {
	auth.BaseAccount
	ExternalStore
}

// NewPassportAccount returns a reference to a new PassportAccount given a name and an
// auth.BaseAccount.
func NewPassportAccount(pd PersonalData, baseAcct auth.BaseAccount, store ExternalStore) *PassportAccount {
	store.SetPersonalData(pd)
	return &PassportAccount{BaseAccount: baseAcct, ExternalStore: store}
}

// GetAccountDecoder returns the AccountDecoder function for the custom
// PassportAccount.
func GetAccountDecoder(cdc *wire.Codec) auth.AccountDecoder {
	return func(accBytes []byte) (auth.Account, error) {
		if len(accBytes) == 0 {
			return nil, sdk.ErrTxDecode("accBytes are empty")
		}

		acct := new(PassportAccount)
		err := cdc.UnmarshalBinaryBare(accBytes, &acct)
		if err != nil {
			panic(err)
		}

		return acct, err
	}
}
