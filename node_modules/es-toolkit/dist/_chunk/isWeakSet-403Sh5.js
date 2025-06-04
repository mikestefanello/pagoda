'use strict';

const isPlainObject = require('./isPlainObject-Xaozpc.js');
const noop = require('./noop-2IwLUk.js');

function isArrayBuffer(value) {
    return value instanceof ArrayBuffer;
}

function isBuffer(x) {
    return typeof Buffer !== 'undefined' && Buffer.isBuffer(x);
}

function isDate(value) {
    return value instanceof Date;
}

function eq(value, other) {
    return value === other || (Number.isNaN(value) && Number.isNaN(other));
}

function isEqualWith(a, b, areValuesEqual) {
    return isEqualWithImpl(a, b, undefined, undefined, undefined, undefined, areValuesEqual);
}
function isEqualWithImpl(a, b, property, aParent, bParent, stack, areValuesEqual) {
    const result = areValuesEqual(a, b, property, aParent, bParent, stack);
    if (result !== undefined) {
        return result;
    }
    if (typeof a === typeof b) {
        switch (typeof a) {
            case 'bigint':
            case 'string':
            case 'boolean':
            case 'symbol':
            case 'undefined': {
                return a === b;
            }
            case 'number': {
                return a === b || Object.is(a, b);
            }
            case 'function': {
                return a === b;
            }
            case 'object': {
                return areObjectsEqual(a, b, stack, areValuesEqual);
            }
        }
    }
    return areObjectsEqual(a, b, stack, areValuesEqual);
}
function areObjectsEqual(a, b, stack, areValuesEqual) {
    if (Object.is(a, b)) {
        return true;
    }
    let aTag = isPlainObject.getTag(a);
    let bTag = isPlainObject.getTag(b);
    if (aTag === isPlainObject.argumentsTag) {
        aTag = isPlainObject.objectTag;
    }
    if (bTag === isPlainObject.argumentsTag) {
        bTag = isPlainObject.objectTag;
    }
    if (aTag !== bTag) {
        return false;
    }
    switch (aTag) {
        case isPlainObject.stringTag:
            return a.toString() === b.toString();
        case isPlainObject.numberTag: {
            const x = a.valueOf();
            const y = b.valueOf();
            return eq(x, y);
        }
        case isPlainObject.booleanTag:
        case isPlainObject.dateTag:
        case isPlainObject.symbolTag:
            return Object.is(a.valueOf(), b.valueOf());
        case isPlainObject.regexpTag: {
            return a.source === b.source && a.flags === b.flags;
        }
        case isPlainObject.functionTag: {
            return a === b;
        }
    }
    stack = stack ?? new Map();
    const aStack = stack.get(a);
    const bStack = stack.get(b);
    if (aStack != null && bStack != null) {
        return aStack === b;
    }
    stack.set(a, b);
    stack.set(b, a);
    try {
        switch (aTag) {
            case isPlainObject.mapTag: {
                if (a.size !== b.size) {
                    return false;
                }
                for (const [key, value] of a.entries()) {
                    if (!b.has(key) || !isEqualWithImpl(value, b.get(key), key, a, b, stack, areValuesEqual)) {
                        return false;
                    }
                }
                return true;
            }
            case isPlainObject.setTag: {
                if (a.size !== b.size) {
                    return false;
                }
                const aValues = Array.from(a.values());
                const bValues = Array.from(b.values());
                for (let i = 0; i < aValues.length; i++) {
                    const aValue = aValues[i];
                    const index = bValues.findIndex(bValue => {
                        return isEqualWithImpl(aValue, bValue, undefined, a, b, stack, areValuesEqual);
                    });
                    if (index === -1) {
                        return false;
                    }
                    bValues.splice(index, 1);
                }
                return true;
            }
            case isPlainObject.arrayTag:
            case isPlainObject.uint8ArrayTag:
            case isPlainObject.uint8ClampedArrayTag:
            case isPlainObject.uint16ArrayTag:
            case isPlainObject.uint32ArrayTag:
            case isPlainObject.bigUint64ArrayTag:
            case isPlainObject.int8ArrayTag:
            case isPlainObject.int16ArrayTag:
            case isPlainObject.int32ArrayTag:
            case isPlainObject.bigInt64ArrayTag:
            case isPlainObject.float32ArrayTag:
            case isPlainObject.float64ArrayTag: {
                if (typeof Buffer !== 'undefined' && Buffer.isBuffer(a) !== Buffer.isBuffer(b)) {
                    return false;
                }
                if (a.length !== b.length) {
                    return false;
                }
                for (let i = 0; i < a.length; i++) {
                    if (!isEqualWithImpl(a[i], b[i], i, a, b, stack, areValuesEqual)) {
                        return false;
                    }
                }
                return true;
            }
            case isPlainObject.arrayBufferTag: {
                if (a.byteLength !== b.byteLength) {
                    return false;
                }
                return areObjectsEqual(new Uint8Array(a), new Uint8Array(b), stack, areValuesEqual);
            }
            case isPlainObject.dataViewTag: {
                if (a.byteLength !== b.byteLength || a.byteOffset !== b.byteOffset) {
                    return false;
                }
                return areObjectsEqual(new Uint8Array(a), new Uint8Array(b), stack, areValuesEqual);
            }
            case isPlainObject.errorTag: {
                return a.name === b.name && a.message === b.message;
            }
            case isPlainObject.objectTag: {
                const areEqualInstances = areObjectsEqual(a.constructor, b.constructor, stack, areValuesEqual) ||
                    (isPlainObject.isPlainObject(a) && isPlainObject.isPlainObject(b));
                if (!areEqualInstances) {
                    return false;
                }
                const aKeys = [...Object.keys(a), ...isPlainObject.getSymbols(a)];
                const bKeys = [...Object.keys(b), ...isPlainObject.getSymbols(b)];
                if (aKeys.length !== bKeys.length) {
                    return false;
                }
                for (let i = 0; i < aKeys.length; i++) {
                    const propKey = aKeys[i];
                    const aProp = a[propKey];
                    if (!Object.hasOwn(b, propKey)) {
                        return false;
                    }
                    const bProp = b[propKey];
                    if (!isEqualWithImpl(aProp, bProp, propKey, a, b, stack, areValuesEqual)) {
                        return false;
                    }
                }
                return true;
            }
            default: {
                return false;
            }
        }
    }
    finally {
        stack.delete(a);
        stack.delete(b);
    }
}

