package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUpdatePost = "update_post"

var _ sdk.Msg = &MsgUpdatePost{}

func NewMsgUpdatePost(creator string, title string, body string, id uint64) *MsgUpdatePost {
	return &MsgUpdatePost{
		Creator: creator,
		Title:   title,
		Body:    body,
		Id:      id,
	}
}

func (msg *MsgUpdatePost) Route() string {
	return RouterKey
}

func (msg *MsgUpdatePost) Type() string {
	return TypeMsgUpdatePost
}

func (msg *MsgUpdatePost) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdatePost) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdatePost) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
