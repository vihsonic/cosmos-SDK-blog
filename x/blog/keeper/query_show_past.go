package keeper

import (
	"context"

	"fmt"

    "blog/x/blog/types"

    sdk "github.com/cosmos/cosmos-sdk/types"
    sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

func (k Keeper) ShowPast(goCtx context.Context, req *types.QueryShowPastRequest) (*types.QueryShowPastResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	fmt.Println("dfd");
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	post, found := k.GetPost(ctx, req.Id)
    if !found {
        return nil, sdkerrors.ErrKeyNotFound
    }

    return &types.QueryShowPastResponse{Post: post}, nil
	_ = ctx

	return &types.QueryShowPastResponse{}, nil
}
