'use strict';

function after(n, func) {
    if (!Number.isInteger(n) || n < 0) {
        throw new Error(`n must be a non-negative integer.`);
    }
    let counter = 0;
    return (...args) => {
        if (++counter >= n) {
            return func(...args);
        }
        return undefined;
    };
}

function ary(func, n) {
    return function (...args) {
        return func.apply(this, args.slice(0, n));
    };
}

function debounce(func, debounceMs, { signal, edges } = {}) {
    let pendingThis = undefined;
    let pendingArgs = null;
    const leading = edges != null && edges.includes('leading');
    const trailing = edges == null || edges.includes('trailing');
    const invoke = () => {
        if (pendingArgs !== null) {
            func.apply(pendingThis, pendingArgs);
            pendingThis = undefined;
            pendingArgs = null;
        }
    };
    const onTimerEnd = () => {
        if (trailing) {
            invoke();
        }
        cancel();
    };
    let timeoutId = null;
    const schedule = () => {
        if (timeoutId != null) {
            clearTimeout(timeoutId);
        }
        timeoutId = setTimeout(() => {
            timeoutId = null;
            onTimerEnd();
        }, debounceMs);
    };
    const cancelTimer = () => {
        if (timeoutId !== null) {
            clearTimeout(timeoutId);
            timeoutId = null;
        }
    };
    const cancel = () => {
        cancelTimer();
        pendingThis = undefined;
        pendingArgs = null;
    };
    const flush = () => {
        cancelTimer();
        invoke();
    };
    const debounced = function (...args) {
        if (signal?.aborted) {
            return;
        }
        pendingThis = this;
        pendingArgs = args;
        const isFirstCall = timeoutId == null;
        schedule();
        if (leading && isFirstCall) {
            invoke();
        }
    };
    debounced.schedule = schedule;
    debounced.cancel = cancel;
    debounced.flush = flush;
    signal?.addEventListener('abort', cancel, { once: true });
    return debounced;
}

function flow(...funcs) {
    return function (...args) {
        let result = funcs.length ? funcs[0].apply(this, args) : args[0];
        for (let i = 1; i < funcs.length; i++) {
            result = funcs[i].call(this, result);
        }
        return result;
    };
}

function flowRight(...funcs) {
    return flow(...funcs.reverse());
}

function identity(x) {
    return x;
}

function negate(func) {
    return ((...args) => !func(...args));
}

function once(func) {
    let called = false;
    let cache;
    return function (...args) {
        if (!called) {
            called = true;
            cache = func(...args);
        }
        return cache;
    };
}

function partial(func, ...partialArgs) {
    return partialImpl(func, placeholderSymbol$1, ...partialArgs);
}
function partialImpl(func, placeholder, ...partialArgs) {
    const partialed = function (...providedArgs) {
        let providedArgsIndex = 0;
        const substitutedArgs = partialArgs
            .slice()
            .map(arg => (arg === placeholder ? providedArgs[providedArgsIndex++] : arg));
        const remainingArgs = providedArgs.slice(providedArgsIndex);
        return func.apply(this, substitutedArgs.concat(remainingArgs));
    };
    if (func.prototype) {
        partialed.prototype = Object.create(func.prototype);
    }
    return partialed;
}
const placeholderSymbol$1 = Symbol('partial.placeholder');
partial.placeholder = placeholderSymbol$1;

function partialRight(func, ...partialArgs) {
    return partialRightImpl(func, placeholderSymbol, ...partialArgs);
}
function partialRightImpl(func, placeholder, ...partialArgs) {
    const partialedRight = function (...providedArgs) {
        const placeholderLength = partialArgs.filter(arg => arg === placeholder).length;
        const rangeLength = Math.max(providedArgs.length - placeholderLength, 0);
        const remainingArgs = providedArgs.slice(0, rangeLength);
        let providedArgsIndex = rangeLength;
        const substitutedArgs = partialArgs
            .slice()
            .map(arg => (arg === placeholder ? providedArgs[providedArgsIndex++] : arg));
        return func.apply(this, remainingArgs.concat(substitutedArgs));
    };
    if (func.prototype) {
        partialedRight.prototype = Object.create(func.prototype);
    }
    return partialedRight;
}
const placeholderSymbol = Symbol('partialRight.placeholder');
partialRight.placeholder = placeholderSymbol;

function rest(func, startIndex = func.length - 1) {
    return function (...args) {
        const rest = args.slice(startIndex);
        const params = args.slice(0, startIndex);
        while (params.length < startIndex) {
            params.push(undefined);
        }
        return func.apply(this, [...params, rest]);
    };
}

function unary(func) {
    return ary(func, 1);
}

exports.after = after;
exports.ary = ary;
exports.debounce = debounce;
exports.flow = flow;
exports.flowRight = flowRight;
exports.identity = identity;
exports.negate = negate;
exports.once = once;
exports.partial = partial;
exports.partialImpl = partialImpl;
exports.partialRight = partialRight;
exports.partialRightImpl = partialRightImpl;
exports.rest = rest;
exports.unary = unary;
