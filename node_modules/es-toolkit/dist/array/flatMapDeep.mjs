import { flattenDeep } from './flattenDeep.mjs';

function flatMapDeep(arr, iteratee) {
    return flattenDeep(arr.map((item) => iteratee(item)));
}

export { flatMapDeep };
