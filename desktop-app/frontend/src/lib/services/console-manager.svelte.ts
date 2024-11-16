import type { ServerResponseCmdStart } from './dtos';
import { ConsoleState, type ConsoleInstance } from './models';
import { v4 as uuidv4 } from 'uuid';

const consoles = $state(new Map<string, ConsoleInstance>());

export function createAndConnectNewConsole(): string {
	const newConsole = {
		id: uuidv4(),
		lines: [],
		state: ConsoleState.READY
	};
	consoles.set(newConsole.id, newConsole);
	return newConsole.id;
}

export function getConsole(id: string): ConsoleInstance | undefined {
	const c = consoles.get(id);
	return c;
}

export function handleCmdStart(message: ServerResponseCmdStart): void {
	console.log('New console received:', message);
	// const newConsole = {
	// 	id: message.data.id,
	// 	lines: []
	// };
	// consoles.push(newConsole);
}
