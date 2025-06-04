import { property } from './property.mjs';
import { findKey as findKey$1 } from '../../object/findKey.mjs';
import { isObject } from '../predicate/isObject.mjs';
import { matches } from '../predicate/matches.mjs';
import { matchesProperty } from '../predicate/matchesProperty.mjs';

function findKey(obj, predicate) {
    if (!isObject(obj)) {
        return undefined;
    }
    return findKeyImpl(obj, predicate);
}
function findKeyImpl(obj, predicate) {
    if (typeof predicate === 'function') {
        return findKey$1(obj, predicate);
    }
    if (typeof predicate === 'object') {
        if (Array.isArray(predicate)) {
            const key = predicate[0];
            const value = predicate[1];
            return findKey$1(obj, matchesProperty(key, value));
        }
        return findKey$1(obj, matches(predicate));
    }
    if (typeof predicate === 'string') {
        return findKey$1(obj, property(predicate));
    }
}

export { findKey };
