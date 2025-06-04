import { Poll } from './poll';
import { PollOptions } from './types';
declare class Polls {
    protected polls: Poll[];
    constructor();
    add(interval: number, cb: VoidFunction, options: PollOptions): {
        stop: VoidFunction;
        start: VoidFunction;
    };
    clear(): void;
    protected setupVisibilityListener(): void;
}
export declare const polls: Polls;
export {};
