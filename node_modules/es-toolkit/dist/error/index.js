'use strict';

Object.defineProperty(exports, Symbol.toStringTag, { value: 'Module' });

const AbortError = require('../_chunk/AbortError-Cg4ZQ1.js');

class TimeoutError extends Error {
    constructor(message = 'The operation was timed out') {
        super(message);
        this.name = 'TimeoutError';
    }
}

exports.AbortError = AbortError.AbortError;
exports.TimeoutError = TimeoutError;
