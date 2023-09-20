package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/evmos/evmos/v9/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgDeployContract_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeployContract
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeployContract{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeployContract{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
