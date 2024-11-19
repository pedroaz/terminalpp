import type { MessageType } from './constants';

export type ServerResponse<T> = {
	type: MessageType;
	success: boolean;
	data: T;
};

export type CmdStartData = {
	succes: boolean;
	pid: number;
};

export type UpdateExecutionData = {
	cmdLine: string;
};
