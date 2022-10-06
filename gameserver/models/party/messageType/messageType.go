package messageType

type MessageType byte

const (
	Expelled     MessageType = 1
	Left         MessageType = 2
	None         MessageType = 3
	Disconnected MessageType = 4
)
