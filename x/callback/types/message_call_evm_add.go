package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCallEvmAdd = "call_evm_add"

var _ sdk.Msg = &MsgCallEvmAdd{}

func NewMsgCallEvmAdd(creator string, contractAddress string, funcName string, arg string) *MsgCallEvmAdd {
	return &MsgCallEvmAdd{
		Creator:         creator,
		ContractAddress: contractAddress,
		FuncName:        funcName,
		Arg:             arg,
	}
}

func (msg *MsgCallEvmAdd) Route() string {
	return RouterKey
}

func (msg *MsgCallEvmAdd) Type() string {
	return TypeMsgCallEvmAdd
}

func (msg *MsgCallEvmAdd) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCallEvmAdd) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCallEvmAdd) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
