'use strict';

Object.defineProperties(exports, { __esModule: { value: true }, [Symbol.toStringTag]: { value: 'Module' } });

const zip$1 = require('../_chunk/zip-Cyyp17.js');
const isWeakSet$1 = require('../_chunk/isWeakSet-403Sh5.js');
const isPlainObject = require('../_chunk/isPlainObject-DINLyA.js');
const unary = require('../_chunk/unary-BVQ0iC.js');
const isPlainObject$1 = require('../_chunk/isPlainObject-Xaozpc.js');
const range$1 = require('../_chunk/range-DSpBDL.js');
const randomInt = require('../_chunk/randomInt-CF7bZK.js');
const noop = require('../_chunk/noop-2IwLUk.js');
const snakeCase$1 = require('../_chunk/snakeCase-6cG1f4.js');
const upperFirst$1 = require('../_chunk/upperFirst-Cx78bs.js');

function castArray(value) {
    if (arguments.length === 0) {
        return [];
    }
    return Array.isArray(value) ? value : [value];
}

function toArray$1(value) {
    return Array.isArray(value) ? value : Array.from(value);
}

function isArrayLike(value) {
    return value != null && typeof value !== 'function' && isWeakSet$1.isLength(value.length);
}

function chunk(arr, size = 1) {
    size = Math.max(Math.floor(size), 0);
    if (size === 0 || !isArrayLike(arr)) {
        return [];
    }
    return zip$1.chunk(toArray$1(arr), size);
}

function compact(arr) {
    if (!isArrayLike(arr)) {
        return [];
    }
    return zip$1.compact(Array.from(arr));
}

function concat(...values) {
    return zip$1.flatten(values);
}

function isArrayLikeObject(value) {
    return isPlainObject.isObjectLike(value) && isArrayLike(value);
}

function difference(arr, ...values) {
    if (!isArrayLikeObject(arr)) {
        return [];
    }
    const arr1 = toArray$1(arr);
    const arr2 = [];
    for (let i = 0; i < values.length; i++) {
        const value = values[i];
        if (isArrayLikeObject(value)) {
            arr2.push(...Array.from(value));
        }
    }
    return zip$1.difference(arr1, arr2);
}

function last(array) {
    if (!isArrayLike(array)) {
        return undefined;
    }
    return zip$1.last(toArray$1(array));
}

function flattenArrayLike(values) {
    const result = [];
    for (let i = 0; i < values.length; i++) {
        const arrayLike = values[i];
        if (!isArrayLikeObject(arrayLike)) {
            continue;
        }
        for (let j = 0; j < arrayLike.length; j++) {
            result.push(arrayLike[j]);
        }
    }
    return result;
}

function isDeepKey(key) {
    switch (typeof key) {
        case 'number':
        case 'symbol': {
            return false;
        }
        case 'string': {
            return key.includes('.') || key.includes('[') || key.includes(']');
        }
    }
}

function toKey(value) {
    if (typeof value === 'string' || typeof value === 'symbol') {
        return value;
    }
    if (Object.is(value?.valueOf?.(), -0)) {
        return '-0';
    }
    return String(value);
}

function toPath(deepKey) {
    const result = [];
    const length = deepKey.length;
    if (length === 0) {
        return result;
    }
    let index = 0;
    let key = '';
    let quoteChar = '';
    let bracket = false;
    if (deepKey.charCodeAt(0) === 46) {
        result.push('');
        index++;
    }
    while (index < length) {
        const char = deepKey[index];
        if (quoteChar) {
            if (char === '\\' && index + 1 < length) {
                index++;
                key += deepKey[index];
            }
            else if (char === quoteChar) {
                quoteChar = '';
            }
            else {
                key += char;
            }
        }
        else if (bracket) {
            if (char === '"' || char === "'") {
                quoteChar = char;
            }
            else if (char === ']') {
                bracket = false;
                result.push(key);
                key = '';
            }
            else {
                key += char;
            }
        }
        else {
            if (char === '[') {
                bracket = true;
                if (key) {
                    result.push(key);
                    key = '';
                }
            }
            else if (char === '.') {
                if (key) {
                    result.push(key);
                    key = '';
                }
            }
            else {
                key += char;
            }
        }
        index++;
    }
    if (key) {
        result.push(key);
    }
    return result;
}

function get(object, path, defaultValue) {
    if (object == null) {
        return defaultValue;
    }
    switch (typeof path) {
        case 'string': {
            const result = object[path];
            if (result === undefined) {
                if (isDeepKey(path)) {
                    return get(object, toPath(path), defaultValue);
                }
                else {
                    return defaultValue;
                }
            }
            return result;
        }
        case 'number':
        case 'symbol': {
            if (typeof path === 'number') {
                path = toKey(path);
            }
            const result = object[path];
            if (result === undefined) {
                return defaultValue;
            }
            return result;
        }
        default: {
            if (Array.isArray(path)) {
                return getWithPath(object, path, defaultValue);
            }
            if (Object.is(path?.valueOf(), -0)) {
                path = '-0';
            }
            else {
                path = String(path);
            }
            const result = object[path];
            if (result === undefined) {
                return defaultValue;
            }
            return result;
        }
    }
}
function getWithPath(object, path, defaultValue) {
    if (path.length === 0) {
        return defaultValue;
    }
    let current = object;
    for (let index = 0; index < path.length; index++) {
        if (current == null) {
            return defaultValue;
        }
        current = current[path[index]];
    }
    if (current === undefined) {
        return defaultValue;
    }
    return current;
}

function property(path) {
    return function (object) {
        return get(object, path);
    };
}

function isObject(value) {
    return value !== null && (typeof value === 'object' || typeof value === 'function');
}

function isMatch(target, source) {
    if (source === target) {
        return true;
    }
    switch (typeof source) {
        case 'object': {
            if (source == null) {
                return true;
            }
            const keys = Object.keys(source);
            if (target == null) {
                return keys.length === 0;
            }
            if (Array.isArray(source)) {
                return isArrayMatch(target, source);
            }
            if (source instanceof Map) {
                return isMapMatch(target, source);
            }
            if (source instanceof Set) {
                return isSetMatch(target, source);
            }
            for (let i = 0; i < keys.length; i++) {
                const key = keys[i];
                if (!isPlainObject$1.isPrimitive(target) && !(key in target)) {
                    return false;
                }
                if (source[key] === undefined && target[key] !== undefined) {
                    return false;
                }
                if (source[key] === null && target[key] !== null) {
                    return false;
                }
                if (!isMatch(target[key], source[key])) {
                    return false;
                }
            }
            return true;
        }
        case 'function': {
            if (Object.keys(source).length > 0) {
                return isMatch(target, { ...source });
            }
            return false;
        }
        default: {
            if (!isObject(target)) {
                return isWeakSet$1.eq(target, source);
            }
            return !source;
        }
    }
}
function isMapMatch(target, source) {
    if (source.size === 0) {
        return true;
    }
    if (!(target instanceof Map)) {
        return false;
    }
    for (const [key, value] of source.entries()) {
        if (!isMatch(target.get(key), value)) {
            return false;
        }
    }
    return true;
}
function isArrayMatch(target, source) {
    if (source.length === 0) {
        return true;
    }
    if (!Array.isArray(target)) {
        return false;
    }
    const countedIndex = new Set();
    for (let i = 0; i < source.length; i++) {
        const sourceItem = source[i];
        const index = target.findIndex((targetItem, index) => {
            return isMatch(targetItem, sourceItem) && !countedIndex.has(index);
        });
        if (index === -1) {
            return false;
        }
        countedIndex.add(index);
    }
    return true;
}
function isSetMatch(target, source) {
    if (source.size === 0) {
        return true;
    }
    if (!(target instanceof Set)) {
        return false;
    }
    return isArrayMatch([...target], [...source]);
}

function matches(source) {
    source = isPlainObject.cloneDeep(source);
    return (target) => {
        return isMatch(target, source);
    };
}

function cloneDeepWith(obj, cloneValue) {
    return isPlainObject.cloneDeepWith(obj, (value, key, object, stack) => {
        const cloned = cloneValue?.(value, key, object, stack);
        if (cloned != null) {
            return cloned;
        }
        if (typeof obj !== 'object') {
            return undefined;
        }
        switch (Object.prototype.toString.call(obj)) {
            case isPlainObject$1.numberTag:
            case isPlainObject$1.stringTag:
            case isPlainObject$1.booleanTag: {
                const result = new obj.constructor(obj?.valueOf());
                isPlainObject.copyProperties(result, obj);
                return result;
            }
            case isPlainObject$1.argumentsTag: {
                const result = {};
                isPlainObject.copyProperties(result, obj);
                result.length = obj.length;
                result[Symbol.iterator] = obj[Symbol.iterator];
                return result;
            }
            default: {
                return undefined;
            }
        }
    });
}

function cloneDeep(obj) {
    return cloneDeepWith(obj);
}

const IS_UNSIGNED_INTEGER = /^(?:0|[1-9]\d*)$/;
function isIndex(value, length = Number.MAX_SAFE_INTEGER) {
    switch (typeof value) {
        case 'number': {
            return Number.isInteger(value) && value >= 0 && value < length;
        }
        case 'symbol': {
            return false;
        }
        case 'string': {
            return IS_UNSIGNED_INTEGER.test(value);
        }
    }
}

function isArguments(value) {
    return value !== null && typeof value === 'object' && isPlainObject$1.getTag(value) === '[object Arguments]';
}

function has(object, path) {
    let resolvedPath;
    if (Array.isArray(path)) {
        resolvedPath = path;
    }
    else if (typeof path === 'string' && isDeepKey(path) && object?.[path] == null) {
        resolvedPath = toPath(path);
    }
    else {
        resolvedPath = [path];
    }
    if (resolvedPath.length === 0) {
        return false;
    }
    let current = object;
    for (let i = 0; i < resolvedPath.length; i++) {
        const key = resolvedPath[i];
        if (current == null || !Object.hasOwn(current, key)) {
            const isSparseIndex = (Array.isArray(current) || isArguments(current)) && isIndex(key) && key < current.length;
            if (!isSparseIndex) {
                return false;
            }
        }
        current = current[key];
    }
    return true;
}

function matchesProperty(property, source) {
    switch (typeof property) {
        case 'object': {
            if (Object.is(property?.valueOf(), -0)) {
                property = '-0';
            }
            break;
        }
        case 'number': {
            property = toKey(property);
            break;
        }
    }
    source = cloneDeep(source);
    return function (target) {
        const result = get(target, property);
        if (result === undefined) {
            return has(target, property);
        }
        if (source === undefined) {
            return result === undefined;
        }
        return isMatch(result, source);
    };
}

function iteratee(value) {
    if (value == null) {
        return unary.identity;
    }
    switch (typeof value) {
        case 'function': {
            return value;
        }
        case 'object': {
            if (Array.isArray(value) && value.length === 2) {
                return matchesProperty(value[0], value[1]);
            }
            return matches(value);
        }
        case 'string':
        case 'symbol':
        case 'number': {
            return property(value);
        }
    }
}

function differenceBy(arr, ..._values) {
    if (!isArrayLikeObject(arr)) {
        return [];
    }
    const iteratee$1 = last(_values);
    const values = flattenArrayLike(_values);
    if (isArrayLikeObject(iteratee$1)) {
        return zip$1.difference(Array.from(arr), values);
    }
    return zip$1.differenceBy(Array.from(arr), values, iteratee(iteratee$1));
}

function differenceWith(array, ...values) {
    if (!isArrayLikeObject(array)) {
        return [];
    }
    const comparator = last(values);
    const flattenedValues = flattenArrayLike(values);
    if (typeof comparator === 'function') {
        return zip$1.differenceWith(Array.from(array), flattenedValues, comparator);
    }
    return zip$1.difference(Array.from(array), flattenedValues);
}

function drop(collection, itemsCount = 1, guard) {
    if (!isArrayLike(collection)) {
        return [];
    }
    itemsCount = guard ? 1 : zip$1.toInteger(itemsCount);
    return zip$1.drop(toArray$1(collection), itemsCount);
}

function dropRight(collection, itemsCount = 1, guard) {
    if (!isArrayLike(collection)) {
        return [];
    }
    itemsCount = guard ? 1 : zip$1.toInteger(itemsCount);
    return zip$1.dropRight(toArray$1(collection), itemsCount);
}

function dropRightWhile(arr, predicate) {
    if (!isArrayLike(arr)) {
        return [];
    }
    return dropRightWhileImpl(Array.from(arr), predicate);
}
function dropRightWhileImpl(arr, predicate) {
    switch (typeof predicate) {
        case 'function': {
            return zip$1.dropRightWhile(arr, (item, index, arr) => Boolean(predicate(item, index, arr)));
        }
        case 'object': {
            if (Array.isArray(predicate) && predicate.length === 2) {
                const key = predicate[0];
                const value = predicate[1];
                return zip$1.dropRightWhile(arr, matchesProperty(key, value));
            }
            else {
                return zip$1.dropRightWhile(arr, matches(predicate));
            }
        }
        case 'symbol':
        case 'number':
        case 'string': {
            return zip$1.dropRightWhile(arr, property(predicate));
        }
    }
}

function dropWhile(arr, predicate) {
    if (!isArrayLike(arr)) {
        return [];
    }
    return dropWhileImpl(toArray$1(arr), predicate);
}
function dropWhileImpl(arr, predicate) {
    switch (typeof predicate) {
        case 'function': {
            return zip$1.dropWhile(arr, (item, index, arr) => Boolean(predicate(item, index, arr)));
        }
        case 'object': {
            if (Array.isArray(predicate) && predicate.length === 2) {
                const key = predicate[0];
                const value = predicate[1];
                return zip$1.dropWhile(arr, matchesProperty(key, value));
            }
            else {
                return zip$1.dropWhile(arr, matches(predicate));
            }
        }
        case 'number':
        case 'symbol':
        case 'string': {
            return zip$1.dropWhile(arr, property(predicate));
        }
    }
}

function isIterateeCall(value, index, object) {
    if (!isObject(object)) {
        return false;
    }
    if ((typeof index === 'number' && isArrayLike(object) && isIndex(index) && index < object.length) ||
        (typeof index === 'string' && index in object)) {
        return isWeakSet$1.eq(object[index], value);
    }
    return false;
}

function every(source, doesMatch, guard) {
    if (!source) {
        return true;
    }
    if (guard && isIterateeCall(source, doesMatch, guard)) {
        doesMatch = undefined;
    }
    if (!doesMatch) {
        doesMatch = unary.identity;
    }
    let predicate;
    switch (typeof doesMatch) {
        case 'function': {
            predicate = doesMatch;
            break;
        }
        case 'object': {
            if (Array.isArray(doesMatch) && doesMatch.length === 2) {
                const key = doesMatch[0];
                const value = doesMatch[1];
                predicate = matchesProperty(key, value);
            }
            else {
                predicate = matches(doesMatch);
            }
            break;
        }
        case 'symbol':
        case 'number':
        case 'string': {
            predicate = property(doesMatch);
        }
    }
    if (!isArrayLike(source)) {
        const keys = Object.keys(source);
        for (let i = 0; i < keys.length; i++) {
            const key = keys[i];
            const value = source[key];
            if (!predicate(value, key, source)) {
                return false;
            }
        }
        return true;
    }
    for (let i = 0; i < source.length; i++) {
        if (!predicate(source[i], i, source)) {
            return false;
        }
    }
    return true;
}

function isString(value) {
    return typeof value === 'string' || value instanceof String;
}

function fill(array, value, start = 0, end = array ? array.length : 0) {
    if (!isArrayLike(array)) {
        return [];
    }
    if (isString(array)) {
        return array;
    }
    start = Math.floor(start);
    end = Math.floor(end);
    if (!start) {
        start = 0;
    }
    if (!end) {
        end = 0;
    }
    return zip$1.fill(array, value, start, end);
}

function filter(source, predicate) {
    if (!source) {
        return [];
    }
    predicate = iteratee(predicate);
    if (!Array.isArray(source)) {
        const result = [];
        const keys = Object.keys(source);
        const length = isArrayLike(source) ? source.length : keys.length;
        for (let i = 0; i < length; i++) {
            const key = keys[i];
            const value = source[key];
            if (predicate(value, key, source)) {
                result.push(value);
            }
        }
        return result;
    }
    const result = [];
    const length = source.length;
    for (let i = 0; i < length; i++) {
        const value = source[i];
        if (predicate(value, i, source)) {
            result.push(value);
        }
    }
    return result;
}

