// Singleton Instance
let websocket: WebSocket | null = null;
let connected = $state(false);

function connectWebSocket(): void {
	if (
		websocket &&
		(websocket.readyState === WebSocket.OPEN || websocket.readyState === WebSocket.CONNECTING)
	) {
		console.warn('WebSocket is already open or connecting.');
		return;
	}

	websocket = new WebSocket('http://localhost:45679/ws');

	websocket.onopen = () => {
		console.log('WebSocket connection opened.');
		connected = true;
	};

	websocket.onmessage = (event: MessageEvent) => {
		console.log('Message received from server:', event.data);
	};

	websocket.onerror = (error: Event) => {
		console.error('WebSocket error:', error);
		connected = false;
	};

	websocket.onclose = (event: CloseEvent) => {
		console.log('WebSocket connection closed:', event.reason);
		connected = false;
	};
}

function sendMessage(message: any): void {
	if (websocket && websocket.readyState === WebSocket.OPEN) {
		websocket.send(message);
	} else {
		console.warn('WebSocket is not open.');
	}
}

function closeConnection(): void {
	if (websocket) {
		websocket.close();
		websocket = null;
	}
}

function addEventListener(event: string, callback: (data: any) => void): void {
	if (websocket) {
		websocket.addEventListener(event, callback);
	} else {
		console.warn('WebSocket is not open.');
	}
}

function isWebsocketConnected(): boolean {
	return connected;
}

export { connectWebSocket, sendMessage, closeConnection, addEventListener, isWebsocketConnected };
