package keeper

import (
	"context"

	"blog/x/blog/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UpdatePost(goCtx context.Context, msg *types.MsgUpdatePost) (*types.MsgUpdatePostResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx
	var post = types.Post{
        Creator: msg.Creator,
        Title:   msg.Title,
        Body:    msg.Body,
		Id: msg.Id,
    }
    k.SetPost(ctx, post);
    return &types.MsgUpdatePostResponse{}, nil
}
