import { flatten } from './flatten.mjs';

function flatMap(arr, iteratee, depth = 1) {
    return flatten(arr.map(item => iteratee(item)), depth);
}

export { flatMap };
