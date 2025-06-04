import { iteratee } from '../util/iteratee.mjs';

function find(source, _doesMatch, fromIndex = 0) {
    if (!source) {
        return undefined;
    }
    if (fromIndex < 0) {
        fromIndex = Math.max(source.length + fromIndex, 0);
    }
    const doesMatch = iteratee(_doesMatch);
    if (typeof doesMatch === 'function' && !Array.isArray(source)) {
        const keys = Object.keys(source);
        for (let i = fromIndex; i < keys.length; i++) {
            const key = keys[i];
            const value = source[key];
            if (doesMatch(value, key, source)) {
                return value;
            }
        }
        return undefined;
    }
    const values = Array.isArray(source) ? source.slice(fromIndex) : Object.values(source).slice(fromIndex);
    return values.find(doesMatch);
}

export { find };
