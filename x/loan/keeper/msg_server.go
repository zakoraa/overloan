package keeper

import "github.com/cosmos/cosmos-sdk/x/loan/types"

// msgServer mengimplementasikan service Msg dari proto
type msgServer struct {
	Keeper
}

// NewMsgServerImpl mengembalikan implementasi MsgServer module loan
func NewMsgServerImpl(k Keeper) types.MsgServer {
	return &msgServer{Keeper: k}
}

var _ types.MsgServer = msgServer{}