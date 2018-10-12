package passport

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	dbm "github.com/tendermint/tendermint/libs/db"
	"github.com/tendermint/tendermint/libs/log"
)

func setupMultiStore() (sdk.MultiStore, *sdk.KVStoreKey) {
	db := dbm.NewMemDB()
	passKey := sdk.NewKVStoreKey("passkey")
	ms := store.NewCommitMultiStore(db)
	ms.MountStoreWithDB(passKey, sdk.StoreTypeIAVL, db)
	ms.LoadLatestVersion()
	return ms, passKey
}

func TestKeeper(t *testing.T) {
	ms, passKey := setupMultiStore()
	passKeeper := NewKeeper(passKey, NewIpfsStore("https://ipfs.infura.io:5001"))
	ctx := sdk.NewContext(ms, abci.Header{}, false, log.NewNopLogger())
	addr := sdk.AccAddress([]byte("some-address"))

	record, err := passKeeper.CreatePassport(ctx, addr, PersonalData{Name: "test"})
	require.Nil(t, err)
	require.NotNil(t, record)
	require.Equal(t, "ipfs", record.Type)
	require.NotNil(t, record.Timestamp)
	require.NotNil(t, record.Path)
	t.Logf("Path: %s\n", record.Path)

	_, err = passKeeper.CreatePassport(ctx, addr, PersonalData{Name: "test2"})
	require.NotNil(t, err)
	record2, err := passKeeper.UpdatePassport(ctx, addr, PersonalData{Name: "test3"})
	require.Nil(t, err)
	require.NotNil(t, record2)
	require.Equal(t, "ipfs", record2.Type)
	require.NotNil(t, record2.Timestamp)
	require.NotNil(t, record2.Path)
	t.Logf("Path: %s\n", record2.Path)
	require.NotEqual(t, record.Path, record2.Path)
}

func TestKeeperAuthorize(t *testing.T) {
	ms, passKey := setupMultiStore()
	passKeeper := NewKeeper(passKey, NewIpfsStore("https://ipfs.infura.io:5001"))
	ctx := sdk.NewContext(ms, abci.Header{}, false, log.NewNopLogger())
	addr := sdk.AccAddress([]byte("some-address"))
	receiver := sdk.AccAddress([]byte("receiver-address"))
	err := passKeeper.AuthorizePassport(ctx, addr, receiver)
	require.Nil(t, err)
}
