package passport

import (
	"testing"

	"github.com/stretchr/testify/require"

	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/wire"
	auth "github.com/cosmos/cosmos-sdk/x/auth"
)

func TestPassportHandler(t *testing.T) {
	ms, passKey := setupMultiStore()
	cdc := wire.NewCodec()
	auth.RegisterBaseAccount(cdc)

	ctx := sdk.NewContext(ms, abci.Header{}, false, log.NewNopLogger())
	keeper := NewKeeper(passKey, NewIpfsStore("https://ipfs.infura.io:5001"))

	handler := NewHandler(keeper)

	addr := sdk.AccAddress([]byte("test"))

	msg := NewMsgCreate(addr, PersonalData{Name: "test"})
	result := handler(ctx, msg)
	require.NotNil(t, result)
	require.Equal(t, sdk.ABCICodeOK, result.Code)
	require.NotNil(t, result.Data)
	var record Record
	cdc.MustUnmarshalBinaryBare(result.Data, &record)
	require.NotNil(t, &record)
	require.Equal(t, "ipfs", record.Type)
	require.NotNil(t, record.Timestamp)
	require.NotNil(t, record.Path)
	t.Logf("Path: %s\n", record.Path)

	msg2 := NewMsgCreate(addr, PersonalData{Name: "test2"})
	result = handler(ctx, msg2)
	require.NotEqual(t, sdk.ABCICodeOK, result.Code)

	msg3 := NewMsgUpdate(addr, PersonalData{Name: "test3"})
	result = handler(ctx, msg3)
	require.NotNil(t, result)
	require.Equal(t, sdk.ABCICodeOK, result.Code)
	require.NotNil(t, result.Data)
	cdc.MustUnmarshalBinaryBare(result.Data, &record)
	require.NotNil(t, &record)
	require.Equal(t, "ipfs", record.Type)
	require.NotNil(t, record.Timestamp)
	require.NotNil(t, record.Path)
	t.Logf("Path: %s\n", record.Path)
}