function isEqual(a, b) {
    return isEqualWith(a, b, noop.noop);
}

function isFunction(value) {
    return typeof value === 'function';
}

function isLength(value) {
    return Number.isSafeInteger(value) && value >= 0;
}

function isMap(value) {
    return value instanceof Map;
}

function isNil(x) {
    return x == null;
}

function isNull(x) {
    return x === null;
}

function isRegExp(value) {
    return value instanceof RegExp;
}

function isSet(value) {
    return value instanceof Set;
}

function isSymbol(value) {
    return typeof value === 'symbol';
}

function isUndefined(x) {
    return x === undefined;
}

function isWeakMap(value) {
    return value instanceof WeakMap;
}

function isWeakSet(value) {
    return value instanceof WeakSet;
}

exports.eq = eq;
exports.isArrayBuffer = isArrayBuffer;
exports.isBuffer = isBuffer;
exports.isDate = isDate;
exports.isEqual = isEqual;
exports.isEqualWith = isEqualWith;
exports.isFunction = isFunction;
exports.isLength = isLength;
exports.isMap = isMap;
exports.isNil = isNil;
exports.isNull = isNull;
exports.isRegExp = isRegExp;
exports.isSet = isSet;
exports.isSymbol = isSymbol;
exports.isUndefined = isUndefined;
exports.isWeakMap = isWeakMap;
exports.isWeakSet = isWeakSet;
