import { isPlainObject } from '../predicate/isPlainObject.mjs';

function merge(target, source) {
    const sourceKeys = Object.keys(source);
    for (let i = 0; i < sourceKeys.length; i++) {
        const key = sourceKeys[i];
        const sourceValue = source[key];
        const targetValue = target[key];
        if (Array.isArray(sourceValue)) {
            if (Array.isArray(targetValue)) {
                target[key] = merge(targetValue, sourceValue);
            }
            else {
                target[key] = merge([], sourceValue);
            }
        }
        else if (isPlainObject(sourceValue)) {
            if (isPlainObject(targetValue)) {
                target[key] = merge(targetValue, sourceValue);
            }
            else {
                target[key] = merge({}, sourceValue);
            }
        }
        else if (targetValue === undefined || sourceValue !== undefined) {
            target[key] = sourceValue;
        }
    }
    return target;
}

export { merge };
