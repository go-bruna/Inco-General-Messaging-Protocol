package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/evmos/evmos/v9/x/callback/types"
)

// GetTestDataCount get the total number of testData
func (k Keeper) GetTestDataCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.TestDataCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetTestDataCount set the total number of testData
func (k Keeper) SetTestDataCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.TestDataCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendTestData appends a testData in the store with a new id and update the count
func (k Keeper) AppendTestData(
	ctx sdk.Context,
	testData types.TestData,
) uint64 {
	// Create the testData
	count := k.GetTestDataCount(ctx)

	// Set the ID of the appended value
	testData.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TestDataKey))
	appendedValue := k.cdc.MustMarshal(&testData)
	store.Set(GetTestDataIDBytes(testData.Id), appendedValue)

	// Update testData count
	k.SetTestDataCount(ctx, count+1)

	return count
}

// SetTestData set a specific testData in the store
func (k Keeper) SetTestData(ctx sdk.Context, testData types.TestData) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TestDataKey))
	b := k.cdc.MustMarshal(&testData)
	store.Set(GetTestDataIDBytes(testData.Id), b)
}

// GetTestData returns a testData from its id
func (k Keeper) GetTestData(ctx sdk.Context, id uint64) (val types.TestData, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TestDataKey))
	b := store.Get(GetTestDataIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveTestData removes a testData from the store
func (k Keeper) RemoveTestData(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TestDataKey))
	store.Delete(GetTestDataIDBytes(id))
}

// GetAllTestData returns all testData
func (k Keeper) GetAllTestData(ctx sdk.Context) (list []types.TestData) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TestDataKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.TestData
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetTestDataIDBytes returns the byte representation of the ID
func GetTestDataIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetTestDataIDFromBytes returns ID in uint64 format from a byte array
func GetTestDataIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
