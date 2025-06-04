import { snakeCase as snakeCase$1 } from '../../string/snakeCase.mjs';
import { normalizeForCase } from '../_internal/normalizeForCase.mjs';

function snakeCase(str) {
    return snakeCase$1(normalizeForCase(str));
}

export { snakeCase };
