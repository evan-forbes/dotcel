package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgClaimDeposit = "claim_deposit"

var _ sdk.Msg = &MsgClaimDeposit{}

func NewMsgClaimDeposit(signer string, hash string) *MsgClaimDeposit {
	return &MsgClaimDeposit{
		Signer: signer,
		Hash:   hash,
	}
}

func (msg *MsgClaimDeposit) Route() string {
	return RouterKey
}

func (msg *MsgClaimDeposit) Type() string {
	return TypeMsgClaimDeposit
}

func (msg *MsgClaimDeposit) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{signer}
}

func (msg *MsgClaimDeposit) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgClaimDeposit) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address (%s)", err)
	}
	return nil
}
