import { kebabCase as kebabCase$1 } from '../../string/kebabCase.mjs';
import { normalizeForCase } from '../_internal/normalizeForCase.mjs';

function kebabCase(str) {
    return kebabCase$1(normalizeForCase(str));
}

export { kebabCase };
