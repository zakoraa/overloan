package keeper

import (
	loanv1 "github.com/cosmos/cosmos-sdk/api/cosmos/loan/v1"
)

var _ loanv1.QueryServer = queryServer{}

// NewQueryServerImpl returns an implementation of the QueryServer interface
// for the provided Keeper.
func NewQueryServerImpl(k Keeper) loanv1.QueryServer {
    return queryServer{k: k}
}

type queryServer struct {
    loanv1.UnimplementedQueryServer
    k Keeper
}
