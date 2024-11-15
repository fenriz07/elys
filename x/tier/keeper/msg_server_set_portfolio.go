package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/elys-network/elys/x/tier/types"
)

func (k msgServer) SetPortfolio(goCtx context.Context, msg *types.MsgSetPortfolio) (*types.MsgSetPortfolioResponse, error) {
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	user := sdk.MustAccAddressFromBech32(msg.User)
	k.RetrieveAllPortfolio(ctx, user)

	return &types.MsgSetPortfolioResponse{}, nil
}
