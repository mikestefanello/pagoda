'use strict';

Object.defineProperty(exports, Symbol.toStringTag, { value: 'Module' });

const zip = require('../_chunk/zip-Cyyp17.js');

function at(arr, indices) {
    const result = new Array(indices.length);
    const length = arr.length;
    for (let i = 0; i < indices.length; i++) {
        let index = indices[i];
        index = Number.isInteger(index) ? index : Math.trunc(index) || 0;
        if (index < 0) {
            index += length;
        }
        result[i] = arr[index];
    }
    return result;
}

function countBy(arr, mapper) {
    const result = {};
    for (let i = 0; i < arr.length; i++) {
        const item = arr[i];
        const key = mapper(item);
        result[key] = (result[key] ?? 0) + 1;
    }
    return result;
}

function flatMap(arr, iteratee, depth = 1) {
    return zip.flatten(arr.map(item => iteratee(item)), depth);
}

function flattenDeep(arr) {
    return zip.flatten(arr, Infinity);
}

function flatMapDeep(arr, iteratee) {
    return flattenDeep(arr.map((item) => iteratee(item)));
}

function forEachRight(arr, callback) {
    for (let i = arr.length - 1; i >= 0; i--) {
        const element = arr[i];
        callback(element, i, arr);
    }
}

function isSubset(superset, subset) {
    return zip.difference(subset, superset).length === 0;
}

function isSubsetWith(superset, subset, areItemsEqual) {
    return zip.differenceWith(subset, superset, areItemsEqual).length === 0;
}

function keyBy(arr, getKeyFromItem) {
    const result = {};
    for (let i = 0; i < arr.length; i++) {
        const item = arr[i];
        const key = getKeyFromItem(item);
        result[key] = item;
    }
    return result;
}

function compareValues(a, b, order) {
    if (a < b) {
        return order === 'asc' ? -1 : 1;
    }
    if (a > b) {
        return order === 'asc' ? 1 : -1;
    }
    return 0;
}

function orderBy(arr, criteria, orders) {
    return arr.slice().sort((a, b) => {
        const ordersLength = orders.length;
        for (let i = 0; i < criteria.length; i++) {
            const order = ordersLength > i ? orders[i] : orders[ordersLength - 1];
            const criterion = criteria[i];
            const criterionIsFunction = typeof criterion === 'function';
            const valueA = criterionIsFunction ? criterion(a) : a[criterion];
            const valueB = criterionIsFunction ? criterion(b) : b[criterion];
            const result = compareValues(valueA, valueB, order);
            if (result !== 0) {
                return result;
            }
        }
        return 0;
    });
}

function partition(arr, isInTruthy) {
    const truthy = [];
    const falsy = [];
    for (let i = 0; i < arr.length; i++) {
        const item = arr[i];
        if (isInTruthy(item)) {
            truthy.push(item);
        }
        else {
            falsy.push(item);
        }
    }
    return [truthy, falsy];
}

function pullAt(arr, indicesToRemove) {
    const removed = at(arr, indicesToRemove);
    const indices = new Set(indicesToRemove.slice().sort((x, y) => y - x));
    for (const index of indices) {
        arr.splice(index, 1);
    }
    return removed;
}

function sortBy(arr, criteria) {
    return orderBy(arr, criteria, ['asc']);
}

function takeRightWhile(arr, shouldContinueTaking) {
    for (let i = arr.length - 1; i >= 0; i--) {
        if (!shouldContinueTaking(arr[i])) {
            return arr.slice(i + 1);
        }
    }
    return arr.slice();
}

function takeWhile(arr, shouldContinueTaking) {
    const result = [];
    for (let i = 0; i < arr.length; i++) {
        const item = arr[i];
        if (!shouldContinueTaking(item)) {
            break;
        }
        result.push(item);
    }
    return result;
}

function toFilled(arr, value, start = 0, end = arr.length) {
    const length = arr.length;
    const finalStart = Math.max(start >= 0 ? start : length + start, 0);
    const finalEnd = Math.min(end >= 0 ? end : length + end, length);
    const newArr = arr.slice();
    for (let i = finalStart; i < finalEnd; i++) {
        newArr[i] = value;
    }
    return newArr;
}

