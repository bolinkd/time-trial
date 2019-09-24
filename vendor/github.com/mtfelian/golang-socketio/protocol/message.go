package protocol

const (
	MessageTypeOpen        = iota // message with connection options
	MessageTypeClose              // close connection and destroy all handle routines
	MessageTypePing               // ping request message
	MessageTypePong               // pong response message
	MessageTypeEmpty              // empty message
	MessageTypeEmit               // emit request, no response
	MessageTypeAckRequest         // emit request, wait for response (ack)
	MessageTypeAckResponse        // ack response
	MessageTypeUpgrade            // upgrade message
	MessageTypeBlank              // blank message
)

// Message represents socket.io message
type Message struct {
	Type      int
	AckID     int
	EventName string
	Args      string
	Source    string
}
