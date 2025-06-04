import { ProgressSettings } from './types';
declare const _default: {
    configure: (options: Partial<ProgressSettings>) => void;
    isStarted: () => boolean;
    done: (force?: boolean | undefined) => void;
    set: (n: number) => void;
    remove: () => void;
    start: () => void;
    status: null;
    show: () => void;
    hide: () => void;
};
export default _default;
