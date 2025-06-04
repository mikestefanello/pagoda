import { flatten } from './flatten.mjs';

function flattenDeep(value) {
    return flatten(value, Infinity);
}

export { flattenDeep };
