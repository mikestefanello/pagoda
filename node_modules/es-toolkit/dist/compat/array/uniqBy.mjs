import { uniqBy as uniqBy$1 } from '../../array/uniqBy.mjs';
import { isArrayLikeObject } from '../predicate/isArrayLikeObject.mjs';
import { iteratee } from '../util/iteratee.mjs';

function uniqBy(array, iteratee$1) {
    if (!isArrayLikeObject(array)) {
        return [];
    }
    return uniqBy$1(Array.from(array), iteratee(iteratee$1));
}

export { uniqBy };