function union(arr1, arr2) {
    return zip.uniq(arr1.concat(arr2));
}

function unionBy(arr1, arr2, mapper) {
    return zip.uniqBy(arr1.concat(arr2), mapper);
}

function unionWith(arr1, arr2, areItemsEqual) {
    return zip.uniqWith(arr1.concat(arr2), areItemsEqual);
}

function unzipWith(target, iteratee) {
    const maxLength = Math.max(...target.map(innerArray => innerArray.length));
    const result = new Array(maxLength);
    for (let i = 0; i < maxLength; i++) {
        const group = new Array(target.length);
        for (let j = 0; j < target.length; j++) {
            group[j] = target[j][i];
        }
        result[i] = iteratee(...group);
    }
    return result;
}

function xor(arr1, arr2) {
    return zip.difference(union(arr1, arr2), zip.intersection(arr1, arr2));
}

function xorBy(arr1, arr2, mapper) {
    const union = unionBy(arr1, arr2, mapper);
    const intersection = zip.intersectionBy(arr1, arr2, mapper);
    return zip.differenceBy(union, intersection, mapper);
}

function xorWith(arr1, arr2, areElementsEqual) {
    const union = unionWith(arr1, arr2, areElementsEqual);
    const intersection = zip.intersectionWith(arr1, arr2, areElementsEqual);
    return zip.differenceWith(union, intersection, areElementsEqual);
}

function zipObject(keys, values) {
    const result = {};
    for (let i = 0; i < keys.length; i++) {
        result[keys[i]] = values[i];
    }
    return result;
}

function zipWith(arr1, ...rest) {
    const arrs = [arr1, ...rest.slice(0, -1)];
    const combine = rest[rest.length - 1];
    const maxIndex = Math.max(...arrs.map(arr => arr.length));
    const result = Array(maxIndex);
    for (let i = 0; i < maxIndex; i++) {
        const elements = arrs.map(arr => arr[i]);
        result[i] = combine(...elements);
    }
    return result;
}

exports.chunk = zip.chunk;
exports.compact = zip.compact;
exports.difference = zip.difference;
exports.differenceBy = zip.differenceBy;
exports.differenceWith = zip.differenceWith;
exports.drop = zip.drop;
exports.dropRight = zip.dropRight;
exports.dropRightWhile = zip.dropRightWhile;
exports.dropWhile = zip.dropWhile;
exports.fill = zip.fill;
exports.flatten = zip.flatten;
exports.groupBy = zip.groupBy;
exports.head = zip.head;
exports.initial = zip.initial;
exports.intersection = zip.intersection;
exports.intersectionBy = zip.intersectionBy;
exports.intersectionWith = zip.intersectionWith;
exports.last = zip.last;
exports.maxBy = zip.maxBy;
exports.minBy = zip.minBy;
exports.pull = zip.pull;
exports.remove = zip.remove;
exports.sample = zip.sample;
exports.sampleSize = zip.sampleSize;
exports.shuffle = zip.shuffle;
exports.tail = zip.tail;
exports.take = zip.take;
exports.takeRight = zip.takeRight;
exports.uniq = zip.uniq;
exports.uniqBy = zip.uniqBy;
exports.uniqWith = zip.uniqWith;
exports.unzip = zip.unzip;
exports.windowed = zip.windowed;
exports.without = zip.without;
exports.zip = zip.zip;
exports.at = at;
exports.countBy = countBy;
exports.flatMap = flatMap;
exports.flatMapDeep = flatMapDeep;
exports.flattenDeep = flattenDeep;
exports.forEachRight = forEachRight;
exports.isSubset = isSubset;
exports.isSubsetWith = isSubsetWith;
exports.keyBy = keyBy;
exports.orderBy = orderBy;
exports.partition = partition;
exports.pullAt = pullAt;
exports.sortBy = sortBy;
exports.takeRightWhile = takeRightWhile;
exports.takeWhile = takeWhile;
exports.toFilled = toFilled;
exports.union = union;
exports.unionBy = unionBy;
exports.unionWith = unionWith;
exports.unzipWith = unzipWith;
exports.xor = xor;
exports.xorBy = xorBy;
exports.xorWith = xorWith;
exports.zipObject = zipObject;
exports.zipWith = zipWith;
