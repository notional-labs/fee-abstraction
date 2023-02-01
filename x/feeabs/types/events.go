package types

const (
	EventCrossChainSwapSuccess = "cross_chain_swap_success"
	EventCrossChainSwapError   = "cross_chain_swap_error"

	EventTypePacket     = "receive_feechain_verification_packet"
	EventTypeEpochEnd   = "epoch_end"
	EventTypeEpochStart = "epoch_start"

	AttributeKeyAckSuccess  = "success"
	AttributeKeyClientID    = "client_id"
	AttributeKeyAck         = "acknowledgement"
	AttributeKeyAckError    = "ack_error"
	AttributeEpochNumber    = "epoch_number"
	AttributeEpochStartTime = "start_time"
)
