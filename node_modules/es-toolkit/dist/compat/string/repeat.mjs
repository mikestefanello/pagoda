import { isIterateeCall } from '../_internal/isIterateeCall.mjs';
import { toInteger } from '../util/toInteger.mjs';
import { toString } from '../util/toString.mjs';

function repeat(str, n, guard) {
    if (guard ? isIterateeCall(str, n, guard) : n === undefined) {
        n = 1;
    }
    else {
        n = toInteger(n);
    }
    return toString(str).repeat(n);
}

export { repeat };
