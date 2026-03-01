package keeper

import sdk "github.com/cosmos/cosmos-sdk/types"

func (k Keeper) LoanSupplyInvariant(ctx sdk.Context) (string, bool) {

	params := k.GetParams(ctx)

	moduleAddr := k.GetModuleAddress()
	moduleBalance := k.bankKeeper.GetBalance(ctx, moduleAddr, params.SettlementDenom)

	outstanding := k.GetTotalOutstanding(ctx)

	broken := !moduleBalance.Amount.Equal(outstanding)

	return "loan module balance invariant broken", broken
}
