'use strict';

Object.defineProperty(exports, Symbol.toStringTag, { value: 'Module' });

const range = require('../_chunk/range-DSpBDL.js');
const randomInt = require('../_chunk/randomInt-CF7bZK.js');

function median(nums) {
    if (nums.length === 0) {
        return NaN;
    }
    const sorted = nums.slice().sort((a, b) => a - b);
    const middleIndex = Math.floor(sorted.length / 2);
    if (sorted.length % 2 === 0) {
        return (sorted[middleIndex - 1] + sorted[middleIndex]) / 2;
    }
    else {
        return sorted[middleIndex];
    }
}

function medianBy(items, getValue) {
    const nums = items.map(x => getValue(x));
    return median(nums);
}

function rangeRight(start, end, step = 1) {
    if (end == null) {
        end = start;
        start = 0;
    }
    if (!Number.isInteger(step) || step === 0) {
        throw new Error(`The step value must be a non-zero integer.`);
    }
    const length = Math.max(Math.ceil((end - start) / step), 0);
    const result = new Array(length);
    for (let i = 0; i < length; i++) {
        result[i] = start + (length - i - 1) * step;
    }
    return result;
}

function round(value, precision = 0) {
    if (!Number.isInteger(precision)) {
        throw new Error('Precision must be an integer.');
    }
    const multiplier = Math.pow(10, precision);
    return Math.round(value * multiplier) / multiplier;
}

function sumBy(items, getValue) {
    let result = 0;
    for (let i = 0; i < items.length; i++) {
        result += getValue(items[i]);
    }
    return result;
}

exports.clamp = range.clamp;
exports.inRange = range.inRange;
exports.mean = range.mean;
exports.meanBy = range.meanBy;
exports.range = range.range;
exports.sum = range.sum;
exports.random = randomInt.random;
exports.randomInt = randomInt.randomInt;
exports.median = median;
exports.medianBy = medianBy;
exports.rangeRight = rangeRight;
exports.round = round;
exports.sumBy = sumBy;
