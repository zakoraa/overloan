package keeper

import (
	"context"

	"cosmossdk.io/collections"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/cosmos/cosmos-sdk/x/loan/types"
)

func (q queryServer) LoansByBorrower(
	ctx context.Context,
	req *types.QueryLoansByBorrowerRequest,
) (*types.QueryLoansByBorrowerResponse, error) {

	if req == nil {
		return nil, types.ErrInvalidRequest
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// validate borrower address
	borrowerAddr, err := sdk.AccAddressFromBech32(req.Borrower)
	if err != nil {
		return nil, types.ErrInvalidAddress
	}

	loans, pageRes, err := query.CollectionFilteredPaginate(
		sdkCtx,
		q.k.LoansByBorrower,
		req.Pagination,

		// filter keys
		func(key collections.Pair[sdk.AccAddress, uint64], loanID uint64) (bool, error) {
			return key.K1().Equals(borrowerAddr), nil
		},

		// transform result
		func(key collections.Pair[sdk.AccAddress, uint64], loanID uint64) (*types.Loan, error) {

			loan, err := q.k.Loans.Get(sdkCtx, loanID)
			if err != nil {
				return nil, err
			}

			return &loan, nil
		},
	)

	if err != nil {
		return nil, err
	}

	return &types.QueryLoansByBorrowerResponse{
		Loans:      loans,
		Pagination: pageRes,
	}, nil
}