package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgDeletePost = "delete_post"

var _ sdk.Msg = &MsgDeletePost{}

func NewMsgDeletePost(creator string, id uint64) *MsgDeletePost {
	return &MsgDeletePost{
		Creator: creator,
		Id:      id,
	}
}

func (msg *MsgDeletePost) Route() string {
	return RouterKey
}

func (msg *MsgDeletePost) Type() string {
	return TypeMsgDeletePost
}

func (msg *MsgDeletePost) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeletePost) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeletePost) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
