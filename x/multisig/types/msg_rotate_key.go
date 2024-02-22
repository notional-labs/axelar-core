package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/axelarnetwork/axelar-core/x/multisig/exported"
	nexus "github.com/axelarnetwork/axelar-core/x/nexus/exported"
)

var _ sdk.Msg = &RotateKeyRequest{}

// NewRotateKeyRequest constructor for RotateKeyRequest
func NewRotateKeyRequest(sender sdk.AccAddress, chain nexus.ChainName, keyID exported.KeyID) *RotateKeyRequest {
	return &RotateKeyRequest{
		Sender: sender,
		Chain:  chain,
		KeyID:  keyID,
	}
}

// ValidateBasic implements the sdk.Msg interface.
func (m RotateKeyRequest) ValidateBasic() error {
	if err := sdk.VerifyAddressFormat(m.Sender); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, sdkerrors.Wrap(err, "sender").Error())
	}

	if err := m.Chain.Validate(); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	if err := m.KeyID.ValidateBasic(); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	return nil
}

// Route implements sdk.Msg
func (m RotateKeyRequest) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (m RotateKeyRequest) Type() string {
	return "RotateKey"
}

// GetSigners implements the sdk.Msg interface
func (m RotateKeyRequest) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Sender}
}

// GetSignBytes implements sdk.Msg
func (m RotateKeyRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&m)
	return sdk.MustSortJSON(bz)
}
