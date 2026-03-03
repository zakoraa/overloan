package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/cosmos/cosmos-sdk/x/loan/types"
)

func (q queryServer) Loans(
	ctx context.Context,
	req *types.QueryLoansRequest,
) (*types.QueryLoansResponse, error) {

	if req == nil {
		return nil, types.ErrInvalidRequest
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)

	loans, pageRes, err := query.CollectionPaginate(
		sdkCtx,
		q.k.Loans,
		req.Pagination,
		func(_ uint64, value types.Loan) (*types.Loan, error) {
			v := value
			return &v, nil
		},
	)
	if err != nil {
		return nil, err
	}

	return &types.QueryLoansResponse{
		Loans:      loans,
		Pagination: pageRes,
	}, nil
}
