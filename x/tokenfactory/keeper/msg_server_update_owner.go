package keeper

import (
	"context"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"bita/x/tokenfactory/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UpdateOwner(goCtx context.Context, msg *types.MsgUpdateOwner) (*types.MsgUpdateOwnerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetDenom(ctx, msg.Denom)
	if !found {
		return nil, types.ErrDenomNotFound
	}

	if val.Owner != msg.Owner {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}

	var denom = types.Denom{
		Owner:              msg.NewOwner,
		Denom:              val.Denom,
		Description:        val.Description,
		MaxSupply:          val.MaxSupply,
		Supply:             val.Supply,
		Precision:          val.Precision,
		Ticker:             val.Ticker,
		Url:                val.Url,
		CanChangeMaxSupply: val.CanChangeMaxSupply,
	}

	k.SetDenom(ctx, denom)

	return &types.MsgUpdateOwnerResponse{}, nil
}
