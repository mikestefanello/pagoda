import { toInteger } from '../compat/util/toInteger.mjs';

function takeRight(arr, count = 1, guard) {
    count = guard || count === undefined ? 1 : toInteger(count);
    if (count <= 0 || arr == null || arr.length === 0) {
        return [];
    }
    return arr.slice(-count);
}

export { takeRight };
