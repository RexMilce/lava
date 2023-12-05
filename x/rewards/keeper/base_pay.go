package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lavanet/lava/x/rewards/types"
)

// SetBasePay set a specific BasePay in the store from its index
func (k Keeper) SetBasePay(ctx sdk.Context, index types.BasePayIndex, basePay types.BasePay) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BasePayPrefix))
	b := k.cdc.MustMarshal(&basePay)
	store.Set([]byte(index.String()), b)
}

// GetBasePay returns a BasePay from its index
func (k Keeper) GetBasePay(
	ctx sdk.Context,
	index types.BasePayIndex,
) (val types.BasePay, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BasePayPrefix))

	b := store.Get([]byte(index.String()))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveBasePay removes a BasePay from the store
func (k Keeper) RemoveBasePay(
	ctx sdk.Context,
	index types.BasePayIndex,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BasePayPrefix))
	store.Delete([]byte(index.String()))
}

// GetAllBasePay returns all BasePay
func (k Keeper) GetAllBasePay(ctx sdk.Context) (list []types.BasePayWithIndex) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BasePayPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.BasePay
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, types.BasePayWithIndex{BasePayIndex: types.BasePayKeyRecover(string(iterator.Key())), BasePay: val})
	}

	return
}
