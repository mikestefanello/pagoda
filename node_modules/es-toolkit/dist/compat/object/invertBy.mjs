import { identity } from '../../function/identity.mjs';
import { isNil } from '../../predicate/isNil.mjs';

function invertBy(object, iteratee) {
    const result = {};
    if (isNil(object)) {
        return result;
    }
    if (iteratee == null) {
        iteratee = identity;
    }
    const keys = Object.keys(object);
    for (let i = 0; i < keys.length; i++) {
        const key = keys[i];
        const value = object[key];
        const valueStr = iteratee(value);
        if (Array.isArray(result[valueStr])) {
            result[valueStr].push(key);
        }
        else {
            result[valueStr] = [key];
        }
    }
    return result;
}

export { invertBy };
