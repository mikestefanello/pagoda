import { cloneDeep } from './cloneDeep.mjs';
import { clone } from '../../object/clone.mjs';
import { isPrimitive } from '../../predicate/isPrimitive.mjs';
import { getSymbols } from '../_internal/getSymbols.mjs';
import { isArguments } from '../predicate/isArguments.mjs';
import { isObjectLike } from '../predicate/isObjectLike.mjs';
import { isPlainObject } from '../predicate/isPlainObject.mjs';
import { isTypedArray } from '../predicate/isTypedArray.mjs';

function mergeWith(object, ...otherArgs) {
    const sources = otherArgs.slice(0, -1);
    const merge = otherArgs[otherArgs.length - 1];
    let result = object;
    for (let i = 0; i < sources.length; i++) {
        const source = sources[i];
        result = mergeWithDeep(result, source, merge, new Map());
    }
    return result;
}
function mergeWithDeep(target, source, merge, stack) {
    if (isPrimitive(target)) {
        target = Object(target);
    }
    if (source == null || typeof source !== 'object') {
        return target;
    }
    if (stack.has(source)) {
        return clone(stack.get(source));
    }
    stack.set(source, target);
    if (Array.isArray(source)) {
        source = source.slice();
        for (let i = 0; i < source.length; i++) {
            source[i] = source[i] ?? undefined;
        }
    }
    const sourceKeys = [...Object.keys(source), ...getSymbols(source)];
    for (let i = 0; i < sourceKeys.length; i++) {
        const key = sourceKeys[i];
        let sourceValue = source[key];
        let targetValue = target[key];
        if (isArguments(sourceValue)) {
            sourceValue = { ...sourceValue };
        }
        if (isArguments(targetValue)) {
            targetValue = { ...targetValue };
        }
        if (typeof Buffer !== 'undefined' && Buffer.isBuffer(sourceValue)) {
            sourceValue = cloneDeep(sourceValue);
        }
        if (Array.isArray(sourceValue)) {
            if (typeof targetValue === 'object' && targetValue != null) {
                const cloned = [];
                const targetKeys = Reflect.ownKeys(targetValue);
                for (let i = 0; i < targetKeys.length; i++) {
                    const targetKey = targetKeys[i];
                    cloned[targetKey] = targetValue[targetKey];
                }
                targetValue = cloned;
            }
            else {
                targetValue = [];
            }
        }
        const merged = merge(targetValue, sourceValue, key, target, source, stack);
        if (merged != null) {
            target[key] = merged;
        }
        else if (Array.isArray(sourceValue)) {
            target[key] = mergeWithDeep(targetValue, sourceValue, merge, stack);
        }
        else if (isObjectLike(targetValue) && isObjectLike(sourceValue)) {
            target[key] = mergeWithDeep(targetValue, sourceValue, merge, stack);
        }
        else if (targetValue == null && isPlainObject(sourceValue)) {
            target[key] = mergeWithDeep({}, sourceValue, merge, stack);
        }
        else if (targetValue == null && isTypedArray(sourceValue)) {
            target[key] = cloneDeep(sourceValue);
        }
        else if (targetValue === undefined || sourceValue !== undefined) {
            target[key] = sourceValue;
        }
    }
    return target;
}

export { mergeWith };
