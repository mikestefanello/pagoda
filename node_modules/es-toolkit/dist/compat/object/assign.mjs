import { keys } from './keys.mjs';
import { eq } from '../util/eq.mjs';

function assign(object, ...sources) {
    for (let i = 0; i < sources.length; i++) {
        assignImpl(object, sources[i]);
    }
    return object;
}
function assignImpl(object, source) {
    const keys$1 = keys(source);
    for (let i = 0; i < keys$1.length; i++) {
        const key = keys$1[i];
        if (!(key in object) || !eq(object[key], source[key])) {
            object[key] = source[key];
        }
    }
}

export { assign };
