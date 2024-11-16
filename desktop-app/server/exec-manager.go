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

type Execution struct {
	id     string
	PID    int
	stdOut []string
	stdErr []string
}

func NewExecutionManager(socketServer *WebSocketServer) *ExecutionManager {
	return &ExecutionManager{
		socketServer: socketServer,
		executions:   make(map[string]Execution),
	}
}

func (manager *ExecutionManager) ExecuteNewCmd(data NewCmdData) NewCmdResponse {
	fmt.Printf("Executing NewCmd with data:  %s\n", data)

	execution := Execution{
		id:  uuid.NewString(),
		PID: 0,
	}
	manager.executions[execution.id] = execution

	cmd := exec.Command("sh", "-c", data.CmdLine)
	cmdError := cmd.Start()

	stdout, cmdError := cmd.StdoutPipe()

	go func() {
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			line := scanner.Text()
			// Process the line as needed. For example, log it or store it.
			fmt.Println("Stdout:", line)
			// manager.socketServer.Conn.WriteJSON(line)
		}
		if scanner.Err() != nil {
			fmt.Println("Error reading stdout:", scanner.Err())
		}
	}()

	if cmdError != nil {
		return NewCmdResponse{Success: false, PID: 0}
	}
	return NewCmdResponse{Success: true, PID: cmd.Process.Pid}
}
