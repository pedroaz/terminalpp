package server

import "encoding/json"

type ClientMessage struct {
	AppCmd string `json:"appCmd"`
	Data   json.RawMessage
}

type ServerResponse struct {
	RequestedAppCmd string          `json:"requestedAppCmd"`
	Success         bool            `json:"success"`
	Data            json.RawMessage `json:"data"`
}

func (response ClientMessage) ErrorMessage(err string) ServerResponse {
	return ServerResponse{
		RequestedAppCmd: response.AppCmd,
		Success:         false,
		Data:            json.RawMessage("{\"error\": \"" + err + "\"}"),
	}
}

type NewCmdData struct {
	CmdLine string `json:"cmdLine"`
}

type NewCmdResponse struct {
	Success bool `json:"success"`
	PID     int  `json:"pid"`
}
