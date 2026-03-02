package keeper

import (
	queryv1beta1 "cosmossdk.io/api/cosmos/base/query/v1beta1"
	"github.com/cosmos/cosmos-sdk/types/query"
	sdkquery "github.com/cosmos/cosmos-sdk/types/query"
)

const (
	defaultLimit uint64 = 20
	maxLimit     uint64 = 100
)

func normalizePagination(req *query.PageRequest) *query.PageRequest {
	if req == nil {
		return &query.PageRequest{
			Limit: defaultLimit,
		}
	}

	pageReq := *req

	if pageReq.Limit == 0 {
		pageReq.Limit = defaultLimit
	}

	if pageReq.Limit > maxLimit {
		pageReq.Limit = maxLimit
	}

	return &pageReq
}

func convertPageRequest(req *queryv1beta1.PageRequest) *sdkquery.PageRequest {
	if req == nil {
		return nil
	}

	return &sdkquery.PageRequest{
		Key:        req.Key,
		Offset:     req.Offset,
		Limit:      req.Limit,
		CountTotal: req.CountTotal,
		Reverse:    req.Reverse,
	}
}

func convertPageResponse(res *sdkquery.PageResponse) *queryv1beta1.PageResponse {
	if res == nil {
		return nil
	}

	return &queryv1beta1.PageResponse{
		NextKey: res.NextKey,
		Total:   res.Total,
	}
}
