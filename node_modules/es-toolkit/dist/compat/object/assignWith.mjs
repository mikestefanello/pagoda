import { keys } from './keys.mjs';
import { eq } from '../util/eq.mjs';

function assignWith(object, ...sources) {
    let getValueToAssign = sources[sources.length - 1];
    if (typeof getValueToAssign === 'function') {
        sources.pop();
    }
    else {
        getValueToAssign = undefined;
    }
    for (let i = 0; i < sources.length; i++) {
        assignWithImpl(object, sources[i], getValueToAssign);
    }
    return object;
}
function assignWithImpl(object, source, getValueToAssign) {
    const keys$1 = keys(source);
    for (let i = 0; i < keys$1.length; i++) {
        const key = keys$1[i];
        const objValue = object[key];
        const srcValue = source[key];
        const newValue = getValueToAssign?.(objValue, srcValue, key, object, source) ?? srcValue;
        if (!(key in object) || !eq(objValue, newValue)) {
            object[key] = newValue;
        }
    }
}

export { assignWith };
