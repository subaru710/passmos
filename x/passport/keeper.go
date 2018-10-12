package passport

import (
	"bytes"
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strconv"
	"time"
)

type Keeper struct {
	key     sdk.StoreKey
	cdc     *codec.Codec
	exstore ExternalStore
}

func NewKeeper(key sdk.StoreKey, store ExternalStore) Keeper {
	cdc := codec.New()
	codec.RegisterCrypto(cdc)
	return Keeper{
		key:     key,
		cdc:     cdc,
		exstore: store,
	}
}

func (k Keeper) CreatePassport(ctx sdk.Context, addr sdk.AccAddress, pd PersonalData) (*Record, sdk.Error) {
	return createPassport(k, ctx, addr, pd)
}

func (k Keeper) UpdatePassport(ctx sdk.Context, addr sdk.AccAddress, pd PersonalData) (*Record, sdk.Error) {
	return updatePassport(k, ctx, addr, pd)
}

func (k Keeper) AuthorizePassport(ctx sdk.Context, addr sdk.AccAddress, receiver sdk.AccAddress) sdk.Error {
	return authorizePassport(k, ctx, addr, receiver)
}

//______________________________________________________________________________________________

func setRecord(store sdk.KVStore, addr sdk.AccAddress, i int64, record []byte) {
	key := append(append(addr, ":"...), strconv.FormatInt(i, 10)...)
	store.Set(key, record)
}

func addAuthorization(store sdk.KVStore, addr sdk.AccAddress, receiver sdk.AccAddress) {
	var buffer bytes.Buffer
	buffer.Write(addr)
	buffer.WriteString(":")
	buffer.Write(receiver)
	key := buffer.Bytes()
	if !store.Has(key) {
		store.Set(key, receiver)
	}
}

func createPassport(k Keeper, ctx sdk.Context, addr sdk.AccAddress, data PersonalData) (*Record, sdk.Error) {
	store := ctx.KVStore(k.key)
	if store.Has(addr) {
		return nil, sdk.ErrInvalidAddress("this address already has a passport")
	}
	path, err := k.exstore.SetPersonalData(data)
	if err != nil {
		//TODO: error code
		fmt.Println(err)
		return nil, sdk.ErrInternal("fail to store to external store")
	}
	record := &Record{
		Type:      k.exstore.Type(),
		Path:      path,
		Timestamp: time.Now().UTC(),
	}
	bz, err := k.cdc.MarshalBinaryBare(record)
	if err != nil {
		panic(err)
	}
	setRecord(store, addr, 0, bz)
	// set the counter
	bz = k.cdc.MustMarshalBinary(1)
	store.Set(addr, bz)
	return record, nil
}

func updatePassport(k Keeper, ctx sdk.Context, addr sdk.AccAddress, data PersonalData) (*Record, sdk.Error) {
	store := ctx.KVStore(k.key)
	if !store.Has(addr) {
		return nil, sdk.ErrInvalidAddress("this address doesn't have a passport")
	}
	path, err := k.exstore.SetPersonalData(data)
	if err != nil {
		//TODO: error code
		return nil, sdk.ErrInternal("fail to store to external store")
	}
	record := &Record{
		Type:      k.exstore.Type(),
		Path:      path,
		Timestamp: time.Now().UTC(),
	}
	bz, err := k.cdc.MarshalBinaryBare(record)
	if err != nil {
		panic(err)
	}
	// increase the counter
	var cnt int64
	bz = store.Get(addr)
	if bz == nil {
		cnt = 0
	} else {
		err := k.cdc.UnmarshalBinary(bz, &cnt)
		if err != nil {
			//TODO: error code
			fmt.Println(err)
			return nil, sdk.ErrInternal("invalid counter")
		}
	}
	setRecord(store, addr, cnt+1, bz)
	bz = k.cdc.MustMarshalBinary(cnt + 1)
	store.Set(addr, bz)
	return record, nil
}

func authorizePassport(k Keeper, ctx sdk.Context, addr sdk.AccAddress, receiver sdk.AccAddress) sdk.Error {
	store := ctx.KVStore(k.key)
	addAuthorization(store, addr, receiver)
	return nil
}
