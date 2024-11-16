import type { MessageType } from './constants';

export type ServerResponse = {
	type: MessageType;
	success: boolean;
	data: any;
};

export type ServerResponseCmdStart = {
	type: string;
	success: boolean;
	data: ServerResponseCmdStartData;
};

type ServerResponseCmdStartData = {
	succes: boolean;
	pid: number;
};
