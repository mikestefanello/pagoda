import { sortedIndex } from './sortedIndex.mjs';
import { eq } from '../util/eq.mjs';

function sortedIndexOf(array, value) {
    if (!array?.length) {
        return -1;
    }
    const index = sortedIndex(array, value);
    if (index < array.length && eq(array[index], value)) {
        return index;
    }
    return -1;
}

export { sortedIndexOf };
