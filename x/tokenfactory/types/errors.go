package types

// DONTCOVER

import (
	"cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/tokenfactory module sentinel errors
var (
	ErrSample              = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrInvalidTickerLength = errors.New(ModuleName, 1101, "invalid ticker length")
	ErrInvalidMaxSupply    = errors.New(ModuleName, 1102, "invalid max supply")
	ErrDenomNotFound       = errors.New(ModuleName, 1103, "denom not found")
)
