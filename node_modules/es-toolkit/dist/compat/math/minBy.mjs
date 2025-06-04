import { minBy as minBy$1 } from '../../array/minBy.mjs';
import { iteratee } from '../util/iteratee.mjs';

function minBy(items, iteratee$1) {
    if (items == null) {
        return undefined;
    }
    return minBy$1(Array.from(items), iteratee(iteratee$1));
}

export { minBy };
