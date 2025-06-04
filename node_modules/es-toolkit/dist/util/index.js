'use strict';

Object.defineProperty(exports, Symbol.toStringTag, { value: 'Module' });

function attempt(func) {
    try {
        return [null, func()];
    }
    catch (error) {
        return [error, null];
    }
}

async function attemptAsync(func) {
    try {
        const result = await func();
        return [null, result];
    }
    catch (error) {
        return [error, null];
    }
}

function invariant(condition, message) {
    if (condition) {
        return;
    }
    if (typeof message === 'string') {
        throw new Error(message);
    }
    throw message;
}

exports.assert = invariant;
exports.attempt = attempt;
exports.attemptAsync = attemptAsync;
exports.invariant = invariant;
