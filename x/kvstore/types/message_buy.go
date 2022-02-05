package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgBuy = "buy"

var _ sdk.Msg = &MsgBuy{}

func NewMsgBuy(signer string, name string, until string, fee string) *MsgBuy {
	return &MsgBuy{
		Signer: signer,
		Name:   name,
		Until:  until,
		Fee:    fee,
	}
}

func (msg *MsgBuy) Route() string {
	return RouterKey
}

func (msg *MsgBuy) Type() string {
	return TypeMsgBuy
}

func (msg *MsgBuy) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{signer}
}

func (msg *MsgBuy) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgBuy) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address (%s)", err)
	}
	return nil
}
