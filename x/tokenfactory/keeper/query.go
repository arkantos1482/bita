package keeper

import (
	"bita/x/tokenfactory/types"
)

var _ types.QueryServer = Keeper{}
