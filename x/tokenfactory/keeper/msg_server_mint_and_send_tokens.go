package keeper

import (
	"context"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/arkantos1482/bita/x/tokenfactory/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) MintAndSendTokens(goCtx context.Context, msg *types.MsgMintAndSendTokens) (*types.MsgMintAndSendTokensResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetDenom(ctx, msg.Denom)
	if !found {
		return nil, types.ErrDenomNotFound
	}

	if val.Owner != msg.Owner {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}

	if val.Supply+msg.Amount > val.MaxSupply {
		return nil, types.ErrInvalidMaxSupply
	}
	moduleAcct := k.accountKeeper.GetModuleAddress(types.ModuleName)

	recipientAddress, err := sdk.AccAddressFromBech32(msg.Recipient)
	if err != nil {
		return nil, err
	}

	var mintCoins sdk.Coins
	mintCoins = mintCoins.Add(sdk.NewCoin(msg.Denom, sdk.NewInt(int64(msg.Amount))))

	errMint := k.bankKeeper.MintCoins(ctx, types.ModuleName, mintCoins)
	if errMint != nil {
		return nil, errMint
	}

	errSend := k.bankKeeper.SendCoins(ctx, moduleAcct, recipientAddress, mintCoins)
	if errSend != nil {
		return nil, errSend
	}

	var denom = types.Denom{
		Owner:              val.Owner,
		Supply:             val.Supply + msg.Amount,
		MaxSupply:          val.MaxSupply,
		Denom:              val.Denom,
		Description:        val.Description,
		Precision:          val.Precision,
		Ticker:             val.Ticker,
		Url:                val.Url,
		CanChangeMaxSupply: val.CanChangeMaxSupply,
	}

	k.SetDenom(ctx, denom)

	return &types.MsgMintAndSendTokensResponse{}, nil
}
