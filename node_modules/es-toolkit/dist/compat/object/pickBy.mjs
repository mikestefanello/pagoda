import { keysIn } from './keysIn.mjs';
import { range } from '../../math/range.mjs';
import { getSymbolsIn } from '../_internal/getSymbolsIn.mjs';
import { isArrayLike } from '../predicate/isArrayLike.mjs';
import { isSymbol } from '../predicate/isSymbol.mjs';

function pickBy(obj, shouldPick) {
    if (obj == null) {
        return {};
    }
    const result = {};
    if (shouldPick == null) {
        return obj;
    }
    const keys = isArrayLike(obj) ? range(0, obj.length) : [...keysIn(obj), ...getSymbolsIn(obj)];
    for (let i = 0; i < keys.length; i++) {
        const key = (isSymbol(keys[i]) ? keys[i] : keys[i].toString());
        const value = obj[key];
        if (shouldPick(value, key, obj)) {
            result[key] = value;
        }
    }
    return result;
}

export { pickBy };
