'use strict';

Object.defineProperty(exports, Symbol.toStringTag, { value: 'Module' });

const isWeakSet = require('../_chunk/isWeakSet-403Sh5.js');
const isPlainObject = require('../_chunk/isPlainObject-Xaozpc.js');

function isBlob(x) {
    if (typeof Blob === 'undefined') {
        return false;
    }
    return x instanceof Blob;
}

function isBoolean(x) {
    return typeof x === 'boolean';
}

function isBrowser() {
    return typeof window !== 'undefined' && window?.document != null;
}

function isError(value) {
    return value instanceof Error;
}

function isFile(x) {
    if (typeof File === 'undefined') {
        return false;
    }
    return isBlob(x) && x instanceof File;
}

function isJSON(value) {
    if (typeof value !== 'string') {
        return false;
    }
    try {
        JSON.parse(value);
        return true;
    }
    catch {
        return false;
    }
}

function isJSONValue(value) {
    switch (typeof value) {
        case 'object': {
            return value === null || isJSONArray(value) || isJSONObject(value);
        }
        case 'string':
        case 'number':
        case 'boolean': {
            return true;
        }
        default: {
            return false;
        }
    }
}
function isJSONArray(value) {
    if (!Array.isArray(value)) {
        return false;
    }
    return value.every(item => isJSONValue(item));
}
function isJSONObject(obj) {
    if (!isPlainObject.isPlainObject(obj)) {
        return false;
    }
    const keys = Reflect.ownKeys(obj);
    for (let i = 0; i < keys.length; i++) {
        const key = keys[i];
        const value = obj[key];
        if (typeof key !== 'string') {
            return false;
        }
        if (!isJSONValue(value)) {
            return false;
        }
    }
    return true;
}

function isNode() {
    return typeof process !== 'undefined' && process?.versions?.node != null;
}

function isNotNil(x) {
    return x != null;
}

function isPromise(value) {
    return value instanceof Promise;
}

function isString(value) {
    return typeof value === 'string';
}

exports.isArrayBuffer = isWeakSet.isArrayBuffer;
exports.isBuffer = isWeakSet.isBuffer;
exports.isDate = isWeakSet.isDate;
exports.isEqual = isWeakSet.isEqual;
exports.isEqualWith = isWeakSet.isEqualWith;
exports.isFunction = isWeakSet.isFunction;
exports.isLength = isWeakSet.isLength;
exports.isMap = isWeakSet.isMap;
exports.isNil = isWeakSet.isNil;
exports.isNull = isWeakSet.isNull;
exports.isRegExp = isWeakSet.isRegExp;
exports.isSet = isWeakSet.isSet;
exports.isSymbol = isWeakSet.isSymbol;
exports.isUndefined = isWeakSet.isUndefined;
exports.isWeakMap = isWeakSet.isWeakMap;
exports.isWeakSet = isWeakSet.isWeakSet;
exports.isPlainObject = isPlainObject.isPlainObject;
exports.isPrimitive = isPlainObject.isPrimitive;
exports.isTypedArray = isPlainObject.isTypedArray;
exports.isBlob = isBlob;
exports.isBoolean = isBoolean;
exports.isBrowser = isBrowser;
exports.isError = isError;
exports.isFile = isFile;
exports.isJSON = isJSON;
exports.isJSONArray = isJSONArray;
exports.isJSONObject = isJSONObject;
exports.isJSONValue = isJSONValue;
exports.isNode = isNode;
exports.isNotNil = isNotNil;
exports.isPromise = isPromise;
exports.isString = isString;
