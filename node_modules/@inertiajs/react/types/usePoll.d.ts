import { PollOptions, ReloadOptions } from '@inertiajs/core';
export default function usePoll(interval: number, requestOptions?: ReloadOptions, options?: PollOptions): {
    stop: VoidFunction;
    start: VoidFunction;
};
