package keeper_test

import (
	"context"
	"testing"

	keepertest "androma/testutil/keeper"
	"androma/x/androma/keeper"
	"androma/x/androma/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.AndromaKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
