package server

import "encoding/json"

// -----------------//
// Client to Server //
// -----------------//
type ClientMessage struct {
	Type MessageType `json:"type"`
	Data json.RawMessage
}

type SendCmdData struct {
	CmdLine string `json:"cmdLine"`
}

// -----------------//
// Server to Client //
// -----------------//

type ServerResponse struct {
	Type    MessageType     `json:"type"`
	Success bool            `json:"success"`
	Data    json.RawMessage `json:"data"`
}

func NewServerResponse(t MessageType, success bool, data any) ServerResponse {
	bytes, _ := json.Marshal(data)
	return ServerResponse{
		Type:    t,
		Success: success,
		Data:    json.RawMessage(bytes),
	}
}
