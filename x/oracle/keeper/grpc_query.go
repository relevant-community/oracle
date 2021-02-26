package keeper

import (
	"github.com/relevant-community/oracle/x/oracle/types"
)

var _ types.QueryServer = Keeper{}
