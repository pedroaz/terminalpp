import { handleCmdStart } from './console-manager.svelte';
import { MessageType } from './constants';
import type { ServerResponse } from './dtos';

export function handleMessage(event: MessageEvent): void {
	console.log('Message received from server:', event.data);
	const res = JSON.parse(event.data) as ServerResponse<any>;
	switch (res.type) {
		case MessageType.SEND_CMD_START:
			handleCmdStart(res);
			break;
		case MessageType.UPDATE_EXECUTION:
			handleCmdStart(res);
			break;
		default:
			console.log('Unknown message type:', res.type);
			break;
	}
}
