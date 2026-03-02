package keeper

import (
	"context"

	loanv1 "github.com/cosmos/cosmos-sdk/api/cosmos/loan/v1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (q queryServer) Loan(
	ctx context.Context,
	req *loanv1.QueryLoanRequest,
) (*loanv1.QueryLoanResponse, error) {

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)

	loan, err := q.k.GetLoan(sdkCtx, req.LoanId)
	if err != nil {
		return nil, err
	}

	return &loanv1.QueryLoanResponse{
		Loan: loan,
	}, nil
}
