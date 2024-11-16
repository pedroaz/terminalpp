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

	switch msg.AppCmd {
	case "new-cmd":
		data := NewCmdData{}
		marshalErr := json.Unmarshal(msg.Data, &data)
		if marshalErr != nil {
			return msg.ErrorMessage(marshalErr.Error())
		}

		executionData := handler.executionManager.ExecuteNewCmd(data)
		responseData, executionErr := json.Marshal(executionData)

		if executionErr != nil {
			return msg.ErrorMessage(executionErr.Error())
		}

		return ServerResponse{
			RequestedAppCmd: msg.AppCmd,
			Success:         true,
			Data:            json.RawMessage(responseData),
		}

	default:
		return ServerResponse{
			RequestedAppCmd: msg.AppCmd,
			Success:         false,
		}
	}
}
