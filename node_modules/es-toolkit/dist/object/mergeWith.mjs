import { isObjectLike } from '../compat/predicate/isObjectLike.mjs';

function mergeWith(target, source, merge) {
    const sourceKeys = Object.keys(source);
    for (let i = 0; i < sourceKeys.length; i++) {
        const key = sourceKeys[i];
        const sourceValue = source[key];
        const targetValue = target[key];
        const merged = merge(targetValue, sourceValue, key, target, source);
        if (merged != null) {
            target[key] = merged;
        }
        else if (Array.isArray(sourceValue)) {
            target[key] = mergeWith(targetValue ?? [], sourceValue, merge);
        }
        else if (isObjectLike(targetValue) && isObjectLike(sourceValue)) {
            target[key] = mergeWith(targetValue ?? {}, sourceValue, merge);
        }
        else if (targetValue === undefined || sourceValue !== undefined) {
            target[key] = sourceValue;
        }
    }
    return target;
}

export { mergeWith };
