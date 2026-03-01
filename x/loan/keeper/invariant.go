package keeper

import sdk "github.com/cosmos/cosmos-sdk/types"

func (k Keeper) LoanSupplyInvariant(ctx sdk.Context) (string, bool) {

	params, err := k.GetParams(ctx)
	if err != nil {
		return "loan params error: " + err.Error(), true
	}

	moduleAddr := k.GetModuleAddress()
	moduleBalance := k.bankKeeper.GetBalance(ctx, moduleAddr, params.SettlementDenom)

	outstanding := k.GetTotalOutstanding(ctx)

	broken := !moduleBalance.Amount.Equal(outstanding)

	if broken {
		return "loan module balance invariant broken", true
	}

	return "", false
}
