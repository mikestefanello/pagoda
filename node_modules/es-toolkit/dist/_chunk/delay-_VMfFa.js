'use strict';

const AbortError = require('./AbortError-Cg4ZQ1.js');

function delay(ms, { signal } = {}) {
    return new Promise((resolve, reject) => {
        const abortError = () => {
            reject(new AbortError.AbortError());
        };
        const abortHandler = () => {
            clearTimeout(timeoutId);
            abortError();
        };
        if (signal?.aborted) {
            return abortError();
        }
        const timeoutId = setTimeout(() => {
            signal?.removeEventListener('abort', abortHandler);
            resolve();
        }, ms);
        signal?.addEventListener('abort', abortHandler, { once: true });
    });
}

exports.delay = delay;
