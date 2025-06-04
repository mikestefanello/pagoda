'use strict';

Object.defineProperty(exports, Symbol.toStringTag, { value: 'Module' });

const isPlainObject = require('../_chunk/isPlainObject-DINLyA.js');
const isPlainObject$1 = require('../_chunk/isPlainObject-Xaozpc.js');
const snakeCase = require('../_chunk/snakeCase-6cG1f4.js');

function flattenObject(object, { delimiter = '.' } = {}) {
    return flattenObjectImpl(object, '', delimiter);
}
function flattenObjectImpl(object, prefix = '', delimiter = '.') {
    const result = {};
    const keys = Object.keys(object);
    for (let i = 0; i < keys.length; i++) {
        const key = keys[i];
        const value = object[key];
        const prefixedKey = prefix ? `${prefix}${delimiter}${key}` : key;
        if (isPlainObject$1.isPlainObject(value) && Object.keys(value).length > 0) {
            Object.assign(result, flattenObjectImpl(value, prefixedKey, delimiter));
            continue;
        }
        if (Array.isArray(value)) {
            Object.assign(result, flattenObjectImpl(value, prefixedKey, delimiter));
            continue;
        }
        result[prefixedKey] = value;
    }
    return result;
}

function merge(target, source) {
    const sourceKeys = Object.keys(source);
    for (let i = 0; i < sourceKeys.length; i++) {
        const key = sourceKeys[i];
        const sourceValue = source[key];
        const targetValue = target[key];
        if (Array.isArray(sourceValue)) {
            if (Array.isArray(targetValue)) {
                target[key] = merge(targetValue, sourceValue);
            }
            else {
                target[key] = merge([], sourceValue);
            }
        }
        else if (isPlainObject$1.isPlainObject(sourceValue)) {
            if (isPlainObject$1.isPlainObject(targetValue)) {
                target[key] = merge(targetValue, sourceValue);
            }
            else {
                target[key] = merge({}, sourceValue);
            }
        }
        else if (targetValue === undefined || sourceValue !== undefined) {
            target[key] = sourceValue;
        }
    }
    return target;
}

function mergeWith(target, source, merge) {
    const sourceKeys = Object.keys(source);
    for (let i = 0; i < sourceKeys.length; i++) {
        const key = sourceKeys[i];
        const sourceValue = source[key];
        const targetValue = target[key];
        const merged = merge(targetValue, sourceValue, key, target, source);
        if (merged != null) {
            target[key] = merged;
        }
        else if (Array.isArray(sourceValue)) {
            target[key] = mergeWith(targetValue ?? [], sourceValue, merge);
        }
        else if (isPlainObject.isObjectLike(targetValue) && isPlainObject.isObjectLike(sourceValue)) {
            target[key] = mergeWith(targetValue ?? {}, sourceValue, merge);
        }
        else if (targetValue === undefined || sourceValue !== undefined) {
            target[key] = sourceValue;
        }
    }
    return target;
}

function omit(obj, keys) {
    const result = { ...obj };
    for (let i = 0; i < keys.length; i++) {
        const key = keys[i];
        delete result[key];
    }
    return result;
}

function omitBy(obj, shouldOmit) {
    const result = {};
    const keys = Object.keys(obj);
    for (let i = 0; i < keys.length; i++) {
        const key = keys[i];
        const value = obj[key];
        if (!shouldOmit(value, key)) {
            result[key] = value;
        }
    }
    return result;
}

function pick(obj, keys) {
    const result = {};
    for (let i = 0; i < keys.length; i++) {
        const key = keys[i];
        if (Object.hasOwn(obj, key)) {
            result[key] = obj[key];
        }
    }
    return result;
}

function pickBy(obj, shouldPick) {
    const result = {};
    const keys = Object.keys(obj);
    for (let i = 0; i < keys.length; i++) {
        const key = keys[i];
        const value = obj[key];
        if (shouldPick(value, key)) {
            result[key] = value;
        }
    }
    return result;
}

function toCamelCaseKeys(obj) {
    if (isPlainObject.isArray(obj)) {
        return obj.map(item => toCamelCaseKeys(item));
    }
    if (isPlainObject$1.isPlainObject(obj)) {
        const result = {};
        const keys = Object.keys(obj);
        for (let i = 0; i < keys.length; i++) {
            const key = keys[i];
            const camelKey = snakeCase.camelCase(key);
            const camelCaseKeys = toCamelCaseKeys(obj[key]);
            result[camelKey] = camelCaseKeys;
        }
        return result;
    }
    return obj;
}

function toMerged(target, source) {
    return merge(isPlainObject.cloneDeep(target), source);
}

function toSnakeCaseKeys(obj) {
    if (isPlainObject.isArray(obj)) {
        return obj.map(item => toSnakeCaseKeys(item));
    }
    if (isPlainObject.isPlainObject(obj)) {
        const result = {};
        const keys = Object.keys(obj);
        for (let i = 0; i < keys.length; i++) {
            const key = keys[i];
            const snakeKey = snakeCase.snakeCase(key);
            const snakeCaseKeys = toSnakeCaseKeys(obj[key]);
            result[snakeKey] = snakeCaseKeys;
        }
        return result;
    }
    return obj;
}

exports.clone = isPlainObject.clone;
exports.cloneDeep = isPlainObject.cloneDeep;
exports.cloneDeepWith = isPlainObject.cloneDeepWith;
exports.findKey = isPlainObject.findKey;
exports.invert = isPlainObject.invert;
exports.mapKeys = isPlainObject.mapKeys;
exports.mapValues = isPlainObject.mapValues;
exports.flattenObject = flattenObject;
exports.merge = merge;
exports.mergeWith = mergeWith;
exports.omit = omit;
exports.omitBy = omitBy;
exports.pick = pick;
exports.pickBy = pickBy;
exports.toCamelCaseKeys = toCamelCaseKeys;
exports.toMerged = toMerged;
exports.toSnakeCaseKeys = toSnakeCaseKeys;
