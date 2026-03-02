package keeper

import (
	"context"

	"cosmossdk.io/collections"
	loanv1 "github.com/cosmos/cosmos-sdk/api/cosmos/loan/v1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/cosmos/cosmos-sdk/x/loan/types"
)

func (q queryServer) LoansByBorrower(
	ctx context.Context,
	req *loanv1.QueryLoansByBorrowerRequest,
) (*loanv1.QueryLoansByBorrowerResponse, error) {

	if req == nil {
		return nil, types.ErrInvalidRequest
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)

	_, err := sdk.AccAddressFromBech32(req.Borrower)
	if err != nil {
		return nil, types.ErrInvalidAddress
	}

	pageReq := normalizePagination(
		convertPageRequest(req.Pagination),
	)

	borrower := req.Borrower

	loans, pageRes, err := query.CollectionFilteredPaginate(
		sdkCtx,
		q.k.LoansByBorrower,
		pageReq,
		func(key collections.Pair[string, uint64], loanID uint64) (bool, error) {
			return key.K1() == borrower, nil
		},
		func(key collections.Pair[string, uint64], loanID uint64) (*loanv1.Loan, error) {
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

	return &loanv1.QueryLoansByBorrowerResponse{
		Loans:      loans,
		Pagination: convertPageResponse(pageRes),
	}, nil
}
