package keeper

import (
	"errors"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	transfertypes "github.com/cosmos/ibc-go/v4/modules/apps/transfer/types"
	clienttypes "github.com/cosmos/ibc-go/v4/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v4/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v4/modules/core/24-host"
	"github.com/notional-labs/feeabstraction/v1/x/feeabs/types"
)

// TODO: use TWAP instead of spotprice
func (k Keeper) handleOsmosisIbcQuery(ctx sdk.Context) error {
	params := k.GetParams(ctx)
	channelID := params.OsmosisQueryChannel
	poolId := params.PoolId // for testing
	baseDenom := params.NativeIbcDenom

	return k.SendOsmosisQueryRequest(ctx, poolId, baseDenom, "uosmo", types.IBCPortID, channelID)
}

// Send request for query EstimateSwapExactAmountIn over IBC
func (k Keeper) SendOsmosisQueryRequest(ctx sdk.Context, poolId uint64, baseDenom string, quoteDenom string, sourcePort, sourceChannel string) error {
	packetData := types.NewOsmosisQueryRequestPacketData(poolId, baseDenom, quoteDenom)

	// Require channelID is the channelID profiles module is bound to
	// boundChannel := k.GetChannelId(ctx)
	// if boundChannel != sourceChannel {
	// 	return sdkerrors.Wrapf(channeltypes.ErrInvalidChannel, "invalid chanel: %s, expected %s", sourceChannel, boundChannel)
	// }

	// Get the next sequence
	sequence, found := k.channelKeeper.GetNextSequenceSend(ctx, sourcePort, sourceChannel)
	if !found {
		return sdkerrors.Wrapf(
			channeltypes.ErrSequenceSendNotFound,
			"source port: %s, source channel: %s", sourcePort, sourceChannel,
		)
	}
	sourceChannelEnd, found := k.channelKeeper.GetChannel(ctx, sourcePort, sourceChannel)
	if !found {
		return sdkerrors.Wrapf(channeltypes.ErrChannelNotFound, "port ID (%s) channel ID (%s)", sourcePort, sourceChannel)
	}

	destinationPort := sourceChannelEnd.GetCounterparty().GetPortID()
	destinationChannel := sourceChannelEnd.GetCounterparty().GetChannelID()

	timeoutHeight := clienttypes.NewHeight(0, 100000000)
	timeoutTimestamp := uint64(0)

	// Begin createOutgoingPacket logic
	channelCap, ok := k.scopedKeeper.GetCapability(ctx, host.ChannelCapabilityPath(sourcePort, sourceChannel))
	if !ok {
		return sdkerrors.Wrap(channeltypes.ErrChannelCapabilityNotFound, "module does not own channel capability")
	}

	packetBytes := packetData.GetBytes()
	// Create the IBC packet
	packet := channeltypes.NewPacket(
		packetBytes,
		sequence,
		sourcePort,
		sourceChannel,
		destinationPort,
		destinationChannel,
		timeoutHeight,
		timeoutTimestamp,
	)

	// Send the IBC packet
	return k.channelKeeper.SendPacket(ctx, channelCap, packet)
}

// OnAcknowledgementIbcOsmosisQueryRequest handle Acknowledgement for OsmosisPriceQuery packet
func (k Keeper) OnAcknowledgementIbcOsmosisQueryRequest(ctx sdk.Context, ack channeltypes.Acknowledgement) error {
	switch dispatchedAck := ack.Response.(type) {
	case *channeltypes.Acknowledgement_Error:
		_ = dispatchedAck.Error
		return nil
	case *channeltypes.Acknowledgement_Result:
		// Unmarshal dispatchedAck result
		spotPrice, err := k.UnmarshalPacketBytesToPrice(dispatchedAck.Result)
		if err != nil {
			return err
		}
		k.SetOsmosisExchangeRate(ctx, spotPrice)
		return nil
	default:
		// The counter-party module doesn't implement the correct acknowledgment format
		return errors.New("invalid acknowledgment format")
	}
}

// TODO: use TWAP instead of spotprice
func (k Keeper) handleSendIbcSwapAmountInRoute(ctx sdk.Context, memo []byte) error {
	params := k.GetParams(ctx)
	channelID := params.OsmosisSwapChannel
	k.Logger(ctx).Error("handleSendIbcSwapAmountInRoute send swap")
	return k.SendIbcSwapAmountInRoute(
		ctx,
		memo,
		types.IBCPortID,
		channelID,
	)
}

