import { trimStart as trimStart$1 } from '../../string/trimStart.mjs';

function trimStart(str, chars, guard) {
    if (str == null) {
        return '';
    }
    if (guard != null || chars == null) {
        return str.toString().trimStart();
    }
    switch (typeof chars) {
        case 'string': {
            return trimStart$1(str, chars.toString().split(''));
        }
        case 'object': {
            if (Array.isArray(chars)) {
                return trimStart$1(str, chars.flatMap(x => x.toString().split('')));
            }
            else {
                return trimStart$1(str, chars.toString().split(''));
            }
        }
    }
}

export { trimStart };
