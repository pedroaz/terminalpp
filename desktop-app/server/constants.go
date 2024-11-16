package server

type MessageType string

const (
	UNKNOWN          MessageType = "unknown"
	ERROR            MessageType = "error"
	SEND_CMD         MessageType = "send-cmd"
	SEND_CMD_START   MessageType = "send-cmd-start"
	UPDATE_EXECUTION MessageType = "update-execution"
)
