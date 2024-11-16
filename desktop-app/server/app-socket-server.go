package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type WebSocketServer struct {
	Conn *websocket.Conn
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins
	},
}

func (server *WebSocketServer) handleConnections(w http.ResponseWriter, r *http.Request, ctx *ServerCtx) {
	conn, err := upgrader.Upgrade(w, r, nil)
	server.Conn = conn
	if err != nil {
		fmt.Printf("Error upgrading connection: %v\n", err)
		return
	}
	defer conn.Close()

	fmt.Println("Client connected - starting listening loop")

	for {

		var msg ClientMessage
		err := conn.ReadJSON(&msg)

		if err != nil {
			fmt.Printf("Error reading message: %v\n", err)
		}

		res := ctx.CmdHandler.HandleMsg(msg)

		err = conn.WriteJSON(res)
		if err != nil {
			fmt.Printf("Error writing message: %v\n", err)
			break
		}
	}
}

func (server *WebSocketServer) SendTestMessages() {
	go func() {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()
		for range ticker.C {
			err := server.Conn.WriteJSON("Hello from server every second")
			if err != nil {
				fmt.Printf("Error writing message: %v\n", err)
				return
			}
		}
	}()
}

func (server *WebSocketServer) CreateAndStart(ctx *ServerCtx) {
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		server.handleConnections(w, r, ctx)
	})
	fmt.Println("WebSocket server started on :45679")
	// server.SendTestMessages()
	err := http.ListenAndServe(":45679", nil)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
