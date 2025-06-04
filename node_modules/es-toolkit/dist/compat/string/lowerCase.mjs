import { lowerCase as lowerCase$1 } from '../../string/lowerCase.mjs';
import { normalizeForCase } from '../_internal/normalizeForCase.mjs';

function lowerCase(str) {
    return lowerCase$1(normalizeForCase(str));
}

export { lowerCase };
