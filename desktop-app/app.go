package main

import (
	"context"
	"fmt"

	"github.com/pedroaz/terminalpp/server"
)

type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (app *App) startup(ctx context.Context) {
	app.ctx = ctx

	serverCtx := server.NewServerCtx()
	serverCtx.SocketServer.CreateAndStart(serverCtx)

	// cmdHandler := cmdhandler.CommandHandler{}
	// socketServer := WebSocketServer{}
	// socketServer.CreateAndStart(cmdHandler)
}

func (a *App) shutdown(ctx context.Context) {
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
