import { filter } from './filter.mjs';
import { negate } from '../function/negate.mjs';
import { iteratee } from '../util/iteratee.mjs';

function reject(source, predicate) {
    return filter(source, negate(iteratee(predicate)));
}

export { reject };