function find(source, _doesMatch, fromIndex = 0) {
    if (!source) {
        return undefined;
    }
    if (fromIndex < 0) {
        fromIndex = Math.max(source.length + fromIndex, 0);
    }
    const doesMatch = iteratee(_doesMatch);
    if (typeof doesMatch === 'function' && !Array.isArray(source)) {
        const keys = Object.keys(source);
        for (let i = fromIndex; i < keys.length; i++) {
            const key = keys[i];
            const value = source[key];
            if (doesMatch(value, key, source)) {
                return value;
            }
        }
        return undefined;
    }
    const values = Array.isArray(source) ? source.slice(fromIndex) : Object.values(source).slice(fromIndex);
    return values.find(doesMatch);
}

function findIndex(arr, doesMatch, fromIndex = 0) {
    if (!arr) {
        return -1;
    }
    if (fromIndex < 0) {
        fromIndex = Math.max(arr.length + fromIndex, 0);
    }
    const subArray = Array.from(arr).slice(fromIndex);
    let index = -1;
    switch (typeof doesMatch) {
        case 'function': {
            index = subArray.findIndex(doesMatch);
            break;
        }
        case 'object': {
            if (Array.isArray(doesMatch) && doesMatch.length === 2) {
                const key = doesMatch[0];
                const value = doesMatch[1];
                index = subArray.findIndex(matchesProperty(key, value));
            }
            else {
                index = subArray.findIndex(matches(doesMatch));
            }
            break;
        }
        case 'number':
        case 'symbol':
        case 'string': {
            index = subArray.findIndex(property(doesMatch));
        }
    }
    return index === -1 ? -1 : index + fromIndex;
}

function findLast(source, _doesMatch, fromIndex) {
    if (!source) {
        return undefined;
    }
    const length = Array.isArray(source) ? source.length : Object.keys(source).length;
    fromIndex = zip$1.toInteger(fromIndex ?? length - 1);
    if (fromIndex < 0) {
        fromIndex = Math.max(length + fromIndex, 0);
    }
    else {
        fromIndex = Math.min(fromIndex, length - 1);
    }
    const doesMatch = iteratee(_doesMatch);
    if (typeof doesMatch === 'function' && !Array.isArray(source)) {
        const keys = Object.keys(source);
        for (let i = fromIndex; i >= 0; i--) {
            const key = keys[i];
            const value = source[key];
            if (doesMatch(value, key, source)) {
                return value;
            }
        }
        return undefined;
    }
    const values = Array.isArray(source) ? source.slice(0, fromIndex + 1) : Object.values(source).slice(0, fromIndex + 1);
    return values.findLast(doesMatch);
}

function findLastIndex(arr, doesMatch, fromIndex = arr ? arr.length - 1 : 0) {
    if (!arr) {
        return -1;
    }
    if (fromIndex < 0) {
        fromIndex = Math.max(arr.length + fromIndex, 0);
    }
    else {
        fromIndex = Math.min(fromIndex, arr.length - 1);
    }
    const subArray = toArray$1(arr).slice(0, fromIndex + 1);
    switch (typeof doesMatch) {
        case 'function': {
            return subArray.findLastIndex(doesMatch);
        }
        case 'object': {
            if (Array.isArray(doesMatch) && doesMatch.length === 2) {
                const key = doesMatch[0];
                const value = doesMatch[1];
                return subArray.findLastIndex(matchesProperty(key, value));
            }
            else {
                return subArray.findLastIndex(matches(doesMatch));
            }
        }
        case 'number':
        case 'symbol':
        case 'string': {
            return subArray.findLastIndex(property(doesMatch));
        }
    }
}

function flatten(value, depth = 1) {
    const result = [];
    const flooredDepth = Math.floor(depth);
    if (!isArrayLike(value)) {
        return result;
    }
    const recursive = (arr, currentDepth) => {
        for (let i = 0; i < arr.length; i++) {
            const item = arr[i];
            if (currentDepth < flooredDepth &&
                (Array.isArray(item) ||
                    Boolean(item?.[Symbol.isConcatSpreadable]) ||
                    (item !== null && typeof item === 'object' && Object.prototype.toString.call(item) === '[object Arguments]'))) {
                if (Array.isArray(item)) {
                    recursive(item, currentDepth + 1);
                }
                else {
                    recursive(Array.from(item), currentDepth + 1);
                }
            }
            else {
                result.push(item);
            }
        }
    };
    recursive(Array.from(value), 0);
    return result;
}

function map(collection, _iteratee) {
    if (!collection) {
        return [];
    }
    const keys = isArrayLike(collection) || Array.isArray(collection) ? range$1.range(0, collection.length) : Object.keys(collection);
    const iteratee$1 = iteratee(_iteratee ?? unary.identity);
    const result = new Array(keys.length);
    for (let i = 0; i < keys.length; i++) {
        const key = keys[i];
        const value = collection[key];
        result[i] = iteratee$1(value, key, collection);
    }
    return result;
}

function flatMap(collection, iteratee) {
    if (isWeakSet$1.isNil(collection)) {
        return [];
    }
    const mapped = isWeakSet$1.isNil(iteratee) ? map(collection) : map(collection, iteratee);
    return flatten(mapped, 1);
}

function flattenDeep(value) {
    return flatten(value, Infinity);
}

function flattenDepth(value, depth = 1) {
    return flatten(value, depth);
}

function forEach(collection, callback = unary.identity) {
    if (!collection) {
        return collection;
    }
    const keys = isArrayLike(collection) || Array.isArray(collection) ? range$1.range(0, collection.length) : Object.keys(collection);
    for (let i = 0; i < keys.length; i++) {
        const key = keys[i];
        const value = collection[key];
        const result = callback(value, key, collection);
        if (result === false) {
            break;
        }
    }
    return collection;
}

function forEachRight(collection, callback = unary.identity) {
    if (!collection) {
        return collection;
    }
    const keys = isArrayLike(collection) ? range$1.range(0, collection.length) : Object.keys(collection);
    for (let i = keys.length - 1; i >= 0; i--) {
        const key = keys[i];
        const value = collection[key];
        const result = callback(value, key, collection);
        if (result === false) {
            break;
        }
    }
    return collection;
}

function groupBy(source, _getKeyFromItem) {
    if (source == null) {
        return {};
    }
    const items = isArrayLike(source) ? Array.from(source) : Object.values(source);
    const getKeyFromItem = iteratee(_getKeyFromItem ?? unary.identity);
    return zip$1.groupBy(items, getKeyFromItem);
}

function head(arr) {
    if (!isArrayLike(arr)) {
        return undefined;
    }
    return zip$1.head(toArray$1(arr));
}

function includes(source, target, fromIndex, guard) {
    if (source == null) {
        return false;
    }
    if (guard || !fromIndex) {
        fromIndex = 0;
    }
    else {
        fromIndex = zip$1.toInteger(fromIndex);
    }
    if (isString(source)) {
        if (fromIndex > source.length || target instanceof RegExp) {
            return false;
        }
        if (fromIndex < 0) {
            fromIndex = Math.max(0, source.length + fromIndex);
        }
        return source.includes(target, fromIndex);
    }
    if (Array.isArray(source)) {
        return source.includes(target, fromIndex);
    }
    const keys = Object.keys(source);
    if (fromIndex < 0) {
        fromIndex = Math.max(0, keys.length + fromIndex);
    }
    for (let i = fromIndex; i < keys.length; i++) {
        const value = Reflect.get(source, keys[i]);
        if (isWeakSet$1.eq(value, target)) {
            return true;
        }
    }
    return false;
}

function indexOf(array, searchElement, fromIndex) {
    if (!isArrayLike(array)) {
        return -1;
    }
    if (Number.isNaN(searchElement)) {
        fromIndex = fromIndex ?? 0;
        if (fromIndex < 0) {
            fromIndex = Math.max(0, array.length + fromIndex);
        }
        for (let i = fromIndex; i < array.length; i++) {
            if (Number.isNaN(array[i])) {
                return i;
            }
        }
        return -1;
    }
    return Array.from(array).indexOf(searchElement, fromIndex);
}

function initial(arr) {
    if (!isArrayLike(arr)) {
        return [];
    }
    return zip$1.initial(Array.from(arr));
}

function intersection(...arrays) {
    if (arrays.length === 0) {
        return [];
    }
    if (!isArrayLikeObject(arrays[0])) {
        return [];
    }
    let result = zip$1.uniq(Array.from(arrays[0]));
    for (let i = 1; i < arrays.length; i++) {
        const array = arrays[i];
        if (!isArrayLikeObject(array)) {
            return [];
        }
        result = zip$1.intersection(result, Array.from(array));
    }
    return result;
}

function intersectionBy(array, ...values) {
    if (!isArrayLikeObject(array)) {
        return [];
    }
    const lastValue = zip$1.last(values);
    if (lastValue === undefined) {
        return Array.from(array);
    }
    let result = zip$1.uniq(Array.from(array));
    const count = isArrayLikeObject(lastValue) ? values.length : values.length - 1;
    for (let i = 0; i < count; ++i) {
        const value = values[i];
        if (!isArrayLikeObject(value)) {
            return [];
        }
        if (isArrayLikeObject(lastValue)) {
            result = zip$1.intersectionBy(result, Array.from(value), unary.identity);
        }
        else if (typeof lastValue === 'function') {
            result = zip$1.intersectionBy(result, Array.from(value), value => lastValue(value));
        }
        else if (typeof lastValue === 'string') {
            result = zip$1.intersectionBy(result, Array.from(value), property(lastValue));
        }
    }
    return result;
}

function uniq(arr) {
    if (!isArrayLike(arr)) {
        return [];
    }
    return zip$1.uniq(Array.from(arr));
}

function intersectionWith(firstArr, ...otherArrs) {
    if (firstArr == null) {
        return [];
    }
    const _comparator = last(otherArrs);
    let comparator = isWeakSet$1.eq;
    let uniq$1 = uniq;
    if (typeof _comparator === 'function') {
        comparator = _comparator;
        uniq$1 = uniqPreserve0;
        otherArrs.pop();
    }
    let result = uniq$1(Array.from(firstArr));
    for (let i = 0; i < otherArrs.length; ++i) {
        const otherArr = otherArrs[i];
        if (otherArr == null) {
            return [];
        }
        result = zip$1.intersectionWith(result, Array.from(otherArr), comparator);
    }
    return result;
}
function uniqPreserve0(arr) {
    const result = [];
    const added = new Set();
    for (let i = 0; i < arr.length; i++) {
        const item = arr[i];
        if (added.has(item)) {
            continue;
        }
        result.push(item);
        added.add(item);
    }
    return result;
}

function invokeMap(collection, path, ...args) {
    if (isWeakSet$1.isNil(collection)) {
        return [];
    }
    const values = isArrayLike(collection) ? Array.from(collection) : Object.values(collection);
    const result = [];
    for (let i = 0; i < values.length; i++) {
        const value = values[i];
        if (isWeakSet$1.isFunction(path)) {
            result.push(path.apply(value, args));
            continue;
        }
        const method = get(value, path);
        let thisContext = value;
        if (Array.isArray(path)) {
            const pathExceptLast = path.slice(0, -1);
            if (pathExceptLast.length > 0) {
                thisContext = get(value, pathExceptLast);
            }
        }
        else if (typeof path === 'string' && path.includes('.')) {
            const parts = path.split('.');
            const pathExceptLast = parts.slice(0, -1).join('.');
            thisContext = get(value, pathExceptLast);
        }
        result.push(method == null ? undefined : method.apply(thisContext, args));
    }
    return result;
}

function join(array, separator = ',') {
    if (!isArrayLike(array)) {
        return '';
    }
    return Array.from(array).join(separator);
}

function reduce(collection, iteratee = unary.identity, accumulator) {
    if (!collection) {
        return accumulator;
    }
    let keys;
    let startIndex = 0;
    if (isArrayLike(collection)) {
        keys = range$1.range(0, collection.length);
        if (accumulator == null && collection.length > 0) {
            accumulator = collection[0];
            startIndex += 1;
        }
    }
    else {
        keys = Object.keys(collection);
        if (accumulator == null) {
            accumulator = collection[keys[0]];
            startIndex += 1;
        }
    }
    for (let i = startIndex; i < keys.length; i++) {
        const key = keys[i];
        const value = collection[key];
        accumulator = iteratee(accumulator, value, key, collection);
    }
    return accumulator;
}

function keyBy(collection, iteratee$1) {
    if (!isArrayLike(collection) && !isPlainObject.isObjectLike(collection)) {
        return {};
    }
    const keyFn = iteratee(iteratee$1 ?? unary.identity);
    return reduce(collection, (result, value) => {
        const key = keyFn(value);
        result[key] = value;
        return result;
    }, {});
}

function lastIndexOf(array, searchElement, fromIndex) {
    if (!isArrayLike(array) || array.length === 0) {
        return -1;
    }
    const length = array.length;
    let index = fromIndex ?? length - 1;
    if (fromIndex != null) {
        index = index < 0 ? Math.max(length + index, 0) : Math.min(index, length - 1);
    }
    if (Number.isNaN(searchElement)) {
        for (let i = index; i >= 0; i--) {
            if (Number.isNaN(array[i])) {
                return i;
            }
        }
    }
    return Array.from(array).lastIndexOf(searchElement, index);
}

function nth(array, n = 0) {
    if (!isArrayLikeObject(array) || array.length === 0) {
        return undefined;
    }
    n = zip$1.toInteger(n);
    if (n < 0) {
        n += array.length;
    }
    return array[n];
}

function getPriority(a) {
    if (typeof a === 'symbol') {
        return 1;
    }
    if (a === null) {
        return 2;
    }
    if (a === undefined) {
        return 3;
    }
    if (a !== a) {
        return 4;
    }
    return 0;
}
const compareValues = (a, b, order) => {
    if (a !== b) {
        const aPriority = getPriority(a);
        const bPriority = getPriority(b);
        if (aPriority === bPriority && aPriority === 0) {
            if (a < b) {
                return order === 'desc' ? 1 : -1;
            }
            if (a > b) {
                return order === 'desc' ? -1 : 1;
            }
        }
        return order === 'desc' ? bPriority - aPriority : aPriority - bPriority;
    }
    return 0;
};

