import { flatten } from './flatten.mjs';

function flattenDepth(value, depth = 1) {
    return flatten(value, depth);
}

export { flattenDepth };
