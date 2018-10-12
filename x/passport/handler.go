package passport

import (
	"reflect"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewHandler returns a handler for "passport" type messages.
func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case MsgCreate:
			return handleMsgCreate(ctx, k, msg)
		case MsgUpdate:
			return handleMsgUpdate(ctx, k, msg)
		case MsgAuthorize:
			return handleMsgAuthorize(ctx, k, msg)
		default:
			errMsg := "Unrecognized passport Msg type: " + reflect.TypeOf(msg).Name()
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

// Handle MsgCreate.
func handleMsgCreate(ctx sdk.Context, k Keeper, msg MsgCreate) sdk.Result {
	record, err := k.CreatePassport(ctx, msg.Address, msg.Data)
	if err != nil {
		return err.Result()
	}
	data := k.cdc.MustMarshalBinaryBare(record)
	return sdk.Result{
		Code: sdk.ABCICodeOK,
		Data: data,
	}
}

// Handle MsgUpdate.
func handleMsgUpdate(ctx sdk.Context, k Keeper, msg MsgUpdate) sdk.Result {
	record, err := k.UpdatePassport(ctx, msg.Address, msg.Data)
	if err != nil {
		return err.Result()
	}
	data := k.cdc.MustMarshalBinaryBare(record)
	return sdk.Result{
		Code: sdk.ABCICodeOK,
		Data: data,
	}
}

// Handle MsgAuthorize.
func handleMsgAuthorize(ctx sdk.Context, k Keeper, msg MsgAuthorize) sdk.Result {
	err := k.AuthorizePassport(ctx, msg.Address, msg.Receiver)
	if err != nil {
		return err.Result()
	}
	return sdk.Result{
		Code: sdk.ABCICodeOK,
	}
}
