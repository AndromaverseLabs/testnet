package androma_test

import (
	"testing"

	keepertest "androma/testutil/keeper"
	"androma/testutil/nullify"
	"androma/x/androma"
	"androma/x/androma/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.AndromaKeeper(t)
	androma.InitGenesis(ctx, *k, genesisState)
	got := androma.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
