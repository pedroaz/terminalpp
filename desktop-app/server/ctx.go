package server

type ServerCtx struct {
	ExecutionManager *ExecutionManager
	SocketServer     *WebSocketServer
	CmdHandler       *CommandHandler
}

func NewServerCtx() *ServerCtx {

	socketServer := &WebSocketServer{}
	executionManager := NewExecutionManager(socketServer)
	cmdHandler := NewCommandHandler(executionManager)

	return &ServerCtx{
		ExecutionManager: executionManager,
		SocketServer:     socketServer,
		CmdHandler:       cmdHandler,
	}
}
