package server

import (
	"encoding/json"
	"fmt"
)

type CommandHandler struct {
	executionManager *ExecutionManager
}

func NewCommandHandler(exeutionManager *ExecutionManager) *CommandHandler {
	return &CommandHandler{
		executionManager: exeutionManager,
	}
}

func (handler *CommandHandler) HandleMsg(msg ClientMessage) ServerResponse {
	fmt.Printf("Raw Message Received from Client: %s\n", msg)

	switch msg.Type {
	case SEND_CMD:
		data := UpdateExecutionData{}
		json.Unmarshal(msg.Data, &data)
		executionData := handler.executionManager.ExecuteNewCmd(data)
		responseData, _ := json.Marshal(executionData)
		return ServerResponse{
			Type:    SEND_CMD_START,
			Success: true,
			Data:    json.RawMessage(responseData),
		}
	default:
		return ServerResponse{
			Type:    UNKNOWN,
			Success: false,
		}
	}
}
