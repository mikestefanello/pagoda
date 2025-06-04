import { flatten } from './flatten.mjs';
import { uniq } from '../../array/uniq.mjs';
import { isArrayLikeObject } from '../predicate/isArrayLikeObject.mjs';

function union(...arrays) {
    const validArrays = arrays.filter(isArrayLikeObject);
    const flattened = flatten(validArrays, 1);
    return uniq(flattened);
}

export { union };
