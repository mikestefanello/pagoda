import { debounce } from './debounce.mjs';

function throttle(func, throttleMs = 0, options = {}) {
    if (typeof options !== 'object') {
        options = {};
    }
    const { leading = true, trailing = true, signal } = options;
    return debounce(func, throttleMs, {
        leading,
        trailing,
        signal,
        maxWait: throttleMs,
    });
}

export { throttle };