const regexIsDeepProp = /\.|\[(?:[^[\]]*|(["'])(?:(?!\1)[^\\]|\\.)*?\1)\]/;
const regexIsPlainProp = /^\w*$/;
function isKey(value, object) {
    if (Array.isArray(value)) {
        return false;
    }
    if (typeof value === 'number' || typeof value === 'boolean' || value == null || zip$1.isSymbol(value)) {
        return true;
    }
    return ((typeof value === 'string' && (regexIsPlainProp.test(value) || !regexIsDeepProp.test(value))) ||
        (object != null && Object.hasOwn(object, value)));
}

function orderBy(collection, criteria, orders, guard) {
    if (collection == null) {
        return [];
    }
    orders = guard ? undefined : orders;
    if (!Array.isArray(collection)) {
        collection = Object.values(collection);
    }
    if (!Array.isArray(criteria)) {
        criteria = criteria == null ? [null] : [criteria];
    }
    if (criteria.length === 0) {
        criteria = [null];
    }
    if (!Array.isArray(orders)) {
        orders = orders == null ? [] : [orders];
    }
    orders = orders.map(order => String(order));
    const getValueByNestedPath = (object, path) => {
        let target = object;
        for (let i = 0; i < path.length && target != null; ++i) {
            target = target[path[i]];
        }
        return target;
    };
    const getValueByCriterion = (criterion, object) => {
        if (object == null || criterion == null) {
            return object;
        }
        if (typeof criterion === 'object' && 'key' in criterion) {
            if (Object.hasOwn(object, criterion.key)) {
                return object[criterion.key];
            }
            return getValueByNestedPath(object, criterion.path);
        }
        if (typeof criterion === 'function') {
            return criterion(object);
        }
        if (Array.isArray(criterion)) {
            return getValueByNestedPath(object, criterion);
        }
        if (typeof object === 'object') {
            return object[criterion];
        }
        return object;
    };
    const preparedCriteria = criteria.map(criterion => {
        if (Array.isArray(criterion) && criterion.length === 1) {
            criterion = criterion[0];
        }
        if (criterion == null || typeof criterion === 'function' || Array.isArray(criterion) || isKey(criterion)) {
            return criterion;
        }
        return { key: criterion, path: toPath(criterion) };
    });
    const preparedCollection = collection.map(item => ({
        original: item,
        criteria: preparedCriteria.map(criterion => getValueByCriterion(criterion, item)),
    }));
    return preparedCollection
        .slice()
        .sort((a, b) => {
        for (let i = 0; i < preparedCriteria.length; i++) {
            const comparedResult = compareValues(a.criteria[i], b.criteria[i], orders[i]);
            if (comparedResult !== 0) {
                return comparedResult;
            }
        }
        return 0;
    })
        .map(item => item.original);
}

function partition(source, predicate) {
    if (!source) {
        return [[], []];
    }
    const collection = isArrayLike(source) ? source : Object.values(source);
    predicate = iteratee(predicate);
    const matched = [];
    const unmatched = [];
    for (let i = 0; i < collection.length; i++) {
        const value = collection[i];
        if (predicate(value)) {
            matched.push(value);
        }
        else {
            unmatched.push(value);
        }
    }
    return [matched, unmatched];
}

function pull(arr, ...valuesToRemove) {
    return zip$1.pull(arr, valuesToRemove);
}

function pullAll(arr, valuesToRemove = []) {
    return zip$1.pull(arr, Array.from(valuesToRemove));
}

function pullAllBy(arr, valuesToRemove, _getValue) {
    const getValue = iteratee(_getValue);
    const valuesSet = new Set(Array.from(valuesToRemove).map(x => getValue(x)));
    let resultIndex = 0;
    for (let i = 0; i < arr.length; i++) {
        const value = getValue(arr[i]);
        if (valuesSet.has(value)) {
            continue;
        }
        if (!Object.hasOwn(arr, i)) {
            delete arr[resultIndex++];
            continue;
        }
        arr[resultIndex++] = arr[i];
    }
    arr.length = resultIndex;
    return arr;
}

function copyArray(source, array) {
    const length = source.length;
    if (array == null) {
        array = Array(length);
    }
    for (let i = 0; i < length; i++) {
        array[i] = source[i];
    }
    return array;
}

function pullAllWith(array, values, comparator) {
    if (array?.length == null || values?.length == null) {
        return array;
    }
    if (array === values) {
        values = copyArray(values);
    }
    let resultLength = 0;
    if (comparator == null) {
        comparator = (a, b) => isWeakSet$1.eq(a, b);
    }
    const valuesArray = Array.isArray(values) ? values : Array.from(values);
    const hasUndefined = valuesArray.includes(undefined);
    for (let i = 0; i < array.length; i++) {
        if (i in array) {
            const shouldRemove = valuesArray.some(value => comparator(array[i], value));
            if (!shouldRemove) {
                array[resultLength++] = array[i];
            }
            continue;
        }
        if (!hasUndefined) {
            delete array[resultLength++];
        }
    }
    array.length = resultLength;
    return array;
}

function at(object, ...paths) {
    if (paths.length === 0) {
        return [];
    }
    const allPaths = [];
    for (let i = 0; i < paths.length; i++) {
        const path = paths[i];
        if (!isArrayLike(path) || isString(path)) {
            allPaths.push(path);
            continue;
        }
        for (let j = 0; j < path.length; j++) {
            allPaths.push(path[j]);
        }
    }
    const result = [];
    for (let i = 0; i < allPaths.length; i++) {
        result.push(get(object, allPaths[i]));
    }
    return result;
}

function unset(obj, path) {
    if (obj == null) {
        return true;
    }
    switch (typeof path) {
        case 'symbol':
        case 'number':
        case 'object': {
            if (Array.isArray(path)) {
                return unsetWithPath(obj, path);
            }
            if (typeof path === 'number') {
                path = toKey(path);
            }
            else if (typeof path === 'object') {
                if (Object.is(path?.valueOf(), -0)) {
                    path = '-0';
                }
                else {
                    path = String(path);
                }
            }
            if (obj?.[path] === undefined) {
                return true;
            }
            try {
                delete obj[path];
                return true;
            }
            catch {
                return false;
            }
        }
        case 'string': {
            if (obj?.[path] === undefined && isDeepKey(path)) {
                return unsetWithPath(obj, toPath(path));
            }
            try {
                delete obj[path];
                return true;
            }
            catch {
                return false;
            }
        }
    }
}
function unsetWithPath(obj, path) {
    const parent = get(obj, path.slice(0, -1), obj);
    const lastKey = path[path.length - 1];
    if (parent?.[lastKey] === undefined) {
        return true;
    }
    try {
        delete parent[lastKey];
        return true;
    }
    catch {
        return false;
    }
}

function pullAt(array, ..._indices) {
    const indices = flatten(_indices, 1);
    if (!array) {
        return Array(indices.length);
    }
    const result = at(array, indices);
    const indicesToPull = indices
        .map(index => (isIndex(index, array.length) ? Number(index) : index))
        .sort((a, b) => b - a);
    for (const index of new Set(indicesToPull)) {
        if (isIndex(index, array.length)) {
            Array.prototype.splice.call(array, index, 1);
            continue;
        }
        if (isKey(index, array)) {
            delete array[toKey(index)];
            continue;
        }
        const path = isPlainObject.isArray(index) ? index : toPath(index);
        unset(array, path);
    }
    return result;
}

function reduceRight(collection, iteratee = unary.identity, accumulator) {
    if (!collection) {
        return accumulator;
    }
    let keys;
    let startIndex;
    if (isArrayLike(collection)) {
        keys = range$1.range(0, collection.length).reverse();
        if (accumulator == null && collection.length > 0) {
            accumulator = collection[collection.length - 1];
            startIndex = 1;
        }
        else {
            startIndex = 0;
        }
    }
    else {
        keys = Object.keys(collection).reverse();
        if (accumulator == null) {
            accumulator = collection[keys[0]];
            startIndex = 1;
        }
        else {
            startIndex = 0;
        }
    }
    for (let i = startIndex; i < keys.length; i++) {
        const key = keys[i];
        const value = collection[key];
        accumulator = iteratee(accumulator, value, key, collection);
    }
    return accumulator;
}

function negate(func) {
    if (typeof func !== 'function') {
        throw new TypeError('Expected a function');
    }
    return function (...args) {
        return !func.apply(this, args);
    };
}

function reject(source, predicate) {
    return filter(source, negate(iteratee(predicate)));
}

function remove(arr, shouldRemoveElement) {
    return zip$1.remove(arr, iteratee(shouldRemoveElement));
}

function reverse(array) {
    if (array == null) {
        return array;
    }
    return array.reverse();
}

function sample(collection) {
    if (collection == null) {
        return undefined;
    }
    if (isArrayLike(collection)) {
        return zip$1.sample(toArray$1(collection));
    }
    return zip$1.sample(Object.values(collection));
}

function clamp(value, bound1, bound2) {
    if (Number.isNaN(bound1)) {
        bound1 = 0;
    }
    if (Number.isNaN(bound2)) {
        bound2 = 0;
    }
    return range$1.clamp(value, bound1, bound2);
}

function isMap(value) {
    return isWeakSet$1.isMap(value);
}

function toArray(value) {
    if (value == null) {
        return [];
    }
    if (isArrayLike(value) || isMap(value)) {
        return Array.from(value);
    }
    if (typeof value === 'object') {
        return Object.values(value);
    }
    return [];
}

function sampleSize(collection, size, guard) {
    const arrayCollection = toArray(collection);
    if (guard ? isIterateeCall(collection, size, guard) : size === undefined) {
        size = 1;
    }
    else {
        size = clamp(zip$1.toInteger(size), 0, arrayCollection.length);
    }
    return zip$1.sampleSize(arrayCollection, size);
}

function values(object) {
    return Object.values(object);
}

function isNil(x) {
    return x == null;
}

function shuffle(collection) {
    if (isNil(collection)) {
        return [];
    }
    if (isPlainObject.isArray(collection)) {
        return zip$1.shuffle(collection);
    }
    if (isArrayLike(collection)) {
        return zip$1.shuffle(Array.from(collection));
    }
    if (isPlainObject.isObjectLike(collection)) {
        return zip$1.shuffle(values(collection));
    }
    return [];
}

function size(target) {
    if (isWeakSet$1.isNil(target)) {
        return 0;
    }
    if (target instanceof Map || target instanceof Set) {
        return target.size;
    }
    return Object.keys(target).length;
}

function slice(array, start, end) {
    if (!isArrayLike(array)) {
        return [];
    }
    const length = array.length;
    if (end === undefined) {
        end = length;
    }
    else if (typeof end !== 'number' && isIterateeCall(array, start, end)) {
        start = 0;
        end = length;
    }
    start = zip$1.toInteger(start);
    end = zip$1.toInteger(end);
    if (start < 0) {
        start = Math.max(length + start, 0);
    }
    else {
        start = Math.min(start, length);
    }
    if (end < 0) {
        end = Math.max(length + end, 0);
    }
    else {
        end = Math.min(end, length);
    }
    const resultLength = Math.max(end - start, 0);
    const result = new Array(resultLength);
    for (let i = 0; i < resultLength; ++i) {
        result[i] = array[start + i];
    }
    return result;
}

function some(source, predicate, guard) {
    if (!source) {
        return false;
    }
    if (guard != null) {
        predicate = undefined;
    }
    if (!predicate) {
        predicate = unary.identity;
    }
    const values = Array.isArray(source) ? source : Object.values(source);
    switch (typeof predicate) {
        case 'function': {
            if (!Array.isArray(source)) {
                const keys = Object.keys(source);
                for (let i = 0; i < keys.length; i++) {
                    const key = keys[i];
                    const value = source[key];
                    if (predicate(value, key, source)) {
                        return true;
                    }
                }
                return false;
            }
            for (let i = 0; i < source.length; i++) {
                if (predicate(source[i], i, source)) {
                    return true;
                }
            }
            return false;
        }
        case 'object': {
            if (Array.isArray(predicate) && predicate.length === 2) {
                const key = predicate[0];
                const value = predicate[1];
                const matchFunc = matchesProperty(key, value);
                if (Array.isArray(source)) {
                    for (let i = 0; i < source.length; i++) {
                        if (matchFunc(source[i])) {
                            return true;
                        }
                    }
                    return false;
                }
                return values.some(matchFunc);
            }
            else {
                const matchFunc = matches(predicate);
                if (Array.isArray(source)) {
                    for (let i = 0; i < source.length; i++) {
                        if (matchFunc(source[i])) {
                            return true;
                        }
                    }
                    return false;
                }
                return values.some(matchFunc);
            }
        }
        case 'number':
        case 'symbol':
        case 'string': {
            const propFunc = property(predicate);
            if (Array.isArray(source)) {
                for (let i = 0; i < source.length; i++) {
                    if (propFunc(source[i])) {
                        return true;
                    }
                }
                return false;
            }
            return values.some(propFunc);
        }
    }
}

function sortBy(collection, ...criteria) {
    const length = criteria.length;
    if (length > 1 && isIterateeCall(collection, criteria[0], criteria[1])) {
        criteria = [];
    }
    else if (length > 2 && isIterateeCall(criteria[0], criteria[1], criteria[2])) {
        criteria = [criteria[0]];
    }
    return orderBy(collection, zip$1.flatten(criteria), ['asc']);
}

function isNaN(value) {
    return Number.isNaN(value);
}

const MAX_ARRAY_LENGTH$3 = 4294967295;
const MAX_ARRAY_INDEX = MAX_ARRAY_LENGTH$3 - 1;
function sortedIndexBy(array, value, iteratee$1, retHighest) {
    let low = 0;
    let high = array == null ? 0 : array.length;
    if (high === 0 || isNil(array)) {
        return 0;
    }
    const iterateeFunction = iteratee(iteratee$1);
    const transformedValue = iterateeFunction(value);
    const valIsNaN = isNaN(transformedValue);
    const valIsNull = isWeakSet$1.isNull(transformedValue);
    const valIsSymbol = zip$1.isSymbol(transformedValue);
    const valIsUndefined = isWeakSet$1.isUndefined(transformedValue);
    while (low < high) {
        let setLow;
        const mid = Math.floor((low + high) / 2);
        const computed = iterateeFunction(array[mid]);
        const othIsDefined = !isWeakSet$1.isUndefined(computed);
        const othIsNull = isWeakSet$1.isNull(computed);
        const othIsReflexive = !isNaN(computed);
        const othIsSymbol = zip$1.isSymbol(computed);
        if (valIsNaN) {
            setLow = retHighest || othIsReflexive;
        }
        else if (valIsUndefined) {
            setLow = othIsReflexive && (retHighest || othIsDefined);
        }
        else if (valIsNull) {
            setLow = othIsReflexive && othIsDefined && (retHighest || !othIsNull);
        }
        else if (valIsSymbol) {
            setLow = othIsReflexive && othIsDefined && !othIsNull && (retHighest || !othIsSymbol);
        }
        else if (othIsNull || othIsSymbol) {
            setLow = false;
        }
        else {
            setLow = retHighest ? computed <= transformedValue : computed < transformedValue;
        }
        if (setLow) {
            low = mid + 1;
        }
        else {
            high = mid;
        }
    }
    return Math.min(high, MAX_ARRAY_INDEX);
}

function isNumber(value) {
    return typeof value === 'number' || value instanceof Number;
}

const MAX_ARRAY_LENGTH$2 = 4294967295;
const HALF_MAX_ARRAY_LENGTH$1 = MAX_ARRAY_LENGTH$2 >>> 1;
function sortedIndex(array, value) {
    if (isWeakSet$1.isNil(array)) {
        return 0;
    }
    let low = 0, high = isWeakSet$1.isNil(array) ? low : array.length;
    if (isNumber(value) && value === value && high <= HALF_MAX_ARRAY_LENGTH$1) {
        while (low < high) {
            const mid = (low + high) >>> 1;
            const compute = array[mid];
            if (!isWeakSet$1.isNull(compute) && !isWeakSet$1.isSymbol(compute) && compute < value) {
                low = mid + 1;
            }
            else {
                high = mid;
            }
        }
        return high;
    }
    return sortedIndexBy(array, value, value => value);
}

function sortedIndexOf(array, value) {
    if (!array?.length) {
        return -1;
    }
    const index = sortedIndex(array, value);
    if (index < array.length && isWeakSet$1.eq(array[index], value)) {
        return index;
    }
    return -1;
}

function sortedLastIndexBy(array, value, iteratee) {
    return sortedIndexBy(array, value, iteratee, true);
}

const MAX_ARRAY_LENGTH$1 = 4294967295;
const HALF_MAX_ARRAY_LENGTH = MAX_ARRAY_LENGTH$1 >>> 1;
function sortedLastIndex(array, value) {
    if (isWeakSet$1.isNil(array)) {
        return 0;
    }
    let high = array.length;
    if (!isNumber(value) || Number.isNaN(value) || high > HALF_MAX_ARRAY_LENGTH) {
        return sortedLastIndexBy(array, value, value => value);
    }
    let low = 0;
    while (low < high) {
        const mid = (low + high) >>> 1;
        const compute = array[mid];
        if (!isWeakSet$1.isNull(compute) && !isWeakSet$1.isSymbol(compute) && compute <= value) {
            low = mid + 1;
        }
        else {
            high = mid;
        }
    }
    return high;
}

function sortedLastIndexOf(array, value) {
    if (!array?.length) {
        return -1;
    }
    const index = sortedLastIndex(array, value) - 1;
    if (index >= 0 && isWeakSet$1.eq(array[index], value)) {
        return index;
    }
    return -1;
}

function tail(arr) {
    if (!isArrayLike(arr)) {
        return [];
    }
    return zip$1.tail(toArray$1(arr));
}

function take(arr, count = 1, guard) {
    count = guard ? 1 : zip$1.toInteger(count);
    if (count < 1 || !isArrayLike(arr)) {
        return [];
    }
    return zip$1.take(toArray$1(arr), count);
}

function takeRight(arr, count = 1, guard) {
    count = guard ? 1 : zip$1.toInteger(count);
    if (count <= 0 || !isArrayLike(arr)) {
        return [];
    }
    return zip$1.takeRight(toArray$1(arr), count);
}

function takeRightWhile(_array, predicate) {
    if (!isArrayLikeObject(_array)) {
        return [];
    }
    const array = toArray$1(_array);
    const index = array.findLastIndex(unary.negate(iteratee(predicate)));
    return array.slice(index + 1);
}

function takeWhile(array, predicate) {
    if (!isArrayLikeObject(array)) {
        return [];
    }
    const _array = toArray$1(array);
    const index = _array.findIndex(negate(iteratee(predicate)));
    return index === -1 ? _array : _array.slice(0, index);
}

function union(...arrays) {
    const validArrays = arrays.filter(isArrayLikeObject);
    const flattened = flatten(validArrays, 1);
    return zip$1.uniq(flattened);
}

function unionBy(...values) {
    const lastValue = zip$1.last(values);
    const flattened = flattenArrayLike(values);
    if (isArrayLikeObject(lastValue) || lastValue == null) {
        return zip$1.uniq(flattened);
    }
    return zip$1.uniqBy(flattened, iteratee(lastValue));
}

function unionWith(...values) {
    const lastValue = zip$1.last(values);
    const flattened = flattenArrayLike(values);
    if (isArrayLikeObject(lastValue) || lastValue == null) {
        return zip$1.uniq(flattened);
    }
    return zip$1.uniqWith(flattened, lastValue);
}

function uniqBy(array, iteratee$1) {
    if (!isArrayLikeObject(array)) {
        return [];
    }
    return zip$1.uniqBy(Array.from(array), iteratee(iteratee$1));
}

function uniqWith(arr, comparator) {
    if (!isArrayLike(arr)) {
        return [];
    }
    return typeof comparator === 'function' ? zip$1.uniqWith(Array.from(arr), comparator) : uniq(Array.from(arr));
}

function unzip(array) {
    if (!isArrayLikeObject(array) || !array.length) {
        return [];
    }
    array = isPlainObject.isArray(array) ? array : Array.from(array);
    array = array.filter(item => isArrayLikeObject(item));
    return zip$1.unzip(array);
}

function unzipWith(array, iteratee) {
    if (!isArrayLikeObject(array) || !array.length) {
        return [];
    }
    const unziped = isPlainObject.isArray(array) ? zip$1.unzip(array) : zip$1.unzip(Array.from(array, value => Array.from(value)));
    if (!iteratee) {
        return unziped;
    }
    const result = new Array(unziped.length);
    for (let i = 0; i < unziped.length; i++) {
        const value = unziped[i];
        result[i] = iteratee(...value);
    }
    return result;
}

function without(array, ...values) {
    if (!isArrayLikeObject(array)) {
        return [];
    }
    return zip$1.without(Array.from(array), ...values);
}

function xor(...arrays) {
    const itemCounts = new Map();
    for (let i = 0; i < arrays.length; i++) {
        const array = arrays[i];
        if (!isArrayLikeObject(array)) {
            continue;
        }
        const itemSet = new Set(toArray(array));
        for (const item of itemSet) {
            if (!itemCounts.has(item)) {
                itemCounts.set(item, 1);
            }
            else {
                itemCounts.set(item, itemCounts.get(item) + 1);
            }
        }
    }
    const result = [];
    for (const [item, count] of itemCounts) {
        if (count === 1) {
            result.push(item);
        }
    }
    return result;
}

function xorBy(...values) {
    const lastValue = last(values);
    let mapper = unary.identity;
    if (!isArrayLikeObject(lastValue) && lastValue != null) {
        mapper = iteratee(lastValue);
        values = values.slice(0, -1);
    }
    const arrays = values.filter(isArrayLikeObject);
    const union = unionBy(...arrays, mapper);
    const intersections = zip$1.windowed(arrays, 2).map(([arr1, arr2]) => intersectionBy(arr1, arr2, mapper));
    return differenceBy(union, unionBy(...intersections, mapper), mapper);
}

function xorWith(...values) {
    const lastValue = last(values);
    let comparator = (a, b) => a === b;
    if (typeof lastValue === 'function') {
        comparator = lastValue;
        values = values.slice(0, -1);
    }
    const arrays = values.filter(isArrayLikeObject);
    const union = unionWith(...arrays, comparator);
    const intersections = zip$1.windowed(arrays, 2).map(([arr1, arr2]) => intersectionWith(arr1, arr2, comparator));
    return differenceWith(union, unionWith(...intersections, comparator), comparator);
}

function zip(...arrays) {
    if (!arrays.length) {
        return [];
    }
    return zip$1.zip(...arrays.filter(group => isArrayLikeObject(group)));
}

const assignValue = (object, key, value) => {
    const objValue = object[key];
    if (!(Object.hasOwn(object, key) && isWeakSet$1.eq(objValue, value)) || (value === undefined && !(key in object))) {
        object[key] = value;
    }
};

function zipObject(keys = [], values = []) {
    const result = {};
    for (let i = 0; i < keys.length; i++) {
        assignValue(result, keys[i], values[i]);
    }
    return result;
}

function updateWith(obj, path, updater, customizer) {
    if (obj == null && !isObject(obj)) {
        return obj;
    }
    const resolvedPath = isKey(path, obj)
        ? [path]
        : Array.isArray(path)
            ? path
            : typeof path === 'string'
                ? toPath(path)
                : [path];
    let current = obj;
    for (let i = 0; i < resolvedPath.length && current != null; i++) {
        const key = toKey(resolvedPath[i]);
        let newValue;
        if (i === resolvedPath.length - 1) {
            newValue = updater(current[key]);
        }
        else {
            const objValue = current[key];
            const customizerResult = customizer(objValue);
            newValue =
                customizerResult !== undefined
                    ? customizerResult
                    : isObject(objValue)
                        ? objValue
                        : isIndex(resolvedPath[i + 1])
                            ? []
                            : {};
        }
        assignValue(current, key, newValue);
        current = current[key];
    }
    return obj;
}

function set(obj, path, value) {
    return updateWith(obj, path, () => value, () => undefined);
}

function zipObjectDeep(keys, values) {
    const result = {};
    if (!isArrayLike(keys)) {
        return result;
    }
    if (!isArrayLike(values)) {
        values = [];
    }
    const zipped = zip$1.zip(Array.from(keys), Array.from(values));
    for (let i = 0; i < zipped.length; i++) {
        const [key, value] = zipped[i];
        if (key != null) {
            set(result, key, value);
        }
    }
    return result;
}

function zipWith(...combine) {
    let iteratee = combine.pop();
    if (!isWeakSet$1.isFunction(iteratee)) {
        combine.push(iteratee);
        iteratee = undefined;
    }
    if (!combine?.length) {
        return [];
    }
    const result = unzip(combine);
    if (iteratee == null) {
        return result;
    }
    return result.map(group => iteratee(...group));
}

function after(n, func) {
    if (typeof func !== 'function') {
        throw new TypeError('Expected a function');
    }
    n = zip$1.toInteger(n);
    return function (...args) {
        if (--n < 1) {
            return func.apply(this, args);
        }
    };
}

function ary(func, n = func.length, guard) {
    if (guard) {
        n = func.length;
    }
    if (Number.isNaN(n) || n < 0) {
        n = 0;
    }
    return unary.ary(func, n);
}

function attempt(func, ...args) {
    try {
        return func(...args);
    }
    catch (e) {
        return e instanceof Error ? e : new Error(e);
    }
}

function before(n, func) {
    if (typeof func !== 'function') {
        throw new TypeError('Expected a function');
    }
    let result;
    n = zip$1.toInteger(n);
    return function (...args) {
        if (--n > 0) {
            result = func.apply(this, args);
        }
        if (n <= 1 && func) {
            func = undefined;
        }
        return result;
    };
}

function bind(func, thisObj, ...partialArgs) {
    const bound = function (...providedArgs) {
        const args = [];
        let startIndex = 0;
        for (let i = 0; i < partialArgs.length; i++) {
            const arg = partialArgs[i];
            if (arg === bind.placeholder) {
                args.push(providedArgs[startIndex++]);
            }
            else {
                args.push(arg);
            }
        }
        for (let i = startIndex; i < providedArgs.length; i++) {
            args.push(providedArgs[i]);
        }
        if (this instanceof bound) {
            return new func(...args);
        }
        return func.apply(thisObj, args);
    };
    return bound;
}
const bindPlaceholder = Symbol('bind.placeholder');
bind.placeholder = bindPlaceholder;

function bindKey(object, key, ...partialArgs) {
    const bound = function (...providedArgs) {
        const args = [];
        let startIndex = 0;
        for (let i = 0; i < partialArgs.length; i++) {
            const arg = partialArgs[i];
            if (arg === bindKey.placeholder) {
                args.push(providedArgs[startIndex++]);
            }
            else {
                args.push(arg);
            }
        }
        for (let i = startIndex; i < providedArgs.length; i++) {
            args.push(providedArgs[i]);
        }
        if (this instanceof bound) {
            return new object[key](...args);
        }
        return object[key].apply(object, args);
    };
    return bound;
}
const bindKeyPlaceholder = Symbol('bindKey.placeholder');
bindKey.placeholder = bindKeyPlaceholder;

function curry(func, arity = func.length, guard) {
    arity = guard ? func.length : arity;
    arity = Number.parseInt(arity, 10);
    if (Number.isNaN(arity) || arity < 1) {
        arity = 0;
    }
    const wrapper = function (...partialArgs) {
        const holders = partialArgs.filter(item => item === curry.placeholder);
        const length = partialArgs.length - holders.length;
        if (length < arity) {
            return makeCurry(func, arity - length, partialArgs);
        }
        if (this instanceof wrapper) {
            return new func(...partialArgs);
        }
        return func.apply(this, partialArgs);
    };
    wrapper.placeholder = curryPlaceholder;
    return wrapper;
}
function makeCurry(func, arity, partialArgs) {
    function wrapper(...providedArgs) {
        const holders = providedArgs.filter(item => item === curry.placeholder);
        const length = providedArgs.length - holders.length;
        providedArgs = composeArgs$1(providedArgs, partialArgs);
        if (length < arity) {
            return makeCurry(func, arity - length, providedArgs);
        }
        if (this instanceof wrapper) {
            return new func(...providedArgs);
        }
        return func.apply(this, providedArgs);
    }
    wrapper.placeholder = curryPlaceholder;
    return wrapper;
}
function composeArgs$1(providedArgs, partialArgs) {
    const args = [];
    let startIndex = 0;
    for (let i = 0; i < partialArgs.length; i++) {
        const arg = partialArgs[i];
        if (arg === curry.placeholder && startIndex < providedArgs.length) {
            args.push(providedArgs[startIndex++]);
        }
        else {
            args.push(arg);
        }
    }
    for (let i = startIndex; i < providedArgs.length; i++) {
        args.push(providedArgs[i]);
    }
    return args;
}
const curryPlaceholder = Symbol('curry.placeholder');
curry.placeholder = curryPlaceholder;

function curryRight(func, arity = func.length, guard) {
    arity = guard ? func.length : arity;
    arity = Number.parseInt(arity, 10);
    if (Number.isNaN(arity) || arity < 1) {
        arity = 0;
    }
    const wrapper = function (...partialArgs) {
        const holders = partialArgs.filter(item => item === curryRight.placeholder);
        const length = partialArgs.length - holders.length;
        if (length < arity) {
            return makeCurryRight(func, arity - length, partialArgs);
        }
        if (this instanceof wrapper) {
            return new func(...partialArgs);
        }
        return func.apply(this, partialArgs);
    };
    wrapper.placeholder = curryRightPlaceholder;
    return wrapper;
}
function makeCurryRight(func, arity, partialArgs) {
    function wrapper(...providedArgs) {
        const holders = providedArgs.filter(item => item === curryRight.placeholder);
        const length = providedArgs.length - holders.length;
        providedArgs = composeArgs(providedArgs, partialArgs);
        if (length < arity) {
            return makeCurryRight(func, arity - length, providedArgs);
        }
        if (this instanceof wrapper) {
            return new func(...providedArgs);
        }
        return func.apply(this, providedArgs);
    }
    wrapper.placeholder = curryRightPlaceholder;
    return wrapper;
}
function composeArgs(providedArgs, partialArgs) {
    const placeholderLength = partialArgs.filter(arg => arg === curryRight.placeholder).length;
    const rangeLength = Math.max(providedArgs.length - placeholderLength, 0);
    const args = [];
    let providedIndex = 0;
    for (let i = 0; i < rangeLength; i++) {
        args.push(providedArgs[providedIndex++]);
    }
    for (let i = 0; i < partialArgs.length; i++) {
        const arg = partialArgs[i];
        if (arg === curryRight.placeholder) {
            if (providedIndex < providedArgs.length) {
                args.push(providedArgs[providedIndex++]);
            }
            else {
                args.push(arg);
            }
        }
        else {
            args.push(arg);
        }
    }
    return args;
}
const curryRightPlaceholder = Symbol('curryRight.placeholder');
curryRight.placeholder = curryRightPlaceholder;

function debounce(func, debounceMs = 0, options = {}) {
    if (typeof options !== 'object') {
        options = {};
    }
    const { signal, leading = false, trailing = true, maxWait } = options;
    const edges = Array(2);
    if (leading) {
        edges[0] = 'leading';
    }
    if (trailing) {
        edges[1] = 'trailing';
    }
    let result = undefined;
    let pendingAt = null;
    const _debounced = unary.debounce(function (...args) {
        result = func.apply(this, args);
        pendingAt = null;
    }, debounceMs, { signal, edges });
    const debounced = function (...args) {
        if (maxWait != null) {
            if (pendingAt === null) {
                pendingAt = Date.now();
            }
            if (Date.now() - pendingAt >= maxWait) {
                result = func.apply(this, args);
                pendingAt = Date.now();
                _debounced.cancel();
                _debounced.schedule();
                return result;
            }
        }
        _debounced.apply(this, args);
        return result;
    };
    const flush = () => {
        _debounced.flush();
        return result;
    };
    debounced.cancel = _debounced.cancel;
    debounced.flush = flush;
    return debounced;
}

function defer(func, ...args) {
    if (typeof func !== 'function') {
        throw new TypeError('Expected a function');
    }
    return setTimeout(func, 1, ...args);
}

function delay(func, wait, ...args) {
    if (typeof func !== 'function') {
        throw new TypeError('Expected a function');
    }
    return setTimeout(func, zip$1.toNumber(wait) || 0, ...args);
}

function flip(func) {
    return function (...args) {
        return func.apply(this, args.reverse());
    };
}

function flow(...funcs) {
    const flattenFuncs = zip$1.flatten(funcs, 1);
    if (flattenFuncs.some(func => typeof func !== 'function')) {
        throw new TypeError('Expected a function');
    }
    return unary.flow(...flattenFuncs);
}

function flowRight(...funcs) {
    const flattenFuncs = zip$1.flatten(funcs, 1);
    if (flattenFuncs.some(func => typeof func !== 'function')) {
        throw new TypeError('Expected a function');
    }
    return unary.flowRight(...flattenFuncs);
}

function memoize(func, resolver) {
    if (typeof func !== 'function' || (resolver != null && typeof resolver !== 'function')) {
        throw new TypeError('Expected a function');
    }
    const memoized = function (...args) {
        const key = resolver ? resolver.apply(this, args) : args[0];
        const cache = memoized.cache;
        if (cache.has(key)) {
            return cache.get(key);
        }
        const result = func.apply(this, args);
        memoized.cache = cache.set(key, result) || cache;
        return result;
    };
    const CacheConstructor = memoize.Cache || Map;
    memoized.cache = new CacheConstructor();
    return memoized;
}
memoize.Cache = Map;

function nthArg(n = 0) {
    return function (...args) {
        return args.at(zip$1.toInteger(n));
    };
}

function partial(func, ...partialArgs) {
    return unary.partialImpl(func, partial.placeholder, ...partialArgs);
}
partial.placeholder = Symbol('compat.partial.placeholder');

function partialRight(func, ...partialArgs) {
    return unary.partialRightImpl(func, partialRight.placeholder, ...partialArgs);
}
partialRight.placeholder = Symbol('compat.partialRight.placeholder');

function rearg(func, ...indices) {
    const flattenIndices = flatten(indices);
    return function (...args) {
        const reorderedArgs = flattenIndices.map(i => args[i]).slice(0, args.length);
        for (let i = reorderedArgs.length; i < args.length; i++) {
            reorderedArgs.push(args[i]);
        }
        return func.apply(this, reorderedArgs);
    };
}

function rest(func, start = func.length - 1) {
    start = Number.parseInt(start, 10);
    if (Number.isNaN(start) || start < 0) {
        start = func.length - 1;
    }
    return unary.rest(func, start);
}

function spread(func, argsIndex = 0) {
    argsIndex = Number.parseInt(argsIndex, 10);
    if (Number.isNaN(argsIndex) || argsIndex < 0) {
        argsIndex = 0;
    }
    return function (...args) {
        const array = args[argsIndex];
        const params = args.slice(0, argsIndex);
        if (array) {
            params.push(...array);
        }
        return func.apply(this, params);
    };
}

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

function wrap(value, wrapper) {
    return function (...args) {
        const wrapFn = isWeakSet$1.isFunction(wrapper) ? wrapper : unary.identity;
        return wrapFn.apply(this, [value, ...args]);
    };
}

function toString(value) {
    if (value == null) {
        return '';
    }
    if (typeof value === 'string') {
        return value;
    }
    if (Array.isArray(value)) {
        return value.map(toString).join(',');
    }
    const result = String(value);
    if (result === '0' && Object.is(Number(value), -0)) {
        return '-0';
    }
    return result;
}

function add(value, other) {
    if (value === undefined && other === undefined) {
        return 0;
    }
    if (value === undefined || other === undefined) {
        return value ?? other;
    }
    if (typeof value === 'string' || typeof other === 'string') {
        value = toString(value);
        other = toString(other);
    }
    else {
        value = zip$1.toNumber(value);
        other = zip$1.toNumber(other);
    }
    return value + other;
}

function decimalAdjust(type, number, precision = 0) {
    number = Number(number);
    if (Object.is(number, -0)) {
        number = '-0';
    }
    precision = Math.min(Number.parseInt(precision, 10), 292);
    if (precision) {
        const [magnitude, exponent = 0] = number.toString().split('e');
        let adjustedValue = Math[type](Number(`${magnitude}e${Number(exponent) + precision}`));
        if (Object.is(adjustedValue, -0)) {
            adjustedValue = '-0';
        }
        const [newMagnitude, newExponent = 0] = adjustedValue.toString().split('e');
        return Number(`${newMagnitude}e${Number(newExponent) - precision}`);
    }
    return Math[type](Number(number));
}

function ceil(number, precision = 0) {
    return decimalAdjust('ceil', number, precision);
}

function divide(value, other) {
    if (value === undefined && other === undefined) {
        return 1;
    }
    if (value === undefined || other === undefined) {
        return value ?? other;
    }
    if (typeof value === 'string' || typeof other === 'string') {
        value = toString(value);
        other = toString(other);
    }
    else {
        value = zip$1.toNumber(value);
        other = zip$1.toNumber(other);
    }
    return value / other;
}

function floor(number, precision = 0) {
    return decimalAdjust('floor', number, precision);
}

function inRange(value, minimum, maximum) {
    if (!minimum) {
        minimum = 0;
    }
    if (maximum != null && !maximum) {
        maximum = 0;
    }
    if (minimum != null && typeof minimum !== 'number') {
        minimum = Number(minimum);
    }
    if (maximum == null && minimum === 0) {
        return false;
    }
    if (maximum != null && typeof maximum !== 'number') {
        maximum = Number(maximum);
    }
    if (maximum != null && minimum > maximum) {
        [minimum, maximum] = [maximum, minimum];
    }
    if (minimum === maximum) {
        return false;
    }
    return range$1.inRange(value, minimum, maximum);
}

function max(items) {
    if (!items || items.length === 0) {
        return undefined;
    }
    let maxResult = undefined;
    for (let i = 0; i < items.length; i++) {
        const current = items[i];
        if (current == null || Number.isNaN(current) || typeof current === 'symbol') {
            continue;
        }
        if (maxResult === undefined || current > maxResult) {
            maxResult = current;
        }
    }
    return maxResult;
}

function maxBy(items, iteratee$1) {
    if (items == null) {
        return undefined;
    }
    return zip$1.maxBy(Array.from(items), iteratee(iteratee$1));
}

function sumBy(array, iteratee$1) {
    if (!array || !array.length) {
        return 0;
    }
    if (iteratee$1 != null) {
        iteratee$1 = iteratee(iteratee$1);
    }
    let result = undefined;
    for (let i = 0; i < array.length; i++) {
        const current = iteratee$1 ? iteratee$1(array[i]) : array[i];
        if (current !== undefined) {
            if (result === undefined) {
                result = current;
            }
            else {
                result += current;
            }
        }
    }
    return result;
}

function sum(array) {
    return sumBy(array);
}

function mean(nums) {
    const length = nums ? nums.length : 0;
    return length === 0 ? NaN : sum(nums) / length;
}

function meanBy(items, iteratee$1) {
    if (items == null) {
        return NaN;
    }
    return range$1.meanBy(Array.from(items), iteratee(iteratee$1));
}

function min(items) {
    if (!items || items.length === 0) {
        return undefined;
    }
    let minResult = undefined;
    for (let i = 0; i < items.length; i++) {
        const current = items[i];
        if (current == null || Number.isNaN(current) || typeof current === 'symbol') {
            continue;
        }
        if (minResult === undefined || current < minResult) {
            minResult = current;
        }
    }
    return minResult;
}

function minBy(items, iteratee$1) {
    if (items == null) {
        return undefined;
    }
    return zip$1.minBy(Array.from(items), iteratee(iteratee$1));
}

function multiply(value, other) {
    if (value === undefined && other === undefined) {
        return 1;
    }
    if (value === undefined || other === undefined) {
        return value ?? other;
    }
    if (typeof value === 'string' || typeof other === 'string') {
        value = toString(value);
        other = toString(other);
    }
    else {
        value = zip$1.toNumber(value);
        other = zip$1.toNumber(other);
    }
    return value * other;
}

function parseInt(string, radix = 0, guard) {
    if (guard) {
        radix = 0;
    }
    return Number.parseInt(string, radix);
}

function random(...args) {
    let minimum = 0;
    let maximum = 1;
    let floating = false;
    switch (args.length) {
        case 1: {
            if (typeof args[0] === 'boolean') {
                floating = args[0];
            }
            else {
                maximum = args[0];
            }
            break;
        }
        case 2: {
            if (typeof args[1] === 'boolean') {
                maximum = args[0];
                floating = args[1];
            }
            else {
                minimum = args[0];
                maximum = args[1];
            }
        }
        case 3: {
            if (typeof args[2] === 'object' && args[2] != null && args[2][args[1]] === args[0]) {
                minimum = 0;
                maximum = args[0];
                floating = false;
            }
            else {
                minimum = args[0];
                maximum = args[1];
                floating = args[2];
            }
        }
    }
    if (typeof minimum !== 'number') {
        minimum = Number(minimum);
    }
    if (typeof maximum !== 'number') {
        minimum = Number(maximum);
    }
    if (!minimum) {
        minimum = 0;
    }
    if (!maximum) {
        maximum = 0;
    }
    if (minimum > maximum) {
        [minimum, maximum] = [maximum, minimum];
    }
    minimum = clamp(minimum, -Number.MAX_SAFE_INTEGER, Number.MAX_SAFE_INTEGER);
    maximum = clamp(maximum, -Number.MAX_SAFE_INTEGER, Number.MAX_SAFE_INTEGER);
    if (minimum === maximum) {
        return minimum;
    }
    if (floating) {
        return randomInt.random(minimum, maximum + 1);
    }
    else {
        return randomInt.randomInt(minimum, maximum + 1);
    }
}

function range(start, end, step) {
    if (step && typeof step !== 'number' && isIterateeCall(start, end, step)) {
        end = step = undefined;
    }
    start = zip$1.toFinite(start);
    if (end === undefined) {
        end = start;
        start = 0;
    }
    else {
        end = zip$1.toFinite(end);
    }
    step = step === undefined ? (start < end ? 1 : -1) : zip$1.toFinite(step);
    const length = Math.max(Math.ceil((end - start) / (step || 1)), 0);
    const result = new Array(length);
    for (let index = 0; index < length; index++) {
        result[index] = start;
        start += step;
    }
    return result;
}

function rangeRight(start, end, step) {
    if (step && typeof step !== 'number' && isIterateeCall(start, end, step)) {
        end = step = undefined;
    }
    start = zip$1.toFinite(start);
    if (end === undefined) {
        end = start;
        start = 0;
    }
    else {
        end = zip$1.toFinite(end);
    }
    step = step === undefined ? (start < end ? 1 : -1) : zip$1.toFinite(step);
    const length = Math.max(Math.ceil((end - start) / (step || 1)), 0);
    const result = new Array(length);
    for (let index = length - 1; index >= 0; index--) {
        result[index] = start;
        start += step;
    }
    return result;
}

function round(number, precision = 0) {
    return decimalAdjust('round', number, precision);
}

function subtract(value, other) {
    if (value === undefined && other === undefined) {
        return 0;
    }
    if (value === undefined || other === undefined) {
        return value ?? other;
    }
    if (typeof value === 'string' || typeof other === 'string') {
        value = toString(value);
        other = toString(other);
    }
    else {
        value = zip$1.toNumber(value);
        other = zip$1.toNumber(other);
    }
    return value - other;
}

function isPrototype(value) {
    const constructor = value?.constructor;
    const prototype = typeof constructor === 'function' ? constructor.prototype : Object.prototype;
    return value === prototype;
}

function isTypedArray(x) {
    return isPlainObject$1.isTypedArray(x);
}

function times(n, getValue) {
    n = zip$1.toInteger(n);
    if (n < 1 || !Number.isSafeInteger(n)) {
        return [];
    }
    const result = new Array(n);
    for (let i = 0; i < n; i++) {
        result[i] = typeof getValue === 'function' ? getValue(i) : i;
    }
    return result;
}

function keys(object) {
    if (isArrayLike(object)) {
        return arrayLikeKeys(object);
    }
    const result = Object.keys(Object(object));
    if (!isPrototype(object)) {
        return result;
    }
    return result.filter(key => key !== 'constructor');
}
function arrayLikeKeys(object) {
    const indices = times(object.length, index => `${index}`);
    const filteredKeys = new Set(indices);
    if (isWeakSet$1.isBuffer(object)) {
        filteredKeys.add('offset');
        filteredKeys.add('parent');
    }
    if (isTypedArray(object)) {
        filteredKeys.add('buffer');
        filteredKeys.add('byteLength');
        filteredKeys.add('byteOffset');
    }
    return [...indices, ...Object.keys(object).filter(key => !filteredKeys.has(key))];
}

function assign(object, ...sources) {
    for (let i = 0; i < sources.length; i++) {
        assignImpl(object, sources[i]);
    }
    return object;
}
function assignImpl(object, source) {
    const keys$1 = keys(source);
    for (let i = 0; i < keys$1.length; i++) {
        const key = keys$1[i];
        if (!(key in object) || !isWeakSet$1.eq(object[key], source[key])) {
            object[key] = source[key];
        }
    }
}

function keysIn(object) {
    if (object == null) {
        return [];
    }
    switch (typeof object) {
        case 'object':
        case 'function': {
            if (isArrayLike(object)) {
                return arrayLikeKeysIn(object);
            }
            if (isPrototype(object)) {
                return prototypeKeysIn(object);
            }
            return keysInImpl(object);
        }
        default: {
            return keysInImpl(Object(object));
        }
    }
}
function keysInImpl(object) {
    const result = [];
    for (const key in object) {
        result.push(key);
    }
    return result;
}
function prototypeKeysIn(object) {
    const keys = keysInImpl(object);
    return keys.filter(key => key !== 'constructor');
}
function arrayLikeKeysIn(object) {
    const indices = times(object.length, index => `${index}`);
    const filteredKeys = new Set(indices);
    if (isWeakSet$1.isBuffer(object)) {
        filteredKeys.add('offset');
        filteredKeys.add('parent');
    }
    if (isTypedArray(object)) {
        filteredKeys.add('buffer');
        filteredKeys.add('byteLength');
        filteredKeys.add('byteOffset');
    }
    return [...indices, ...keysInImpl(object).filter(key => !filteredKeys.has(key))];
}

function assignIn(object, ...sources) {
    for (let i = 0; i < sources.length; i++) {
        assignInImpl(object, sources[i]);
    }
    return object;
}
function assignInImpl(object, source) {
    const keys = keysIn(source);
    for (let i = 0; i < keys.length; i++) {
        const key = keys[i];
        if (!(key in object) || !isWeakSet$1.eq(object[key], source[key])) {
            object[key] = source[key];
        }
    }
}

function assignInWith(object, ...sources) {
    let getValueToAssign = sources[sources.length - 1];
    if (typeof getValueToAssign === 'function') {
        sources.pop();
    }
    else {
        getValueToAssign = undefined;
    }
    for (let i = 0; i < sources.length; i++) {
        assignInWithImpl(object, sources[i], getValueToAssign);
    }
    return object;
}
function assignInWithImpl(object, source, getValueToAssign) {
    const keys = keysIn(source);
    for (let i = 0; i < keys.length; i++) {
        const key = keys[i];
        const objValue = object[key];
        const srcValue = source[key];
        const newValue = getValueToAssign?.(objValue, srcValue, key, object, source) ?? srcValue;
        if (!(key in object) || !isWeakSet$1.eq(objValue, newValue)) {
            object[key] = newValue;
        }
    }
}

function assignWith(object, ...sources) {
    let getValueToAssign = sources[sources.length - 1];
    if (typeof getValueToAssign === 'function') {
        sources.pop();
    }
    else {
        getValueToAssign = undefined;
    }
    for (let i = 0; i < sources.length; i++) {
        assignWithImpl(object, sources[i], getValueToAssign);
    }
    return object;
}
function assignWithImpl(object, source, getValueToAssign) {
    const keys$1 = keys(source);
    for (let i = 0; i < keys$1.length; i++) {
        const key = keys$1[i];
        const objValue = object[key];
        const srcValue = source[key];
        const newValue = getValueToAssign?.(objValue, srcValue, key, object, source) ?? srcValue;
        if (!(key in object) || !isWeakSet$1.eq(objValue, newValue)) {
            object[key] = newValue;
        }
    }
}

function clone(obj) {
    if (isPlainObject$1.isPrimitive(obj)) {
        return obj;
    }
    const tag = isPlainObject$1.getTag(obj);
    if (!isCloneableObject(obj)) {
        return {};
    }
    if (isPlainObject.isArray(obj)) {
        const result = Array.from(obj);
        if (obj.length > 0 && typeof obj[0] === 'string' && Object.hasOwn(obj, 'index')) {
            result.index = obj.index;
            result.input = obj.input;
        }
        return result;
    }
    if (isTypedArray(obj)) {
        const typedArray = obj;
        const Ctor = typedArray.constructor;
        return new Ctor(typedArray.buffer, typedArray.byteOffset, typedArray.length);
    }
    if (tag === isPlainObject$1.arrayBufferTag) {
        return new ArrayBuffer(obj.byteLength);
    }
    if (tag === isPlainObject$1.dataViewTag) {
        const dataView = obj;
        const buffer = dataView.buffer;
        const byteOffset = dataView.byteOffset;
        const byteLength = dataView.byteLength;
        const clonedBuffer = new ArrayBuffer(byteLength);
        const srcView = new Uint8Array(buffer, byteOffset, byteLength);
        const destView = new Uint8Array(clonedBuffer);
        destView.set(srcView);
        return new DataView(clonedBuffer);
    }
    if (tag === isPlainObject$1.booleanTag || tag === isPlainObject$1.numberTag || tag === isPlainObject$1.stringTag) {
        const Ctor = obj.constructor;
        const clone = new Ctor(obj.valueOf());
        if (tag === isPlainObject$1.stringTag) {
            cloneStringObjectProperties(clone, obj);
        }
        else {
            copyOwnProperties(clone, obj);
        }
        return clone;
    }
    if (tag === isPlainObject$1.dateTag) {
        return new Date(Number(obj));
    }
    if (tag === isPlainObject$1.regexpTag) {
        const regExp = obj;
        const clone = new RegExp(regExp.source, regExp.flags);
        clone.lastIndex = regExp.lastIndex;
        return clone;
    }
    if (tag === isPlainObject$1.symbolTag) {
        return Object(Symbol.prototype.valueOf.call(obj));
    }
    if (tag === isPlainObject$1.mapTag) {
        const map = obj;
        const result = new Map();
        map.forEach((obj, key) => {
            result.set(key, obj);
        });
        return result;
    }
    if (tag === isPlainObject$1.setTag) {
        const set = obj;
        const result = new Set();
        set.forEach(obj => {
            result.add(obj);
        });
        return result;
    }
    if (tag === isPlainObject$1.argumentsTag) {
        const args = obj;
        const result = {};
        copyOwnProperties(result, args);
        result.length = args.length;
        result[Symbol.iterator] = args[Symbol.iterator];
        return result;
    }
    const result = {};
    copyPrototype(result, obj);
    copyOwnProperties(result, obj);
    copySymbolProperties(result, obj);
    return result;
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
function copyOwnProperties(target, source) {
    for (const key in source) {
        if (Object.hasOwn(source, key)) {
            target[key] = source[key];
        }
    }
}
function copySymbolProperties(target, source) {
    const symbols = Object.getOwnPropertySymbols(source);
    for (let i = 0; i < symbols.length; i++) {
        const symbol = symbols[i];
        if (Object.prototype.propertyIsEnumerable.call(source, symbol)) {
            target[symbol] = source[symbol];
        }
    }
}
function cloneStringObjectProperties(target, source) {
    const stringLength = source.valueOf().length;
    for (const key in source) {
        if (Object.hasOwn(source, key) && (Number.isNaN(Number(key)) || Number(key) >= stringLength)) {
            target[key] = source[key];
        }
    }
}
function copyPrototype(target, source) {
    const proto = Object.getPrototypeOf(source);
    if (proto !== null) {
        const Ctor = source.constructor;
        if (typeof Ctor === 'function') {
            Object.setPrototypeOf(target, proto);
        }
    }
}

function cloneWith(value, customizer) {
    if (!customizer) {
        return clone(value);
    }
    const result = customizer(value);
    if (result !== undefined) {
        return result;
    }
    return clone(value);
}

function create(prototype, properties) {
    const proto = isObject(prototype) ? Object.create(prototype) : {};
    if (properties != null) {
        const propsKeys = keys(properties);
        for (let i = 0; i < propsKeys.length; i++) {
            const key = propsKeys[i];
            const propsValue = properties[key];
            assignValue(proto, key, propsValue);
        }
    }
    return proto;
}

function defaults(object, ...sources) {
    object = Object(object);
    const objectProto = Object.prototype;
    let length = sources.length;
    const guard = length > 2 ? sources[2] : undefined;
    if (guard && isIterateeCall(sources[0], sources[1], guard)) {
        length = 1;
    }
    for (let i = 0; i < length; i++) {
        const source = sources[i];
        const keys = Object.keys(source);
        for (let j = 0; j < keys.length; j++) {
            const key = keys[j];
            const value = object[key];
            if (value === undefined ||
                (!Object.hasOwn(object, key) && isWeakSet$1.eq(value, objectProto[key]))) {
                object[key] = source[key];
            }
        }
    }
    return object;
}

function findKey(obj, predicate) {
    if (!isObject(obj)) {
        return undefined;
    }
    return findKeyImpl(obj, predicate);
}
function findKeyImpl(obj, predicate) {
    if (typeof predicate === 'function') {
        return isPlainObject.findKey(obj, predicate);
    }
    if (typeof predicate === 'object') {
        if (Array.isArray(predicate)) {
            const key = predicate[0];
            const value = predicate[1];
            return isPlainObject.findKey(obj, matchesProperty(key, value));
        }
        return isPlainObject.findKey(obj, matches(predicate));
    }
    if (typeof predicate === 'string') {
        return isPlainObject.findKey(obj, property(predicate));
    }
}

function forIn(object, iteratee = unary.identity) {
    if (object == null) {
        return object;
    }
    for (const key in object) {
        const result = iteratee(object[key], key, object);
        if (result === false) {
            break;
        }
    }
    return object;
}

function forInRight(object, iteratee = unary.identity) {
    if (object == null) {
        return object;
    }
    const keys = [];
    for (const key in object) {
        keys.push(key);
    }
    for (let i = keys.length - 1; i >= 0; i--) {
        const key = keys[i];
        const result = iteratee(object[key], key, object);
        if (result === false) {
            break;
        }
    }
    return object;
}

function forOwn(object, iteratee = unary.identity) {
    if (object == null) {
        return object;
    }
    const iterable = Object(object);
    const keys$1 = keys(object);
    for (let i = 0; i < keys$1.length; ++i) {
        const key = keys$1[i];
        if (iteratee(iterable[key], key, iterable) === false) {
            break;
        }
    }
    return object;
}

function forOwnRight(object, iteratee = unary.identity) {
    if (object == null) {
        return object;
    }
    const iterable = Object(object);
    const keys$1 = keys(object);
    for (let i = keys$1.length - 1; i >= 0; --i) {
        const key = keys$1[i];
        if (iteratee(iterable[key], key, iterable) === false) {
            break;
        }
    }
    return object;
}

function fromPairs(pairs) {
    if (!isArrayLike(pairs) && !(pairs instanceof Map)) {
        return {};
    }
    const result = {};
    for (const [key, value] of pairs) {
        result[key] = value;
    }
    return result;
}

function functions(object) {
    if (object == null) {
        return [];
    }
    return keys(object).filter(key => typeof object[key] === 'function');
}

function functionsIn(object) {
    if (object == null) {
        return [];
    }
    const result = [];
    for (const key in object) {
        if (isWeakSet$1.isFunction(object[key])) {
            result.push(key);
        }
    }
    return result;
}

function hasIn(object, path) {
    let resolvedPath;
    if (Array.isArray(path)) {
        resolvedPath = path;
    }
    else if (typeof path === 'string' && isDeepKey(path) && object?.[path] == null) {
        resolvedPath = toPath(path);
    }
    else {
        resolvedPath = [path];
    }
    if (resolvedPath.length === 0) {
        return false;
    }
    let current = object;
    for (let i = 0; i < resolvedPath.length; i++) {
        const key = resolvedPath[i];
        if (current == null || !(key in Object(current))) {
            const isSparseIndex = (Array.isArray(current) || isArguments(current)) && isIndex(key) && key < current.length;
            if (!isSparseIndex) {
                return false;
            }
        }
        current = current[key];
    }
    return true;
}

function invertBy(object, iteratee) {
    const result = {};
    if (isWeakSet$1.isNil(object)) {
        return result;
    }
    if (iteratee == null) {
        iteratee = unary.identity;
    }
    const keys = Object.keys(object);
    for (let i = 0; i < keys.length; i++) {
        const key = keys[i];
        const value = object[key];
        const valueStr = iteratee(value);
        if (Array.isArray(result[valueStr])) {
            result[valueStr].push(key);
        }
        else {
            result[valueStr] = [key];
        }
    }
    return result;
}

const functionToString = Function.prototype.toString;
const REGEXP_SYNTAX_CHARS = /[\\^$.*+?()[\]{}|]/g;
const IS_NATIVE_FUNCTION_REGEXP = RegExp(`^${functionToString
    .call(Object.prototype.hasOwnProperty)
    .replace(REGEXP_SYNTAX_CHARS, '\\$&')
    .replace(/hasOwnProperty|(function).*?(?=\\\()| for .+?(?=\\\])/g, '$1.*?')}$`);
function isNative(value) {
    if (typeof value !== 'function') {
        return false;
    }
    if (globalThis?.['__core-js_shared__'] != null) {
        throw new Error('Unsupported core-js use. Try https://npms.io/search?q=ponyfill.');
    }
    return IS_NATIVE_FUNCTION_REGEXP.test(functionToString.call(value));
}

function mapKeys(object, getNewKey) {
    getNewKey = getNewKey ?? unary.identity;
    switch (typeof getNewKey) {
        case 'string':
        case 'symbol':
        case 'number':
        case 'object': {
            return isPlainObject.mapKeys(object, property(getNewKey));
        }
        case 'function': {
            return isPlainObject.mapKeys(object, getNewKey);
        }
    }
}

function mapValues(object, getNewValue) {
    getNewValue = getNewValue ?? unary.identity;
    switch (typeof getNewValue) {
        case 'string':
        case 'symbol':
        case 'number':
        case 'object': {
            return isPlainObject.mapValues(object, property(getNewValue));
        }
        case 'function': {
            return isPlainObject.mapValues(object, getNewValue);
        }
    }
}

function mergeWith(object, ...otherArgs) {
    const sources = otherArgs.slice(0, -1);
    const merge = otherArgs[otherArgs.length - 1];
    let result = object;
    for (let i = 0; i < sources.length; i++) {
        const source = sources[i];
        result = mergeWithDeep(result, source, merge, new Map());
    }
    return result;
}
function mergeWithDeep(target, source, merge, stack) {
    if (isPlainObject$1.isPrimitive(target)) {
        target = Object(target);
    }
    if (source == null || typeof source !== 'object') {
        return target;
    }
    if (stack.has(source)) {
        return isPlainObject.clone(stack.get(source));
    }
    stack.set(source, target);
    if (Array.isArray(source)) {
        source = source.slice();
        for (let i = 0; i < source.length; i++) {
            source[i] = source[i] ?? undefined;
        }
    }
    const sourceKeys = [...Object.keys(source), ...isPlainObject$1.getSymbols(source)];
    for (let i = 0; i < sourceKeys.length; i++) {
        const key = sourceKeys[i];
        let sourceValue = source[key];
        let targetValue = target[key];
        if (isArguments(sourceValue)) {
            sourceValue = { ...sourceValue };
        }
        if (isArguments(targetValue)) {
            targetValue = { ...targetValue };
        }
        if (typeof Buffer !== 'undefined' && Buffer.isBuffer(sourceValue)) {
            sourceValue = cloneDeep(sourceValue);
        }
        if (Array.isArray(sourceValue)) {
            if (typeof targetValue === 'object' && targetValue != null) {
                const cloned = [];
                const targetKeys = Reflect.ownKeys(targetValue);
                for (let i = 0; i < targetKeys.length; i++) {
                    const targetKey = targetKeys[i];
                    cloned[targetKey] = targetValue[targetKey];
                }
                targetValue = cloned;
            }
            else {
                targetValue = [];
            }
        }
        const merged = merge(targetValue, sourceValue, key, target, source, stack);
        if (merged != null) {
            target[key] = merged;
        }
        else if (Array.isArray(sourceValue)) {
            target[key] = mergeWithDeep(targetValue, sourceValue, merge, stack);
        }
        else if (isPlainObject.isObjectLike(targetValue) && isPlainObject.isObjectLike(sourceValue)) {
            target[key] = mergeWithDeep(targetValue, sourceValue, merge, stack);
        }
        else if (targetValue == null && isPlainObject.isPlainObject(sourceValue)) {
            target[key] = mergeWithDeep({}, sourceValue, merge, stack);
        }
        else if (targetValue == null && isTypedArray(sourceValue)) {
            target[key] = cloneDeep(sourceValue);
        }
        else if (targetValue === undefined || sourceValue !== undefined) {
            target[key] = sourceValue;
        }
    }
    return target;
}

function merge(object, ...sources) {
    return mergeWith(object, ...sources, noop.noop);
}

function omit(obj, ...keysArr) {
    if (obj == null) {
        return {};
    }
    const result = isPlainObject.cloneDeep(obj);
    for (let i = 0; i < keysArr.length; i++) {
        let keys = keysArr[i];
        switch (typeof keys) {
            case 'object': {
                if (!Array.isArray(keys)) {
                    keys = Array.from(keys);
                }
                for (let j = 0; j < keys.length; j++) {
                    const key = keys[j];
                    unset(result, key);
                }
                break;
            }
            case 'string':
            case 'symbol':
            case 'number': {
                unset(result, keys);
                break;
            }
        }
    }
    return result;
}

function getSymbolsIn(object) {
    const result = [];
    while (object) {
        result.push(...isPlainObject$1.getSymbols(object));
        object = Object.getPrototypeOf(object);
    }
    return result;
}

function omitBy(obj, shouldOmit) {
    if (obj == null) {
        return {};
    }
    const result = {};
    if (shouldOmit == null) {
        return {};
    }
    const keys = isArrayLike(obj) ? range$1.range(0, obj.length) : [...keysIn(obj), ...getSymbolsIn(obj)];
    for (let i = 0; i < keys.length; i++) {
        const key = (zip$1.isSymbol(keys[i]) ? keys[i] : keys[i].toString());
        const value = obj[key];
        if (!shouldOmit(value, key, obj)) {
            result[key] = value;
        }
    }
    return result;
}

function pick(obj, ...keysArr) {
    if (isNil(obj)) {
        return {};
    }
    const result = {};
    for (let i = 0; i < keysArr.length; i++) {
        let keys = keysArr[i];
        switch (typeof keys) {
            case 'object': {
                if (!Array.isArray(keys)) {
                    if (isArrayLike(keys)) {
                        keys = Array.from(keys);
                    }
                    else {
                        keys = [keys];
                    }
                }
                break;
            }
            case 'string':
            case 'symbol':
            case 'number': {
                keys = [keys];
                break;
            }
        }
        for (const key of keys) {
            const value = get(obj, key);
            if (value === undefined && !has(obj, key)) {
                continue;
            }
            if (typeof key === 'string' && Object.hasOwn(obj, key)) {
                result[key] = value;
            }
            else {
                set(result, key, value);
            }
        }
    }
    return result;
}

function pickBy(obj, shouldPick) {
    if (obj == null) {
        return {};
    }
    const result = {};
    if (shouldPick == null) {
        return obj;
    }
    const keys = isArrayLike(obj) ? range$1.range(0, obj.length) : [...keysIn(obj), ...getSymbolsIn(obj)];
    for (let i = 0; i < keys.length; i++) {
        const key = (zip$1.isSymbol(keys[i]) ? keys[i] : keys[i].toString());
        const value = obj[key];
        if (shouldPick(value, key, obj)) {
            result[key] = value;
        }
    }
    return result;
}

function propertyOf(object) {
    return function (path) {
        return get(object, path);
    };
}

function result(object, path, defaultValue) {
    if (isKey(path, object)) {
        path = [path];
    }
    else if (!Array.isArray(path)) {
        path = toPath(toString(path));
    }
    const pathLength = Math.max(path.length, 1);
    for (let index = 0; index < pathLength; index++) {
        const value = object == null ? undefined : object[toKey(path[index])];
        if (value === undefined) {
            return typeof defaultValue === 'function' ? defaultValue.call(object) : defaultValue;
        }
        object = typeof value === 'function' ? value.call(object) : value;
    }
    return object;
}

function setWith(obj, path, value, customizer) {
    let customizerFn;
    if (typeof customizer === 'function') {
        customizerFn = customizer;
    }
    else {
        customizerFn = () => undefined;
    }
    return updateWith(obj, path, () => value, customizerFn);
}

function toDefaulted(object, ...sources) {
    const cloned = cloneDeep(object);
    return defaults(cloned, ...sources);
}

function mapToEntries(map) {
    const arr = new Array(map.size);
    const keys = map.keys();
    const values = map.values();
    for (let i = 0; i < arr.length; i++) {
        arr[i] = [keys.next().value, values.next().value];
    }
    return arr;
}

function setToEntries(set) {
    const arr = new Array(set.size);
    const values = set.values();
    for (let i = 0; i < arr.length; i++) {
        const value = values.next().value;
        arr[i] = [value, value];
    }
    return arr;
}

function toPairs(object) {
    if (object instanceof Set) {
        return setToEntries(object);
    }
    if (object instanceof Map) {
        return mapToEntries(object);
    }
    const keys$1 = keys(object);
    const result = new Array(keys$1.length);
    for (let i = 0; i < keys$1.length; i++) {
        const key = keys$1[i];
        const value = object[key];
        result[i] = [key, value];
    }
    return result;
}

function toPairsIn(object) {
    if (object instanceof Set) {
        return setToEntries(object);
    }
    if (object instanceof Map) {
        return mapToEntries(object);
    }
    const keys = keysIn(object);
    const result = new Array(keys.length);
    for (let i = 0; i < keys.length; i++) {
        const key = keys[i];
        const value = object[key];
        result[i] = [key, value];
    }
    return result;
}

function isBuffer(x) {
    return isWeakSet$1.isBuffer(x);
}

function transform(object, iteratee$1 = unary.identity, accumulator) {
    const isArrayOrBufferOrTypedArray = Array.isArray(object) || isBuffer(object) || isTypedArray(object);
    iteratee$1 = iteratee(iteratee$1);
    if (accumulator == null) {
        if (isArrayOrBufferOrTypedArray) {
            accumulator = [];
        }
        else if (isObject(object) && isWeakSet$1.isFunction(object.constructor)) {
            accumulator = Object.create(Object.getPrototypeOf(object));
        }
        else {
            accumulator = {};
        }
    }
    if (object == null) {
        return accumulator;
    }
    forEach(object, (value, key, object) => iteratee$1(accumulator, value, key, object));
    return accumulator;
}

function update(obj, path, updater) {
    return updateWith(obj, path, updater, () => undefined);
}

function valuesIn(object) {
    const keys = keysIn(object);
    const result = new Array(keys.length);
    for (let i = 0; i < keys.length; i++) {
        const key = keys[i];
        result[i] = object[key];
    }
    return result;
}

function bindAll(object, ...methodNames) {
    if (object == null) {
        return object;
    }
    if (!isObject(object)) {
        return object;
    }
    if (isPlainObject.isArray(object) && methodNames.length === 0) {
        return object;
    }
    const methods = [];
    for (let i = 0; i < methodNames.length; i++) {
        const name = methodNames[i];
        if (isPlainObject.isArray(name)) {
            methods.push(...name);
        }
        else if (name && typeof name === 'object' && 'length' in name) {
            methods.push(...Array.from(name));
        }
        else {
            methods.push(name);
        }
    }
    if (methods.length === 0) {
        return object;
    }
    for (let i = 0; i < methods.length; i++) {
        const key = methods[i];
        const stringKey = toString(key);
        const func = object[stringKey];
        if (isWeakSet$1.isFunction(func)) {
            object[stringKey] = func.bind(object);
        }
    }
    return object;
}

function conformsTo(target, source) {
    if (source == null) {
        return true;
    }
    if (target == null) {
        return Object.keys(source).length === 0;
    }
    const keys = Object.keys(source);
    for (let i = 0; i < keys.length; i++) {
        const key = keys[i];
        const predicate = source[key];
        const value = target[key];
        if ((value === undefined && !(key in target)) || !predicate(value)) {
            return false;
        }
    }
    return true;
}

function conforms(source) {
    source = isPlainObject.cloneDeep(source);
    return function (object) {
        return conformsTo(object, source);
    };
}

function isArrayBuffer(value) {
    return isWeakSet$1.isArrayBuffer(value);
}

function isBoolean(value) {
    return typeof value === 'boolean' || value instanceof Boolean;
}

function isDate(value) {
    return isWeakSet$1.isDate(value);
}

function isElement(value) {
    return isPlainObject.isObjectLike(value) && value.nodeType === 1 && !isPlainObject.isPlainObject(value);
}

function isEmpty(value) {
    if (value == null) {
        return true;
    }
    if (isArrayLike(value)) {
        if (typeof value.splice !== 'function' &&
            typeof value !== 'string' &&
            (typeof Buffer === 'undefined' || !Buffer.isBuffer(value)) &&
            !isTypedArray(value) &&
            !isArguments(value)) {
            return false;
        }
        return value.length === 0;
    }
    if (typeof value === 'object') {
        if (value instanceof Map || value instanceof Set) {
            return value.size === 0;
        }
        const keys = Object.keys(value);
        if (isPrototype(value)) {
            return keys.filter(x => x !== 'constructor').length === 0;
        }
        return keys.length === 0;
    }
    return true;
}

function isEqualWith(a, b, areValuesEqual = noop.noop) {
    if (typeof areValuesEqual !== 'function') {
        areValuesEqual = noop.noop;
    }
    return isWeakSet$1.isEqualWith(a, b, (...args) => {
        const result = areValuesEqual(...args);
        if (result !== undefined) {
            return Boolean(result);
        }
        if (a instanceof Map && b instanceof Map) {
            return isEqualWith(Array.from(a), Array.from(b), unary.after(2, areValuesEqual));
        }
        if (a instanceof Set && b instanceof Set) {
            return isEqualWith(Array.from(a), Array.from(b), unary.after(2, areValuesEqual));
        }
    });
}

function isError(value) {
    return isPlainObject$1.getTag(value) === '[object Error]';
}

function isFinite(value) {
    return Number.isFinite(value);
}

function isInteger(value) {
    return Number.isInteger(value);
}

function isRegExp(value) {
    return isWeakSet$1.isRegExp(value);
}

function isSafeInteger(value) {
    return Number.isSafeInteger(value);
}

function isSet(value) {
    return isWeakSet$1.isSet(value);
}

function isWeakMap(value) {
    return isWeakSet$1.isWeakMap(value);
}

function isWeakSet(value) {
    return isWeakSet$1.isWeakSet(value);
}

function normalizeForCase(str) {
    if (typeof str !== 'string') {
        str = toString(str);
    }
    return str.replace(/['\u2019]/g, '');
}

function camelCase(str) {
    return snakeCase$1.camelCase(normalizeForCase(str));
}

function deburr(str) {
    return upperFirst$1.deburr(toString(str));
}

function endsWith(str, target, position = str.length) {
    return str.endsWith(target, position);
}

function escape(string) {
    return upperFirst$1.escape(toString(string));
}

function escapeRegExp(str) {
    return upperFirst$1.escapeRegExp(toString(str));
}

function kebabCase(str) {
    return upperFirst$1.kebabCase(normalizeForCase(str));
}

function lowerCase(str) {
    return upperFirst$1.lowerCase(normalizeForCase(str));
}

function lowerFirst(str) {
    return upperFirst$1.lowerFirst(toString(str));
}

function pad(str, length, chars = ' ') {
    return upperFirst$1.pad(toString(str), length, chars);
}

function padEnd(str, length = 0, chars = ' ') {
    return toString(str).padEnd(length, chars);
}

function padStart(str, length = 0, chars = ' ') {
    return toString(str).padStart(length, chars);
}

function repeat(str, n, guard) {
    if (guard ? isIterateeCall(str, n, guard) : n === undefined) {
        n = 1;
    }
    else {
        n = zip$1.toInteger(n);
    }
    return toString(str).repeat(n);
}

function replace(target = '', pattern, replacement) {
    if (arguments.length < 3) {
        return toString(target);
    }
    return toString(target).replace(pattern, replacement);
}

function snakeCase(str) {
    return snakeCase$1.snakeCase(normalizeForCase(str));
}

function split(string = '', separator, limit) {
    return toString(string).split(separator, limit);
}

function startCase(str) {
    const words = snakeCase$1.words(normalizeForCase(str).trim());
    let result = '';
    for (let i = 0; i < words.length; i++) {
        const word = words[i];
        if (result) {
            result += ' ';
        }
        if (word === word.toUpperCase()) {
            result += word;
        }
        else {
            result += word[0].toUpperCase() + word.slice(1).toLowerCase();
        }
    }
    return result;
}

function startsWith(str, target, position = 0) {
    return str.startsWith(target, position);
}

const esTemplateRegExp = /\$\{([^\\}]*(?:\\.[^\\}]*)*)\}/g;
const unEscapedRegExp = /['\n\r\u2028\u2029\\]/g;
const noMatchExp = /($^)/;
const escapeMap = new Map([
    ['\\', '\\'],
    ["'", "'"],
    ['\n', 'n'],
    ['\r', 'r'],
    ['\u2028', 'u2028'],
    ['\u2029', 'u2029'],
]);
function escapeString(match) {
    return `\\${escapeMap.get(match)}`;
}
const templateSettings = {
    escape: /<%-([\s\S]+?)%>/g,
    evaluate: /<%([\s\S]+?)%>/g,
    interpolate: /<%=([\s\S]+?)%>/g,
    variable: '',
    imports: {
        _: {
            escape,
            template,
        },
    },
};
function template(string, options, guard) {
    string = toString(string);
    if (guard) {
        options = templateSettings;
    }
    options = defaults({ ...options }, templateSettings);
    const delimitersRegExp = new RegExp([
        options.escape?.source ?? noMatchExp.source,
        options.interpolate?.source ?? noMatchExp.source,
        options.interpolate ? esTemplateRegExp.source : noMatchExp.source,
        options.evaluate?.source ?? noMatchExp.source,
        '$',
    ].join('|'), 'g');
    let lastIndex = 0;
    let isEvaluated = false;
    let source = `__p += ''`;
    for (const match of string.matchAll(delimitersRegExp)) {
        const [fullMatch, escapeValue, interpolateValue, esTemplateValue, evaluateValue] = match;
        const { index } = match;
        source += ` + '${string.slice(lastIndex, index).replace(unEscapedRegExp, escapeString)}'`;
        if (escapeValue) {
            source += ` + _.escape(${escapeValue})`;
        }
        if (interpolateValue) {
            source += ` + ((${interpolateValue}) == null ? '' : ${interpolateValue})`;
        }
        else if (esTemplateValue) {
            source += ` + ((${esTemplateValue}) == null ? '' : ${esTemplateValue})`;
        }
        if (evaluateValue) {
            source += `;\n${evaluateValue};\n __p += ''`;
            isEvaluated = true;
        }
        lastIndex = index + fullMatch.length;
    }
    const imports = defaults({ ...options.imports }, templateSettings.imports);
    const importsKeys = Object.keys(imports);
    const importValues = Object.values(imports);
    const sourceURL = `//# sourceURL=${options.sourceURL ? String(options.sourceURL).replace(/[\r\n]/g, ' ') : `es-toolkit.templateSource[${Date.now()}]`}\n`;
    const compiledFunction = `function(${options.variable || 'obj'}) {
    let __p = '';
    ${options.variable ? '' : 'if (obj == null) { obj = {}; }'}
    ${isEvaluated ? `function print() { __p += Array.prototype.join.call(arguments, ''); }` : ''}
    ${options.variable ? source : `with(obj) {\n${source}\n}`}
    return __p;
  }`;
    const result = attempt(() => new Function(...importsKeys, `${sourceURL}return ${compiledFunction}`)(...importValues));
    result.source = compiledFunction;
    if (result instanceof Error) {
        throw result;
    }
    return result;
}

function toLower(value) {
    return toString(value).toLowerCase();
}

function toUpper(value) {
    return toString(value).toUpperCase();
}

function trim(str, chars, guard) {
    if (str == null) {
        return '';
    }
    if (guard != null || chars == null) {
        return str.toString().trim();
    }
    switch (typeof chars) {
        case 'string': {
            return upperFirst$1.trim(str, chars.toString().split(''));
        }
        case 'object': {
            if (Array.isArray(chars)) {
                return upperFirst$1.trim(str, chars.flatMap(x => x.toString().split('')));
            }
            else {
                return upperFirst$1.trim(str, chars.toString().split(''));
            }
        }
    }
}

function trimEnd(str, chars, guard) {
    if (str == null) {
        return '';
    }
    if (guard != null || chars == null) {
        return str.toString().trimEnd();
    }
    switch (typeof chars) {
        case 'string': {
            return upperFirst$1.trimEnd(str, chars.toString().split(''));
        }
        case 'object': {
            if (Array.isArray(chars)) {
                return upperFirst$1.trimEnd(str, chars.flatMap(x => x.toString().split('')));
            }
            else {
                return upperFirst$1.trimEnd(str, chars.toString().split(''));
            }
        }
    }
}

function trimStart(str, chars, guard) {
    if (str == null) {
        return '';
    }
    if (guard != null || chars == null) {
        return str.toString().trimStart();
    }
    switch (typeof chars) {
        case 'string': {
            return upperFirst$1.trimStart(str, chars.toString().split(''));
        }
        case 'object': {
            if (Array.isArray(chars)) {
                return upperFirst$1.trimStart(str, chars.flatMap(x => x.toString().split('')));
            }
            else {
                return upperFirst$1.trimStart(str, chars.toString().split(''));
            }
        }
    }
}

function unescape(str) {
    return upperFirst$1.unescape(toString(str));
}

function upperCase(str) {
    return upperFirst$1.upperCase(normalizeForCase(str));
}

function upperFirst(str) {
    return upperFirst$1.upperFirst(toString(str));
}

const rNonCharLatin = '\\x00-\\x2f\\x3a-\\x40\\x5b-\\x60\\x7b-\\xbf\\xd7\\xf7';
const rUnicodeUpper = '\\p{Lu}';
const rUnicodeLower = '\\p{Ll}';
const rMisc = '(?:[\\p{Lm}\\p{Lo}]\\p{M}*)';
const rNumber = '\\d';
const rUnicodeOptContrLower = "(?:['\u2019](?:d|ll|m|re|s|t|ve))?";
const rUnicodeOptContrUpper = "(?:['\u2019](?:D|LL|M|RE|S|T|VE))?";
const rUnicodeBreak = `[\\p{Z}\\p{P}${rNonCharLatin}]`;
const rUnicodeMiscUpper = `(?:${rUnicodeUpper}|${rMisc})`;
const rUnicodeMiscLower = `(?:${rUnicodeLower}|${rMisc})`;
const rUnicodeWord = RegExp([
    `${rUnicodeUpper}?${rUnicodeLower}+${rUnicodeOptContrLower}(?=${rUnicodeBreak}|${rUnicodeUpper}|$)`,
    `${rUnicodeMiscUpper}+${rUnicodeOptContrUpper}(?=${rUnicodeBreak}|${rUnicodeUpper}${rUnicodeMiscLower}|$)`,
    `${rUnicodeUpper}?${rUnicodeMiscLower}+${rUnicodeOptContrLower}`,
    `${rUnicodeUpper}+${rUnicodeOptContrUpper}`,
    `${rNumber}*(?:1ST|2ND|3RD|(?![123])${rNumber}TH)(?=\\b|[a-z_])`,
    `${rNumber}*(?:1st|2nd|3rd|(?![123])${rNumber}th)(?=\\b|[A-Z_])`,
    `${rNumber}+`,
    '\\p{Emoji_Presentation}',
    '\\p{Extended_Pictographic}',
].join('|'), 'gu');
function words(str, pattern = rUnicodeWord, guard) {
    const input = toString(str);
    pattern = guard ? rUnicodeWord : pattern;
    const words = Array.from(input.match(pattern) ?? []);
    return words.filter(x => x !== '');
}

function cond(pairs) {
    const length = pairs.length;
    const processedPairs = pairs.map(pair => {
        const predicate = pair[0];
        const func = pair[1];
        if (!isWeakSet$1.isFunction(func)) {
            throw new TypeError('Expected a function');
        }
        return [iteratee(predicate), func];
    });
    return function (...args) {
        for (let i = 0; i < length; i++) {
            const pair = processedPairs[i];
            const predicate = pair[0];
            const func = pair[1];
            if (predicate.apply(this, args)) {
                return func.apply(this, args);
            }
        }
    };
}

function constant(value) {
    return () => value;
}

function defaultTo(value, defaultValue) {
    if (value == null || Number.isNaN(value)) {
        return defaultValue;
    }
    return value;
}

function gt(value, other) {
    if (typeof value === 'string' && typeof other === 'string') {
        return value > other;
    }
    return zip$1.toNumber(value) > zip$1.toNumber(other);
}

function gte(value, other) {
    if (typeof value === 'string' && typeof other === 'string') {
        return value >= other;
    }
    return zip$1.toNumber(value) >= zip$1.toNumber(other);
}

function invoke(object, path, args = []) {
    if (object == null) {
        return;
    }
    switch (typeof path) {
        case 'string': {
            if (typeof object === 'object' && Object.hasOwn(object, path)) {
                return invokeImpl(object, [path], args);
            }
            return invokeImpl(object, toPath(path), args);
        }
        case 'number':
        case 'symbol': {
            return invokeImpl(object, [path], args);
        }
        default: {
            if (Array.isArray(path)) {
                return invokeImpl(object, path, args);
            }
            else {
                return invokeImpl(object, [path], args);
            }
        }
    }
}
function invokeImpl(object, path, args) {
    const parent = get(object, path.slice(0, -1), object);
    if (parent == null) {
        return undefined;
    }
    let lastKey = last(path);
    const lastValue = lastKey?.valueOf();
    if (typeof lastValue === 'number') {
        lastKey = toKey(lastValue);
    }
    else {
        lastKey = String(lastKey);
    }
    const func = get(parent, lastKey);
    return func?.apply(parent, args);
}

function lt(value, other) {
    if (typeof value === 'string' && typeof other === 'string') {
        return value < other;
    }
    return zip$1.toNumber(value) < zip$1.toNumber(other);
}

function lte(value, other) {
    if (typeof value === 'string' && typeof other === 'string') {
        return value <= other;
    }
    return zip$1.toNumber(value) <= zip$1.toNumber(other);
}

function method(path, ...args) {
    return function (object) {
        return invoke(object, path, args);
    };
}

function methodOf(object, ...args) {
    return function (path) {
        return invoke(object, path, args);
    };
}

function now() {
    return Date.now();
}

function over(...iteratees) {
    if (iteratees.length === 1 && Array.isArray(iteratees[0])) {
        iteratees = iteratees[0];
    }
    const funcs = iteratees.map(item => iteratee(item));
    return function (...args) {
        return funcs.map(func => func.apply(this, args));
    };
}

function overEvery(...predicates) {
    return function (...values) {
        for (let i = 0; i < predicates.length; ++i) {
            const predicate = predicates[i];
            if (!Array.isArray(predicate)) {
                if (!iteratee(predicate).apply(this, values)) {
                    return false;
                }
                continue;
            }
            for (let j = 0; j < predicate.length; ++j) {
                if (!iteratee(predicate[j]).apply(this, values)) {
                    return false;
                }
            }
        }
        return true;
    };
}

function overSome(...predicates) {
    return function (...values) {
        for (let i = 0; i < predicates.length; ++i) {
            const predicate = predicates[i];
            if (!Array.isArray(predicate)) {
                if (iteratee(predicate).apply(this, values)) {
                    return true;
                }
                continue;
            }
            for (let j = 0; j < predicate.length; ++j) {
                if (iteratee(predicate[j]).apply(this, values)) {
                    return true;
                }
            }
        }
        return false;
    };
}

function stubArray() {
    return [];
}

function stubFalse() {
    return false;
}

function stubObject() {
    return {};
}

function stubString() {
    return '';
}

function stubTrue() {
    return true;
}

const MAX_ARRAY_LENGTH = 4_294_967_295;

function toLength(value) {
    if (value == null) {
        return 0;
    }
    const length = Math.floor(Number(value));
    return clamp(length, 0, MAX_ARRAY_LENGTH);
}

function toPlainObject(value) {
    const plainObject = {};
    const valueKeys = keysIn(value);
    for (let i = 0; i < valueKeys.length; i++) {
        const key = valueKeys[i];
        const objValue = value[key];
        if (key === '__proto__') {
            Object.defineProperty(plainObject, key, {
                configurable: true,
                enumerable: true,
                value: objValue,
                writable: true,
            });
        }
        else {
            plainObject[key] = objValue;
        }
    }
    return plainObject;
}

const MAX_SAFE_INTEGER = Number.MAX_SAFE_INTEGER;

function toSafeInteger(value) {
    if (value == null) {
        return 0;
    }
    return clamp(zip$1.toInteger(value), -MAX_SAFE_INTEGER, MAX_SAFE_INTEGER);
}

let idCounter = 0;
function uniqueId(prefix = '') {
    const id = ++idCounter;
    return `${prefix}${id}`;
}

const compat = /*#__PURE__*/Object.freeze(/*#__PURE__*/Object.defineProperty({
    __proto__: null,
    add,
    after,
    ary,
    assign,
    assignIn,
    assignInWith,
    assignWith,
    at,
    attempt,
    before,
    bind,
    bindAll,
    bindKey,
    camelCase,
    capitalize: snakeCase$1.capitalize,
    castArray,
    ceil,
    chunk,
    clamp,
    clone,
    cloneDeep,
    cloneDeepWith,
    cloneWith,
    compact,
    concat,
    cond,
    conforms,
    conformsTo,
    constant,
    create,
    curry,
    curryRight,
    debounce,
    deburr,
    defaultTo,
    defaults,
    defer,
    delay,
    difference,
    differenceBy,
    differenceWith,
    divide,
    drop,
    dropRight,
    dropRightWhile,
    dropWhile,
    each: forEach,
    eachRight: forEachRight,
    endsWith,
    eq: isWeakSet$1.eq,
    escape,
    escapeRegExp,
    every,
    extend: assignIn,
    extendWith: assignInWith,
    fill,
    filter,
    find,
    findIndex,
    findKey,
    findLast,
    findLastIndex,
    first: head,
    flatMap,
    flatten,
    flattenDeep,
    flattenDepth,
    flip,
    floor,
    flow,
    flowRight,
    forEach,
    forEachRight,
    forIn,
    forInRight,
    forOwn,
    forOwnRight,
    fromPairs,
    functions,
    functionsIn,
    get,
    groupBy,
    gt,
    gte,
    has,
    hasIn,
    head,
    identity: unary.identity,
    inRange,
    includes,
    indexOf,
    initial,
    intersection,
    intersectionBy,
    intersectionWith,
    invert: isPlainObject.invert,
    invertBy,
    invoke,
    invokeMap,
    isArguments,
    isArray: isPlainObject.isArray,
    isArrayBuffer,
    isArrayLike,
    isArrayLikeObject,
    isBoolean,
    isBuffer,
    isDate,
    isElement,
    isEmpty,
    isEqual: isWeakSet$1.isEqual,
    isEqualWith,
    isError,
    isFinite,
    isFunction: isWeakSet$1.isFunction,
    isInteger,
    isLength: isWeakSet$1.isLength,
    isMap,
    isMatch,
    isNaN,
    isNative,
    isNil,
    isNull: isWeakSet$1.isNull,
    isNumber,
    isObject,
    isObjectLike: isPlainObject.isObjectLike,
    isPlainObject: isPlainObject.isPlainObject,
    isRegExp,
    isSafeInteger,
    isSet,
    isString,
    isSymbol: zip$1.isSymbol,
    isTypedArray,
    isUndefined: isWeakSet$1.isUndefined,
    isWeakMap,
    isWeakSet,
    iteratee,
    join,
    kebabCase,
    keyBy,
    keys,
    keysIn,
    last,
    lastIndexOf,
    lowerCase,
    lowerFirst,
    lt,
    lte,
    map,
    mapKeys,
    mapValues,
    matches,
    matchesProperty,
    max,
    maxBy,
    mean,
    meanBy,
    memoize,
    merge,
    mergeWith,
    method,
    methodOf,
    min,
    minBy,
    multiply,
    negate,
    noop: noop.noop,
    now,
    nth,
    nthArg,
    omit,
    omitBy,
    once: unary.once,
    orderBy,
    over,
    overEvery,
    overSome,
    pad,
    padEnd,
    padStart,
    parseInt,
    partial,
    partialRight,
    partition,
    pick,
    pickBy,
    property,
    propertyOf,
    pull,
    pullAll,
    pullAllBy,
    pullAllWith,
    pullAt,
    random,
    range,
    rangeRight,
    rearg,
    reduce,
    reduceRight,
    reject,
    remove,
    repeat,
    replace,
    rest,
    result,
    reverse,
    round,
    sample,
    sampleSize,
    set,
    setWith,
    shuffle,
    size,
    slice,
    snakeCase,
    some,
    sortBy,
    sortedIndex,
    sortedIndexBy,
    sortedIndexOf,
    sortedLastIndex,
    sortedLastIndexBy,
    sortedLastIndexOf,
    split,
    spread,
    startCase,
    startsWith,
    stubArray,
    stubFalse,
    stubObject,
    stubString,
    stubTrue,
    subtract,
    sum,
    sumBy,
    tail,
    take,
    takeRight,
    takeRightWhile,
    takeWhile,
    template,
    templateSettings,
    throttle,
    times,
    toArray,
    toDefaulted,
    toFinite: zip$1.toFinite,
    toInteger: zip$1.toInteger,
    toLength,
    toLower,
    toNumber: zip$1.toNumber,
    toPairs,
    toPairsIn,
    toPath,
    toPlainObject,
    toSafeInteger,
    toString,
    toUpper,
    transform,
    trim,
    trimEnd,
    trimStart,
    unary: unary.unary,
    unescape,
    union,
    unionBy,
    unionWith,
    uniq,
    uniqBy,
    uniqWith,
    uniqueId,
    unset,
    unzip,
    unzipWith,
    update,
    updateWith,
    upperCase,
    upperFirst,
    values,
    valuesIn,
    without,
    words,
    wrap,
    xor,
    xorBy,
    xorWith,
    zip,
    zipObject,
    zipObjectDeep,
    zipWith
}, Symbol.toStringTag, { value: 'Module' }));

const toolkit = ((value) => {
    return value;
});
Object.assign(toolkit, compat);
toolkit.partial.placeholder = toolkit;
toolkit.partialRight.placeholder = toolkit;

exports.isSymbol = zip$1.isSymbol;
exports.toFinite = zip$1.toFinite;
exports.toInteger = zip$1.toInteger;
exports.toNumber = zip$1.toNumber;
exports.eq = isWeakSet$1.eq;
exports.isEqual = isWeakSet$1.isEqual;
exports.isFunction = isWeakSet$1.isFunction;
exports.isLength = isWeakSet$1.isLength;
exports.isNull = isWeakSet$1.isNull;
exports.isUndefined = isWeakSet$1.isUndefined;
exports.invert = isPlainObject.invert;
exports.isArray = isPlainObject.isArray;
exports.isObjectLike = isPlainObject.isObjectLike;
exports.isPlainObject = isPlainObject.isPlainObject;
exports.identity = unary.identity;
exports.once = unary.once;
exports.unary = unary.unary;
exports.noop = noop.noop;
exports.capitalize = snakeCase$1.capitalize;
exports.add = add;
exports.after = after;
exports.ary = ary;
exports.assign = assign;
exports.assignIn = assignIn;
exports.assignInWith = assignInWith;
exports.assignWith = assignWith;
exports.at = at;
exports.attempt = attempt;
exports.before = before;
exports.bind = bind;
exports.bindAll = bindAll;
exports.bindKey = bindKey;
exports.camelCase = camelCase;
exports.castArray = castArray;
exports.ceil = ceil;
exports.chunk = chunk;
exports.clamp = clamp;
exports.clone = clone;
exports.cloneDeep = cloneDeep;
exports.cloneDeepWith = cloneDeepWith;
exports.cloneWith = cloneWith;
exports.compact = compact;
exports.concat = concat;
exports.cond = cond;
exports.conforms = conforms;
exports.conformsTo = conformsTo;
exports.constant = constant;
exports.create = create;
exports.curry = curry;
exports.curryRight = curryRight;
exports.debounce = debounce;
exports.deburr = deburr;
exports.default = toolkit;
exports.defaultTo = defaultTo;
exports.defaults = defaults;
exports.defer = defer;
exports.delay = delay;
exports.difference = difference;
exports.differenceBy = differenceBy;
exports.differenceWith = differenceWith;
exports.divide = divide;
exports.drop = drop;
exports.dropRight = dropRight;
exports.dropRightWhile = dropRightWhile;
exports.dropWhile = dropWhile;
exports.each = forEach;
exports.eachRight = forEachRight;
exports.endsWith = endsWith;
exports.escape = escape;
exports.escapeRegExp = escapeRegExp;
exports.every = every;
exports.extend = assignIn;
exports.extendWith = assignInWith;
exports.fill = fill;
exports.filter = filter;
exports.find = find;
exports.findIndex = findIndex;
exports.findKey = findKey;
exports.findLast = findLast;
exports.findLastIndex = findLastIndex;
exports.first = head;
exports.flatMap = flatMap;
exports.flatten = flatten;
exports.flattenDeep = flattenDeep;
exports.flattenDepth = flattenDepth;
exports.flip = flip;
exports.floor = floor;
exports.flow = flow;
exports.flowRight = flowRight;
exports.forEach = forEach;
exports.forEachRight = forEachRight;
exports.forIn = forIn;
exports.forInRight = forInRight;
exports.forOwn = forOwn;
exports.forOwnRight = forOwnRight;
exports.fromPairs = fromPairs;
exports.functions = functions;
exports.functionsIn = functionsIn;
exports.get = get;
exports.groupBy = groupBy;
exports.gt = gt;
exports.gte = gte;
exports.has = has;
exports.hasIn = hasIn;
exports.head = head;
exports.inRange = inRange;
exports.includes = includes;
exports.indexOf = indexOf;
exports.initial = initial;
exports.intersection = intersection;
exports.intersectionBy = intersectionBy;
exports.intersectionWith = intersectionWith;
exports.invertBy = invertBy;
exports.invoke = invoke;
exports.invokeMap = invokeMap;
exports.isArguments = isArguments;
exports.isArrayBuffer = isArrayBuffer;
exports.isArrayLike = isArrayLike;
exports.isArrayLikeObject = isArrayLikeObject;
exports.isBoolean = isBoolean;
exports.isBuffer = isBuffer;
exports.isDate = isDate;
exports.isElement = isElement;
exports.isEmpty = isEmpty;
exports.isEqualWith = isEqualWith;
exports.isError = isError;
exports.isFinite = isFinite;
exports.isInteger = isInteger;
exports.isMap = isMap;
exports.isMatch = isMatch;
exports.isNaN = isNaN;
exports.isNative = isNative;
exports.isNil = isNil;
exports.isNumber = isNumber;
exports.isObject = isObject;
exports.isRegExp = isRegExp;
exports.isSafeInteger = isSafeInteger;
exports.isSet = isSet;
exports.isString = isString;
exports.isTypedArray = isTypedArray;
exports.isWeakMap = isWeakMap;
exports.isWeakSet = isWeakSet;
exports.iteratee = iteratee;
exports.join = join;
exports.kebabCase = kebabCase;
exports.keyBy = keyBy;
exports.keys = keys;
exports.keysIn = keysIn;
exports.last = last;
exports.lastIndexOf = lastIndexOf;
exports.lowerCase = lowerCase;
exports.lowerFirst = lowerFirst;
exports.lt = lt;
exports.lte = lte;
exports.map = map;
exports.mapKeys = mapKeys;
exports.mapValues = mapValues;
exports.matches = matches;
exports.matchesProperty = matchesProperty;
exports.max = max;
exports.maxBy = maxBy;
exports.mean = mean;
exports.meanBy = meanBy;
exports.memoize = memoize;
exports.merge = merge;
exports.mergeWith = mergeWith;
exports.method = method;
exports.methodOf = methodOf;
exports.min = min;
exports.minBy = minBy;
exports.multiply = multiply;
exports.negate = negate;
exports.now = now;
exports.nth = nth;
exports.nthArg = nthArg;
exports.omit = omit;
exports.omitBy = omitBy;
exports.orderBy = orderBy;
exports.over = over;
exports.overEvery = overEvery;
exports.overSome = overSome;
exports.pad = pad;
exports.padEnd = padEnd;
exports.padStart = padStart;
exports.parseInt = parseInt;
exports.partial = partial;
exports.partialRight = partialRight;
exports.partition = partition;
exports.pick = pick;
exports.pickBy = pickBy;
exports.property = property;
exports.propertyOf = propertyOf;
exports.pull = pull;
exports.pullAll = pullAll;
exports.pullAllBy = pullAllBy;
exports.pullAllWith = pullAllWith;
exports.pullAt = pullAt;
exports.random = random;
exports.range = range;
exports.rangeRight = rangeRight;
exports.rearg = rearg;
exports.reduce = reduce;
exports.reduceRight = reduceRight;
exports.reject = reject;
exports.remove = remove;
exports.repeat = repeat;
exports.replace = replace;
exports.rest = rest;
exports.result = result;
exports.reverse = reverse;
exports.round = round;
exports.sample = sample;
exports.sampleSize = sampleSize;
exports.set = set;
exports.setWith = setWith;
exports.shuffle = shuffle;
exports.size = size;
exports.slice = slice;
exports.snakeCase = snakeCase;
exports.some = some;
exports.sortBy = sortBy;
exports.sortedIndex = sortedIndex;
exports.sortedIndexBy = sortedIndexBy;
exports.sortedIndexOf = sortedIndexOf;
exports.sortedLastIndex = sortedLastIndex;
exports.sortedLastIndexBy = sortedLastIndexBy;
exports.sortedLastIndexOf = sortedLastIndexOf;
exports.split = split;
exports.spread = spread;
exports.startCase = startCase;
exports.startsWith = startsWith;
exports.stubArray = stubArray;
exports.stubFalse = stubFalse;
exports.stubObject = stubObject;
exports.stubString = stubString;
exports.stubTrue = stubTrue;
exports.subtract = subtract;
exports.sum = sum;
exports.sumBy = sumBy;
exports.tail = tail;
exports.take = take;
exports.takeRight = takeRight;
exports.takeRightWhile = takeRightWhile;
exports.takeWhile = takeWhile;
exports.template = template;
exports.templateSettings = templateSettings;
exports.throttle = throttle;
exports.times = times;
exports.toArray = toArray;
exports.toDefaulted = toDefaulted;
exports.toLength = toLength;
exports.toLower = toLower;
exports.toPairs = toPairs;
exports.toPairsIn = toPairsIn;
exports.toPath = toPath;
exports.toPlainObject = toPlainObject;
exports.toSafeInteger = toSafeInteger;
exports.toString = toString;
exports.toUpper = toUpper;
exports.transform = transform;
exports.trim = trim;
exports.trimEnd = trimEnd;
exports.trimStart = trimStart;
exports.unescape = unescape;
exports.union = union;
exports.unionBy = unionBy;
exports.unionWith = unionWith;
exports.uniq = uniq;
exports.uniqBy = uniqBy;
exports.uniqWith = uniqWith;
exports.uniqueId = uniqueId;
exports.unset = unset;
exports.unzip = unzip;
exports.unzipWith = unzipWith;
exports.update = update;
exports.updateWith = updateWith;
exports.upperCase = upperCase;
exports.upperFirst = upperFirst;
exports.values = values;
exports.valuesIn = valuesIn;
exports.without = without;
exports.words = words;
exports.wrap = wrap;
exports.xor = xor;
exports.xorBy = xorBy;
exports.xorWith = xorWith;
exports.zip = zip;
exports.zipObject = zipObject;
exports.zipObjectDeep = zipObjectDeep;
exports.zipWith = zipWith;
