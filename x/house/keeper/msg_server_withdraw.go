package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sge-network/sge/x/house/types"
)

// Withdraw performs withdrawal of unused tokens corresponding to a deposit.
func (k msgServer) Withdraw(goCtx context.Context,
	msg *types.MsgWithdraw,
) (*types.MsgWithdrawResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var payload types.WithdrawTicketPayload
	if err := k.ovmKeeper.VerifyTicketUnmarshal(goCtx, msg.Ticket, &payload); err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInTicketVerification, "%s", err)
	}

	if err := payload.Validate(msg.Creator); err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInTicketPayloadValidation, "%s", err)
	}

	id, err := k.Keeper.Withdraw(ctx, msg.Creator, msg.MarketUID,
		msg.ParticipationIndex, msg.Mode, msg.Amount)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "process withdrawal")
	}

	return &types.MsgWithdrawResponse{
		ID:                 id,
		MarketUID:          msg.MarketUID,
		ParticipationIndex: msg.ParticipationIndex,
	}, nil
}
