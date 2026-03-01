package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) SettlementSendRestriction(
	ctx sdk.Context,
	from sdk.AccAddress,
	to sdk.AccAddress,
	amount sdk.Coins,
) error {

	params, err := k.GetParams(ctx)
	if err != nil {
		return err
	}
	moduleAddr := k.GetModuleAddress()

	for _, coin := range amount {

		if coin.Denom != params.SettlementDenom {
			continue
		}

		if from.Equals(moduleAddr) || to.Equals(moduleAddr) {
			return nil
		}

		if from.String() == params.OmnibusGroupPolicy {
			return nil
		}

		return fmt.Errorf("settlement token non-transferable")
	}

	return nil
}
