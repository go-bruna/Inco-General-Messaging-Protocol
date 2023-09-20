package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgDeployContract = "deploy_contract"

var _ sdk.Msg = &MsgDeployContract{}

func NewMsgDeployContract(creator string) *MsgDeployContract {
	return &MsgDeployContract{
		Creator: creator,
	}
}

func (msg *MsgDeployContract) Route() string {
	return RouterKey
}

func (msg *MsgDeployContract) Type() string {
	return TypeMsgDeployContract
}

func (msg *MsgDeployContract) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeployContract) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeployContract) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
