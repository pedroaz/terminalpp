package server

import (
	"bufio"
	"fmt"
	"os/exec"

	"github.com/google/uuid"
)

type ExecutionManager struct {
	socketServer *WebSocketServer
	executions   map[string]Execution
}

type Line struct {
	Text string `json:"text"`
}

type Execution struct {
	id     string
	PID    int
	stdOut []string
	stdErr []string
}

type CmdExecutionStart struct {
	Success bool `json:"success"`
	PID     int  `json:"pid"`
}

func NewExecutionManager(socketServer *WebSocketServer) *ExecutionManager {
	return &ExecutionManager{
		socketServer: socketServer,
		executions:   make(map[string]Execution),
	}
}

func (manager *ExecutionManager) ExecuteNewCmd(data SendCmdData) CmdExecutionStart {
	fmt.Printf("Executing NewCmd with data:  %s\n", data)

	execution := Execution{
		id:  uuid.NewString(),
		PID: 0,
	}
	manager.executions[execution.id] = execution

	cmd := exec.Command("sh", "-c", data.CmdLine)
	stdoutPipe, _ := cmd.StdoutPipe()

	cmdError := cmd.Start()
	if cmdError != nil {
		fmt.Println("Error starting command:", cmdError)
		return CmdExecutionStart{Success: false, PID: 0}
	}

	go func(executionID string) {
		scanner := bufio.NewScanner(stdoutPipe)
		for scanner.Scan() {
			line := scanner.Text()
			fmt.Println("Stdout:", line)
			lineObj := Line{Text: line}
			res := NewServerResponse(UPDATE_EXECUTION, true, lineObj)
			manager.socketServer.Conn.WriteJSON(res)
		}
	}(execution.id)

	return CmdExecutionStart{Success: true, PID: cmd.Process.Pid}
}
