import { isArrayLike } from '../predicate/isArrayLike.mjs';

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

export { fromPairs };
