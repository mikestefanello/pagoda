import { trimEnd as trimEnd$1 } from '../../string/trimEnd.mjs';

function trimEnd(str, chars, guard) {
    if (str == null) {
        return '';
    }
    if (guard != null || chars == null) {
        return str.toString().trimEnd();
    }
    switch (typeof chars) {
        case 'string': {
            return trimEnd$1(str, chars.toString().split(''));
        }
        case 'object': {
            if (Array.isArray(chars)) {
                return trimEnd$1(str, chars.flatMap(x => x.toString().split('')));
            }
            else {
                return trimEnd$1(str, chars.toString().split(''));
            }
        }
    }
}

export { trimEnd };
