package keeper

import (
	loanv1 "cosmossdk.io/api/overloan/loan/v1"
)

// msgServer mengimplementasikan service Msg dari proto
type msgServer struct {
	Keeper
	loanv1.UnimplementedMsgServer
}

// NewMsgServerImpl mengembalikan implementasi MsgServer module loan
func NewMsgServerImpl(k Keeper) loanv1.MsgServer {
	return &msgServer{Keeper: k}
}
