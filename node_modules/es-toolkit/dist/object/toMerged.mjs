import { cloneDeep } from './cloneDeep.mjs';
import { merge } from './merge.mjs';

function toMerged(target, source) {
    return merge(cloneDeep(target), source);
}

export { toMerged };
