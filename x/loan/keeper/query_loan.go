package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/loan/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (q queryServer) Loan(
	ctx context.Context,
	req *types.QueryLoanRequest,
) (*types.QueryLoanResponse, error) {

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)

	loan, err := q.k.GetLoan(sdkCtx, req.LoanId)
	if err != nil {
		return nil, err
	}

	return &types.QueryLoanResponse{
		Loan: loan,
	}, nil
}
