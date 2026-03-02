package keeper

import (
	"context"

	loanv1 "github.com/cosmos/cosmos-sdk/api/cosmos/loan/v1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/cosmos/cosmos-sdk/x/loan/types"
)

func (q queryServer) Loans(
	ctx context.Context,
	req *loanv1.QueryLoansRequest,
) (*loanv1.QueryLoansResponse, error) {

	if req == nil {
		return nil, types.ErrInvalidRequest
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)

	pageReq := normalizePagination(
		convertPageRequest(req.Pagination),
	)

	loans, pageRes, err := query.CollectionPaginate(
		sdkCtx,
		q.k.Loans,
		pageReq,
		func(_ uint64, value loanv1.Loan) (*loanv1.Loan, error) {
			v := value
			return &v, nil
		},
	)
	if err != nil {
		return nil, err
	}

	return &loanv1.QueryLoansResponse{
		Loans:      loans,
		Pagination: convertPageResponse(pageRes),
	}, nil
}