// Send request for swap SwapAmountInRoute over IBC
func (k Keeper) SendIbcSwapAmountInRoute(
	ctx sdk.Context,
	packetData []byte,
	sourcePort string,
	sourceChannel string,
) error {
	// Get source channel endpoint
	sourceChannelEnd, found := k.channelKeeper.GetChannel(ctx, sourcePort, sourceChannel)
	if !found {
		return sdkerrors.Wrapf(channeltypes.ErrChannelNotFound, "port ID (%s) channel ID (%s)", sourcePort, sourceChannel)
	}

	// Get counter-party chain endpoint infor
	destinationPort := sourceChannelEnd.GetCounterparty().GetPortID()
	destinationChannel := sourceChannelEnd.GetCounterparty().GetChannelID()

	// Get next sequence
	sequence, found := k.channelKeeper.GetNextSequenceSend(ctx, sourcePort, sourceChannel)
	if !found {
		return sdkerrors.Wrapf(
			channeltypes.ErrSequenceSendNotFound,
			"source port: %s, source channel: %s", sourcePort, sourceChannel,
		)
	}

	channelCap, ok := k.scopedKeeper.GetCapability(ctx, host.ChannelCapabilityPath(sourcePort, sourceChannel))
	if !ok {
		return sdkerrors.Wrap(channeltypes.ErrChannelCapabilityNotFound, "module does not own channel capability")
	}

	timeoutHeight := clienttypes.NewHeight(0, 100000000)
	timeoutTimestamp := uint64(0)

	// Create the IBC packet
	packet := channeltypes.NewPacket(
		packetData,
		sequence,
		sourcePort,
		sourceChannel,
		destinationPort,
		destinationChannel,
		timeoutHeight,
		timeoutTimestamp,
	)

	// Send the IBC packet
	return k.channelKeeper.SendPacket(ctx, channelCap, packet)
}

// OnAcknowledgementIbcSwapAmountInRoute handle Acknowledgement for SwapAmountInRoute packet
func (k Keeper) OnAcknowledgementIbcSwapAmountInRoute(ctx sdk.Context, ack channeltypes.Acknowledgement) error {
	switch dispatchedAck := ack.Response.(type) {
	case *channeltypes.Acknowledgement_Error:
		_ = dispatchedAck.Error
		return nil
	case *channeltypes.Acknowledgement_Result:
		// Unmarshal dispatchedAck result

		// TODO: implement logic swap success
		return nil
	default:
		// The counter-party module doesn't implement the correct acknowledgment format
		return errors.New("invalid acknowledgment format")
	}
}

// OnTimeOutIbcSwapRequest handle timeout for SwapIbc
func (k Keeper) IbcCrossChainSwapFailCallback(ctx sdk.Context) error {
	return k.transferIBCTokenToOsmosisContract(ctx)
}

// OnTimeOutIbcSwapRequest handle timeout for SwapIbc
func (k Keeper) IbcCrossChainSwapSuccessCallback(ctx sdk.Context, memo []byte) error {
	k.Logger(ctx).Error("IbcCrossChainSwapSuccessCallback")
	return k.handleSendIbcSwapAmountInRoute(ctx, memo)
}

func (k Keeper) transferIBCTokenToOsmosisContract(ctx sdk.Context) error {
	params := k.GetParams(ctx)

	moduleAccountAddress := k.GetModuleAddress()
	token := k.bk.GetBalance(ctx, moduleAccountAddress, params.OsmosisIbcDenom)

	// if token
	if sdk.NewInt(1).GTE(token.Amount) {
		return nil
	}

	memo, err := buildMemo(sdk.NewCoin("uosmo", token.Amount), params.NativeIbcDenom, params.OsmosisSwapContract, moduleAccountAddress.String())
	if err != nil {
		return err
	}

	curChainHeight := ctx.BlockHeight()

	transferMsg := transfertypes.MsgTransfer{
		SourcePort:       transfertypes.PortID,
		SourceChannel:    params.OsmosisTransferChannel,
		Token:            token,
		Sender:           moduleAccountAddress.String(),
		Receiver:         params.OsmosisSwapContract,
		TimeoutHeight:    clienttypes.NewHeight(0, uint64(curChainHeight+10)),
		TimeoutTimestamp: uint64(0),
		Memo:             memo,
	}

	_, err = k.executeTransferMsg(ctx, &transferMsg)
	if err != nil {
		return err
	}

	return nil
}

func (k Keeper) executeTransferMsg(ctx sdk.Context, transferMsg *transfertypes.MsgTransfer) (*transfertypes.MsgTransferResponse, error) {
	if err := transferMsg.ValidateBasic(); err != nil {
		return nil, fmt.Errorf("bad msg %v", err.Error())
	}
	return k.transferKeeper.Transfer(sdk.WrapSDKContext(ctx), transferMsg)

}
