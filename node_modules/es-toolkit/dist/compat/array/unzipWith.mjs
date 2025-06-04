import { unzip } from '../../array/unzip.mjs';
import { isArray } from '../predicate/isArray.mjs';
import { isArrayLikeObject } from '../predicate/isArrayLikeObject.mjs';

function unzipWith(array, iteratee) {
    if (!isArrayLikeObject(array) || !array.length) {
        return [];
    }
    const unziped = isArray(array) ? unzip(array) : unzip(Array.from(array, value => Array.from(value)));
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

export { unzipWith };
