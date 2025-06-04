import { sortedLastIndex } from './sortedLastIndex.mjs';
import { eq } from '../util/eq.mjs';

function sortedLastIndexOf(array, value) {
    if (!array?.length) {
        return -1;
    }
    const index = sortedLastIndex(array, value) - 1;
    if (index >= 0 && eq(array[index], value)) {
        return index;
    }
    return -1;
}

export { sortedLastIndexOf };
