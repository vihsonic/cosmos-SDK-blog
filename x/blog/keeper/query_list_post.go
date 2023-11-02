package keeper

import (
	"context"

	"blog/x/blog/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListPost(goCtx context.Context, req *types.QueryListPostRequest) (*types.QueryListPostResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	
	var posts []types.Post
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx

	store := ctx.KVStore(k.storeKey)
    postStore := prefix.NewStore(store, types.KeyPrefix(types.PostKey))

    pageRes, err := query.Paginate(postStore, req.Pagination, func(key []byte, value []byte) error {
        var post types.Post
        if err := k.cdc.Unmarshal(value, &post); err != nil {
            return err
        }

        posts = append(posts, post)
        return nil
    })

    if err != nil {
        return nil, status.Error(codes.Internal, err.Error())
    }

    return &types.QueryListPostResponse{Post: posts, Pagination: pageRes}, nil

	return &types.QueryListPostResponse{}, nil
}
