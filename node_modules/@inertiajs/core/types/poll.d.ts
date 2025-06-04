import { PollOptions } from './types';
export declare class Poll {
    protected id: number | null;
    protected throttle: boolean;
    protected keepAlive: boolean;
    protected cb: VoidFunction;
    protected interval: number;
    protected cbCount: number;
    constructor(interval: number, cb: VoidFunction, options: PollOptions);
    stop(): void;
    start(): void;
    isInBackground(hidden: boolean): void;
}
