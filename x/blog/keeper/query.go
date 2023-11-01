package keeper

import (
	"blog/x/blog/types"
)

var _ types.QueryServer = Keeper{}
