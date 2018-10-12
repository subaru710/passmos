package passport

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// MsgCreate - create a passport with an external address
type MsgCreate struct {
	Address sdk.AccAddress `json:"address"`
	Data    PersonalData   `json:"data"`
}

var _ sdk.Msg = MsgCreate{}

// NewMsgCreate - construct create msg.
func NewMsgCreate(addr sdk.AccAddress, data PersonalData) MsgCreate {
	return MsgCreate{Address: addr, Data: data}
}

// Implements Msg.
func (msg MsgCreate) Type() string { return "passport" } // TODO: "passport/create"

// Implements Msg.
func (msg MsgCreate) ValidateBasic() sdk.Error {
	if len(msg.Address) == 0 {
		return sdk.ErrInvalidAddress(msg.Address.String())
	}
	//TODO: validate data
	return nil
}

// Implements Msg.
func (msg MsgCreate) GetSignBytes() []byte {
	bin, err := msgCdc.MarshalJSON(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(bin)
}

// Implements Msg.
func (msg MsgCreate) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Address}
}

// MsgUpdate - update personal data for a passport address
type MsgUpdate struct {
	Address sdk.AccAddress `json:"address"`
	Data    PersonalData   `json:"data"`
}

var _ sdk.Msg = MsgUpdate{}

// NewMsgUpdate - construct update msg.
func NewMsgUpdate(addr sdk.AccAddress, data PersonalData) MsgUpdate {
	return MsgUpdate{Address: addr, Data: data}
}

// Implements Msg.
func (msg MsgUpdate) Type() string { return "passport" } // TODO: "passport/update"

// Implements Msg.
func (msg MsgUpdate) ValidateBasic() sdk.Error {
	if len(msg.Address) == 0 {
		return sdk.ErrInvalidAddress(msg.Address.String())
	}
	//TODO: validate data
	return nil
}

// Implements Msg.
func (msg MsgUpdate) GetSignBytes() []byte {
	bin, err := msgCdc.MarshalJSON(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(bin)
}

// Implements Msg.
func (msg MsgUpdate) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Address}
}

// MsgAuthorize - create a passport with an external address
// TODO: partial fields authorization
type MsgAuthorize struct {
	Address  sdk.AccAddress `json:"address"`
	Receiver sdk.AccAddress `json:"receiver"`
}

var _ sdk.Msg = MsgAuthorize{}

// NewMsgAuthorize - construct create msg.
func NewMsgAuthorize(addr sdk.AccAddress, receiver sdk.AccAddress) MsgAuthorize {
	return MsgAuthorize{Address: addr, Receiver: receiver}
}

// Implements Msg.
func (msg MsgAuthorize) Type() string { return "passport" } // TODO: "passport/authorize"

// Implements Msg.
func (msg MsgAuthorize) ValidateBasic() sdk.Error {
	if len(msg.Address) == 0 {
		return sdk.ErrInvalidAddress(msg.Address.String())
	}
	if len(msg.Receiver) == 0 {
		return sdk.ErrInvalidAddress(msg.Address.String())
	}
	return nil
}

// Implements Msg.
func (msg MsgAuthorize) GetSignBytes() []byte {
	bin, err := msgCdc.MarshalJSON(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(bin)
}

// Implements Msg.
func (msg MsgAuthorize) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Address}
}
