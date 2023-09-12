package keeper

import (
	"github.com/evmos/evmos/v9/x/callback/types"
)

var _ types.QueryServer = Keeper{}
