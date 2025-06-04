'use strict';

Object.defineProperty(exports, Symbol.toStringTag, { value: 'Module' });

const unary = require('../_chunk/unary-BVQ0iC.js');
const noop = require('../_chunk/noop-2IwLUk.js');
const delay = require('../_chunk/delay-_VMfFa.js');

async function asyncNoop() { }

function before(n, func) {
    if (!Number.isInteger(n) || n < 0) {
        throw new Error('n must be a non-negative integer.');
    }
    let counter = 0;
    return (...args) => {
        if (++counter < n) {
            return func(...args);
        }
        return undefined;
    };
}

function curry(func) {
    if (func.length === 0 || func.length === 1) {
        return func;
    }
    return function (arg) {
        return makeCurry(func, func.length, [arg]);
    };
}
function makeCurry(origin, argsLength, args) {
    if (args.length === argsLength) {
        return origin(...args);
    }
    else {
        const next = function (arg) {
            return makeCurry(origin, argsLength, [...args, arg]);
        };
        return next;
    }
}

function curryRight(func) {
    if (func.length === 0 || func.length === 1) {
        return func;
    }
    return function (arg) {
        return makeCurryRight(func, func.length, [arg]);
    };
}
function makeCurryRight(origin, argsLength, args) {
    if (args.length === argsLength) {
        return origin(...args);
    }
    else {
        const next = function (arg) {
            return makeCurryRight(origin, argsLength, [arg, ...args]);
        };
        return next;
    }
}

function memoize(fn, options = {}) {
    const { cache = new Map(), getCacheKey } = options;
    const memoizedFn = function (arg) {
        const key = getCacheKey ? getCacheKey(arg) : arg;
        if (cache.has(key)) {
            return cache.get(key);
        }
        const result = fn.call(this, arg);
        cache.set(key, result);
        return result;
    };
    memoizedFn.cache = cache;
    return memoizedFn;
}

const DEFAULT_DELAY = 0;
const DEFAULT_RETRIES = Number.POSITIVE_INFINITY;
async function retry(func, _options) {
    let delay$1;
    let retries;
    let signal;
    if (typeof _options === 'number') {
        delay$1 = DEFAULT_DELAY;
        retries = _options;
        signal = undefined;
    }
    else {
        delay$1 = _options?.delay ?? DEFAULT_DELAY;
        retries = _options?.retries ?? DEFAULT_RETRIES;
        signal = _options?.signal;
    }
    let error;
    for (let attempts = 0; attempts < retries; attempts++) {
        if (signal?.aborted) {
            throw error ?? new Error(`The retry operation was aborted due to an abort signal.`);
        }
        try {
            return await func();
        }
        catch (err) {
            error = err;
            const currentDelay = typeof delay$1 === 'function' ? delay$1(attempts) : delay$1;
            await delay.delay(currentDelay);
        }
    }
    throw error;
}

function spread(func) {
    return function (argsArr) {
        return func.apply(this, argsArr);
    };
}

function throttle(func, throttleMs, { signal, edges = ['leading', 'trailing'] } = {}) {
    let pendingAt = null;
    const debounced = unary.debounce(func, throttleMs, { signal, edges });
    const throttled = function (...args) {
        if (pendingAt == null) {
            pendingAt = Date.now();
        }
        else {
            if (Date.now() - pendingAt >= throttleMs) {
                pendingAt = Date.now();
                debounced.cancel();
            }
        }
        debounced(...args);
    };
    throttled.cancel = debounced.cancel;
    throttled.flush = debounced.flush;
    return throttled;
}

exports.after = unary.after;
exports.ary = unary.ary;
exports.debounce = unary.debounce;
exports.flow = unary.flow;
exports.flowRight = unary.flowRight;
exports.identity = unary.identity;
exports.negate = unary.negate;
exports.once = unary.once;
exports.partial = unary.partial;
exports.partialRight = unary.partialRight;
exports.rest = unary.rest;
exports.unary = unary.unary;
exports.noop = noop.noop;
exports.asyncNoop = asyncNoop;
exports.before = before;
exports.curry = curry;
exports.curryRight = curryRight;
exports.memoize = memoize;
exports.retry = retry;
exports.spread = spread;
exports.throttle = throttle;
