package keeper

import (
	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/x/loan/types"

	loanv1 "cosmossdk.io/api/overloan/loan/v1"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) GetTotalOutstanding(ctx sdk.Context) sdkmath.Int {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.LoanKeyPrefix)

	iterator := store.Iterator(nil, nil)
	defer iterator.Close()

	total := sdkmath.ZeroInt()

	for ; iterator.Valid(); iterator.Next() {
		var loan loanv1.Loan
		k.cdc.MustUnmarshal(iterator.Value(), &loan)

		if loan.Outstanding == nil {
			continue
		}

		amountInt, ok := sdkmath.NewIntFromString(loan.Outstanding.Amount)
		if !ok {
			panic("invalid outstanding amount in loan state")
		}

		total = total.Add(amountInt)
	}
	return total
}
