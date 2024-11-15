package types

import (
	ptypes "github.com/elys-network/elys/x/parameter/types"
	"testing"

	"cosmossdk.io/math"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/elys-network/elys/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgUnstake_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUnstake
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUnstake{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUnstake{
				Creator: sample.AccAddress(),
				Amount:  math.NewInt(200),
				Asset:   ptypes.ATOM,
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
