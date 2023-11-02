package keeper

import (
    "fmt"
    "encoding/binary"

    "blog/x/blog/types"

    "github.com/cosmos/cosmos-sdk/store/prefix"
    sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) AppendPost(ctx sdk.Context, post types.Post) uint64 {
    count := k.GetPostCount(ctx)
    post.Id = count
    store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PostKey))
    appendedValue := k.cdc.MustMarshal(&post)
    store.Set(GetPostIDBytes(post.Id), appendedValue)
    k.SetPostCount(ctx, count+1)
    return count
}

func (k Keeper) GetPostCount(ctx sdk.Context) uint64 {
    store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
    byteKey := types.KeyPrefix(types.PostCountKey)
    bz := store.Get(byteKey)
    if bz == nil {
        return 0
    }
    return binary.BigEndian.Uint64(bz)
}

func GetPostIDBytes(id uint64) []byte {
    bz := make([]byte, 8)
    binary.BigEndian.PutUint64(bz, id)
    return bz
}

func (k Keeper) SetPostCount(ctx sdk.Context, count uint64) {
    store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
    byteKey := types.KeyPrefix(types.PostCountKey)
    bz := make([]byte, 8)
    binary.BigEndian.PutUint64(bz, count)
    store.Set(byteKey, bz)
}

func (k Keeper) GetPost(ctx sdk.Context, id uint64) (val types.Post, found bool) {
    store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PostKey))
    b := store.Get(GetPostIDBytes(id))
    if b == nil {
        return val, false
    }
    k.cdc.MustUnmarshal(b, &val)
    return val, true
}

func (k Keeper) SetPost(ctx sdk.Context, post types.Post) {
    store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PostKey))
    updatedValue := k.cdc.MustMarshal(&post)
    store.Set(GetPostIDBytes(post.Id), updatedValue)
}

func (k Keeper) RemovePost(ctx sdk.Context, id uint64) {
    store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PostKey))
    store.Delete(GetPostIDBytes(id))
}