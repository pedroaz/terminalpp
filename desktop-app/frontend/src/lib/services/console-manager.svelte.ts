import type { ServerResponse, CmdStartData } from './dtos';
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

export function handleCmdStart(message: ServerResponse<CmdStartData>): void {
	console.log('New console received:', message);
}

export function handleUpdateExecution(message: ServerResponse<CmdStartData>): void {
	console.log('New console received:', message);
}
