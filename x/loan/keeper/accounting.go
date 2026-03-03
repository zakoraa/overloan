package keeper

import (
	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/x/loan/types"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) GetTotalOutstanding(ctx sdk.Context) sdkmath.Int {
	store := prefix.NewStore(
		runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx)),
		types.LoanKeyPrefix,
	)

	iterator := store.Iterator(nil, nil)
	defer iterator.Close()

	total := sdkmath.ZeroInt()

	for ; iterator.Valid(); iterator.Next() {
		var loan types.Loan
		k.cdc.MustUnmarshal(iterator.Value(), &loan)

		if loan.Outstanding == nil {
			continue
		}

		amountInt := loan.Outstanding.Amount

		total = total.Add(amountInt)
	}
	return total
}
