export type Line = {
	text: string;
};

export enum ConsoleState {
	LOADING = 'loading',
	READY = 'ready'
}

export type ConsoleInstance = {
	state: ConsoleState;
	id: string;
	lines: string[];
};
