import { MessageType } from './constants';

export function sendCmdMessage(consoleId: string, cmdLine: string): string {
	return JSON.stringify({
		type: MessageType.SEND_CMD,
		data: {
			consoleId: consoleId,
			cmdLine: cmdLine
		}
	});
}
