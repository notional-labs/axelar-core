package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/axelarnetwork/axelar-core/x/multisig/exported"
)

var _ sdk.Msg = &StartKeygenRequest{}

// NewStartKeygenRequest constructor for StartKeygenRequest
func NewStartKeygenRequest(sender sdk.AccAddress, keyID exported.KeyID) *StartKeygenRequest {
	return &StartKeygenRequest{
		Sender: sender,
		KeyID:  keyID,
	}
}

// Route implements sdk.Msg
func (m StartKeygenRequest) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (m StartKeygenRequest) Type() string {
	return "StartKeygen"
}

// ValidateBasic implements the sdk.Msg interface.
func (m StartKeygenRequest) ValidateBasic() error {
	if err := sdk.VerifyAddressFormat(m.Sender); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, sdkerrors.Wrap(err, "sender").Error())
	}

	if err := m.KeyID.ValidateBasic(); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	return nil
}

// GetSigners implements the sdk.Msg interface
func (m StartKeygenRequest) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Sender}
}

// GetSignBytes implements sdk.Msg
func (m StartKeygenRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&m)
	return sdk.MustSortJSON(bz)
}
