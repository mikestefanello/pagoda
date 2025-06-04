'use strict';

const isPlainObject$1 = require('./isPlainObject-Xaozpc.js');

function clone(obj) {
    if (isPlainObject$1.isPrimitive(obj)) {
        return obj;
    }
    if (Array.isArray(obj) ||
        isPlainObject$1.isTypedArray(obj) ||
        obj instanceof ArrayBuffer ||
        (typeof SharedArrayBuffer !== 'undefined' && obj instanceof SharedArrayBuffer)) {
        return obj.slice(0);
    }
    const prototype = Object.getPrototypeOf(obj);
    const Constructor = prototype.constructor;
    if (obj instanceof Date || obj instanceof Map || obj instanceof Set) {
        return new Constructor(obj);
    }
    if (obj instanceof RegExp) {
        const newRegExp = new Constructor(obj);
        newRegExp.lastIndex = obj.lastIndex;
        return newRegExp;
    }
    if (obj instanceof DataView) {
        return new Constructor(obj.buffer.slice(0));
    }
    if (obj instanceof Error) {
        const newError = new Constructor(obj.message);
        newError.stack = obj.stack;
        newError.name = obj.name;
        newError.cause = obj.cause;
        return newError;
    }
    if (typeof File !== 'undefined' && obj instanceof File) {
        const newFile = new Constructor([obj], obj.name, { type: obj.type, lastModified: obj.lastModified });
        return newFile;
    }
    if (typeof obj === 'object') {
        const newObject = Object.create(prototype);
        return Object.assign(newObject, obj);
    }
    return obj;
}

function cloneDeepWith(obj, cloneValue) {
    return cloneDeepWithImpl(obj, undefined, obj, new Map(), cloneValue);
}
function cloneDeepWithImpl(valueToClone, keyToClone, objectToClone, stack = new Map(), cloneValue = undefined) {
    const cloned = cloneValue?.(valueToClone, keyToClone, objectToClone, stack);
    if (cloned != null) {
        return cloned;
    }
    if (isPlainObject$1.isPrimitive(valueToClone)) {
        return valueToClone;
    }
    if (stack.has(valueToClone)) {
        return stack.get(valueToClone);
    }
    if (Array.isArray(valueToClone)) {
        const result = new Array(valueToClone.length);
        stack.set(valueToClone, result);
        for (let i = 0; i < valueToClone.length; i++) {
            result[i] = cloneDeepWithImpl(valueToClone[i], i, objectToClone, stack, cloneValue);
        }
        if (Object.hasOwn(valueToClone, 'index')) {
            result.index = valueToClone.index;
        }
        if (Object.hasOwn(valueToClone, 'input')) {
            result.input = valueToClone.input;
        }
        return result;
    }
    if (valueToClone instanceof Date) {
        return new Date(valueToClone.getTime());
    }
    if (valueToClone instanceof RegExp) {
        const result = new RegExp(valueToClone.source, valueToClone.flags);
        result.lastIndex = valueToClone.lastIndex;
        return result;
    }
    if (valueToClone instanceof Map) {
        const result = new Map();
        stack.set(valueToClone, result);
        for (const [key, value] of valueToClone) {
            result.set(key, cloneDeepWithImpl(value, key, objectToClone, stack, cloneValue));
        }
        return result;
    }
    if (valueToClone instanceof Set) {
        const result = new Set();
        stack.set(valueToClone, result);
        for (const value of valueToClone) {
            result.add(cloneDeepWithImpl(value, undefined, objectToClone, stack, cloneValue));
        }
        return result;
    }
    if (typeof Buffer !== 'undefined' && Buffer.isBuffer(valueToClone)) {
        return valueToClone.subarray();
    }
    if (isPlainObject$1.isTypedArray(valueToClone)) {
        const result = new (Object.getPrototypeOf(valueToClone).constructor)(valueToClone.length);
        stack.set(valueToClone, result);
        for (let i = 0; i < valueToClone.length; i++) {
            result[i] = cloneDeepWithImpl(valueToClone[i], i, objectToClone, stack, cloneValue);
        }
        return result;
    }
    if (valueToClone instanceof ArrayBuffer ||
        (typeof SharedArrayBuffer !== 'undefined' && valueToClone instanceof SharedArrayBuffer)) {
        return valueToClone.slice(0);
    }
    if (valueToClone instanceof DataView) {
        const result = new DataView(valueToClone.buffer.slice(0), valueToClone.byteOffset, valueToClone.byteLength);
        stack.set(valueToClone, result);
        copyProperties(result, valueToClone, objectToClone, stack, cloneValue);
        return result;
    }
    if (typeof File !== 'undefined' && valueToClone instanceof File) {
        const result = new File([valueToClone], valueToClone.name, {
            type: valueToClone.type,
        });
        stack.set(valueToClone, result);
        copyProperties(result, valueToClone, objectToClone, stack, cloneValue);
        return result;
    }
    if (valueToClone instanceof Blob) {
        const result = new Blob([valueToClone], { type: valueToClone.type });
        stack.set(valueToClone, result);
        copyProperties(result, valueToClone, objectToClone, stack, cloneValue);
        return result;
    }
    if (valueToClone instanceof Error) {
        const result = new valueToClone.constructor();
        stack.set(valueToClone, result);
        result.message = valueToClone.message;
        result.name = valueToClone.name;
        result.stack = valueToClone.stack;
        result.cause = valueToClone.cause;
        copyProperties(result, valueToClone, objectToClone, stack, cloneValue);
        return result;
    }
    if (typeof valueToClone === 'object' && isCloneableObject(valueToClone)) {
        const result = Object.create(Object.getPrototypeOf(valueToClone));
        stack.set(valueToClone, result);
        copyProperties(result, valueToClone, objectToClone, stack, cloneValue);
        return result;
    }
    return valueToClone;
}
function copyProperties(target, source, objectToClone = target, stack, cloneValue) {
    const keys = [...Object.keys(source), ...isPlainObject$1.getSymbols(source)];
    for (let i = 0; i < keys.length; i++) {
        const key = keys[i];
        const descriptor = Object.getOwnPropertyDescriptor(target, key);
        if (descriptor == null || descriptor.writable) {
            target[key] = cloneDeepWithImpl(source[key], key, objectToClone, stack, cloneValue);
        }
    }
}
function isCloneableObject(object) {
    switch (isPlainObject$1.getTag(object)) {
        case isPlainObject$1.argumentsTag:
        case isPlainObject$1.arrayTag:
        case isPlainObject$1.arrayBufferTag:
        case isPlainObject$1.dataViewTag:
        case isPlainObject$1.booleanTag:
        case isPlainObject$1.dateTag:
        case isPlainObject$1.float32ArrayTag:
        case isPlainObject$1.float64ArrayTag:
        case isPlainObject$1.int8ArrayTag:
        case isPlainObject$1.int16ArrayTag:
        case isPlainObject$1.int32ArrayTag:
        case isPlainObject$1.mapTag:
        case isPlainObject$1.numberTag:
        case isPlainObject$1.objectTag:
        case isPlainObject$1.regexpTag:
        case isPlainObject$1.setTag:
        case isPlainObject$1.stringTag:
        case isPlainObject$1.symbolTag:
        case isPlainObject$1.uint8ArrayTag:
        case isPlainObject$1.uint8ClampedArrayTag:
        case isPlainObject$1.uint16ArrayTag:
        case isPlainObject$1.uint32ArrayTag: {
            return true;
        }
        default: {
            return false;
        }
    }
}

