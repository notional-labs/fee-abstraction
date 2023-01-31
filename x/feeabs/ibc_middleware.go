package feeabs

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	channeltypes "github.com/cosmos/ibc-go/v4/modules/core/04-channel/types"
	porttypes "github.com/cosmos/ibc-go/v4/modules/core/05-port/types"
	host "github.com/cosmos/ibc-go/v4/modules/core/24-host"
	ibcexported "github.com/cosmos/ibc-go/v4/modules/core/exported"
	"github.com/notional-labs/feeabstraction/v1/x/feeabs/keeper"
	"github.com/notional-labs/feeabstraction/v1/x/feeabs/types"
)

var _ porttypes.Middleware = &IBCMiddleware{}

// IBCMiddleware implements the ICS26 callbacks for the transfer middleware given
// the feeabs keeper and the underlying application.
type IBCMiddleware struct {
	IBCModule porttypes.IBCModule
	keeper    keeper.Keeper
}

// NewIBCMiddleware creates a new IBCMiddleware given the keeper and underlying application
func NewIBCMiddleware(k keeper.Keeper, appModule porttypes.IBCModule) IBCMiddleware {
	return IBCMiddleware{
		IBCModule: appModule,
		keeper:    k,
	}
}

// -------------------------------------------------------------------------------------------------------------------
// OnChanOpenInit implements the IBCModule interface.
func (im IBCMiddleware) OnChanOpenInit(
	ctx sdk.Context,
	order channeltypes.Order,
	connectionHops []string,
	portID string,
	channelID string,
	channelCap *capabilitytypes.Capability,
	counterparty channeltypes.Counterparty,
	version string,
) (string, error) {
	if err := ValidateChannelParams(ctx, im.keeper, order, portID, channelID); err != nil {
		return "", err
	}

	// Claim channel capability passed back by IBC module
	if err := im.keeper.ClaimCapability(ctx, channelCap, host.ChannelCapabilityPath(portID, channelID)); err != nil {
		return "", err
	}

	return version, nil
}

func ValidateChannelParams(
	ctx sdk.Context,
	keeper keeper.Keeper,
	order channeltypes.Order,
	portID string,
	channelID string,
) error {
	if order != channeltypes.UNORDERED {
		return sdkerrors.Wrapf(channeltypes.ErrInvalidChannelOrdering, "expected %s channel, got %s ", channeltypes.UNORDERED, order)
	}

	// Require portID is the portID profiles module is bound to
	boundPort := keeper.GetPort(ctx)
	if boundPort != portID {
		return sdkerrors.Wrapf(porttypes.ErrInvalidPort, "invalid port: %s, expected %s", portID, boundPort)
	}

	return nil
}

// OnChanOpenTry implements the IBCModule interface.
func (im IBCMiddleware) OnChanOpenTry(
	ctx sdk.Context,
	order channeltypes.Order,
	connectionHops []string,
	portID,
	channelID string,
	channelCap *capabilitytypes.Capability,
	counterparty channeltypes.Counterparty,
	counterpartyVersion string,
) (string, error) {
	if err := ValidateChannelParams(ctx, im.keeper, order, portID, channelID); err != nil {
		return "", err
	}

	// Module may have already claimed capability in OnChanOpenInit in the case of crossing hellos
	// (ie chainA and chainB both call ChanOpenInit before one of them calls ChanOpenTry)
	// If module can already authenticate the capability then module already owns it so we don't need to claim
	// Otherwise, module does not have channel capability and we must claim it from IBC
	if !im.keeper.AuthenticateCapability(ctx, channelCap, host.ChannelCapabilityPath(portID, channelID)) {
		// Only claim channel capability passed back by IBC module if we do not already own it
		err := im.keeper.ClaimCapability(ctx, channelCap, host.ChannelCapabilityPath(portID, channelID))
		if err != nil {
			return "", err
		}
	}

	return counterpartyVersion, nil
}

// OnChanOpenAck implements the IBCModule interface.
func (im IBCMiddleware) OnChanOpenAck(
	ctx sdk.Context,
	portID,
	channelID string,
	counterpartyChannelID string,
	counterpartyVersion string,
) error {
	return nil
}

// OnChanOpenConfirm implements the IBCModule interface.
func (im IBCMiddleware) OnChanOpenConfirm(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	return nil
}

// OnChanCloseInit implements the IBCModule interface.
func (im IBCMiddleware) OnChanCloseInit(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	// Disallow user-initiated channel closing for channels
	return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "user cannot close channel")
}

// OnChanCloseConfirm implements the IBCModule interface.
func (im IBCMiddleware) OnChanCloseConfirm(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	// no need to implement
	return nil
}

// OnRecvPacket implements the IBCModule interface.
func (im IBCMiddleware) OnRecvPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	relayer sdk.AccAddress,
) ibcexported.Acknowledgement {
	// no need to implement
	ack := im.IBCModule.OnRecvPacket(ctx, packet, relayer)
	// NOTE: acknowledgement will be written synchronously during IBC handler execution.
	return ack
}

// OnAcknowledgementPacket implements the IBCModule interface.
func (im IBCMiddleware) OnAcknowledgementPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	acknowledgement []byte,
	relayer sdk.AccAddress,
) error {
	var ack channeltypes.Acknowledgement
	if err := types.ModuleCdc.UnmarshalJSON(acknowledgement, &ack); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "cannot unmarshal packet acknowledgement: %v", err)
	}
	// TODO :  Handler ack logic here
	// TODO : update spot price when receive ack from osmosis chain

	return nil
}

// OnTimeoutPacket implements the IBCModule interface.
func (im IBCMiddleware) OnTimeoutPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	relayer sdk.AccAddress,
) error {
	// TODO: Resend request if timeout
	// TODO: emit event
	return nil
}

//---------------------ICS4Wrapper-----------------------------------------
// SendPacket implements the ICS4Wrapper interface from the transfer module
func (im IBCMiddleware) SendPacket(
	ctx sdk.Context,
	chanCap *capabilitytypes.Capability,
	packet ibcexported.PacketI,
) error {
	return nil
}

func (im IBCMiddleware) WriteAcknowledgement(
	ctx sdk.Context,
	chanCap *capabilitytypes.Capability,
	packet ibcexported.PacketI,
	ack ibcexported.Acknowledgement,
) error {
	return nil
}

// GetAppVersion returns the underlying application version.
func (im IBCMiddleware) GetAppVersion(ctx sdk.Context, portID, channelID string) (string, bool) {
	return im.keeper.GetAppVersion(ctx, portID, channelID)
}
