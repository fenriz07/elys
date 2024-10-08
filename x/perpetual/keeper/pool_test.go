package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/elys-network/elys/testutil/keeper"
	"github.com/elys-network/elys/testutil/nullify"
	"github.com/elys-network/elys/x/perpetual/keeper"
	"github.com/elys-network/elys/x/perpetual/types"
	"github.com/stretchr/testify/require"
)

func createNPool(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Pool {
	items := make([]types.Pool, n)
	for i := range items {
		items[i] = types.NewPool((uint64)(i))

		keeper.SetPool(ctx, items[i])
	}
	return items
}

func createNPoolResponse(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.PoolResponse {
	items := make([]types.PoolResponse, n)
	for i := range items {
		items[i] = types.PoolResponse{
			AmmPoolId:                            uint64(i),
			Health:                               sdk.NewDec(100),
			Enabled:                              true,
			Closed:                               false,
			BorrowInterestRate:                   sdk.MustNewDecFromStr("0.000000000000000001"),
			PoolAssetsLong:                       []types.PoolAsset{},
			PoolAssetsShort:                      []types.PoolAsset{},
			LastHeightBorrowInterestRateComputed: 0,
			FundingRate:                          sdk.ZeroDec(),
			NetOpenInterest:                      sdk.ZeroInt(),
		}

		keeper.SetPool(ctx, types.NewPool(uint64(i)))
	}
	return items
}

func TestPoolGet(t *testing.T) {
	keeper, ctx := keepertest.PerpetualKeeper(t)
	items := createNPool(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetPool(ctx,
			item.AmmPoolId,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

func TestPoolRemove(t *testing.T) {
	keeper, ctx := keepertest.PerpetualKeeper(t)
	items := createNPool(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemovePool(ctx,
			item.AmmPoolId,
		)
		_, found := keeper.GetPool(ctx,
			item.AmmPoolId,
		)
		require.False(t, found)
	}
}

func TestPoolGetAll(t *testing.T) {
	keeper, ctx := keepertest.PerpetualKeeper(t)
	items := createNPool(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllPools(ctx)),
	)
}