function cloneDeep(obj) {
    return cloneDeepWithImpl(obj, undefined, obj, new Map(), undefined);
}

function findKey(obj, predicate) {
    const keys = Object.keys(obj);
    return keys.find(key => predicate(obj[key], key, obj));
}

function invert(obj) {
    const result = {};
    const keys = Object.keys(obj);
    for (let i = 0; i < keys.length; i++) {
        const key = keys[i];
        const value = obj[key];
        result[value] = key;
    }
    return result;
}

function mapKeys(object, getNewKey) {
    const result = {};
    const keys = Object.keys(object);
    for (let i = 0; i < keys.length; i++) {
        const key = keys[i];
        const value = object[key];
        result[getNewKey(value, key, object)] = value;
    }
    return result;
}

function mapValues(object, getNewValue) {
    const result = {};
    const keys = Object.keys(object);
    for (let i = 0; i < keys.length; i++) {
        const key = keys[i];
        const value = object[key];
        result[key] = getNewValue(value, key, object);
    }
    return result;
}

function isObjectLike(value) {
    return typeof value === 'object' && value !== null;
}

function isArray(value) {
    return Array.isArray(value);
}

function isPlainObject(object) {
    if (typeof object !== 'object') {
        return false;
    }
    if (object == null) {
        return false;
    }
    if (Object.getPrototypeOf(object) === null) {
        return true;
    }
    if (Object.prototype.toString.call(object) !== '[object Object]') {
        const tag = object[Symbol.toStringTag];
        if (tag == null) {
            return false;
        }
        const isTagReadonly = !Object.getOwnPropertyDescriptor(object, Symbol.toStringTag)?.writable;
        if (isTagReadonly) {
            return false;
        }
        return object.toString() === `[object ${tag}]`;
    }
    let proto = object;
    while (Object.getPrototypeOf(proto) !== null) {
        proto = Object.getPrototypeOf(proto);
    }
    return Object.getPrototypeOf(object) === proto;
}

exports.clone = clone;
exports.cloneDeep = cloneDeep;
exports.cloneDeepWith = cloneDeepWith;
exports.copyProperties = copyProperties;
exports.findKey = findKey;
exports.invert = invert;
exports.isArray = isArray;
exports.isObjectLike = isObjectLike;
exports.isPlainObject = isPlainObject;
exports.mapKeys = mapKeys;
exports.mapValues = mapValues;
