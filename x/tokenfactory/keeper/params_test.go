package keeper_test

import (
	"testing"

	testkeeper "github.com/arkantos1482/bita/testutil/keeper"
	"github.com/arkantos1482/bita/x/tokenfactory/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.TokenfactoryKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
