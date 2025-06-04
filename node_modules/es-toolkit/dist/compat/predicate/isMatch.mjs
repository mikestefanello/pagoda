import { isObject } from './isObject.mjs';
import { isPrimitive } from '../../predicate/isPrimitive.mjs';
import { eq } from '../util/eq.mjs';

function isMatch(target, source) {
    if (source === target) {
        return true;
    }
    switch (typeof source) {
        case 'object': {
            if (source == null) {
                return true;
            }
            const keys = Object.keys(source);
            if (target == null) {
                return keys.length === 0;
            }
            if (Array.isArray(source)) {
                return isArrayMatch(target, source);
            }
            if (source instanceof Map) {
                return isMapMatch(target, source);
            }
            if (source instanceof Set) {
                return isSetMatch(target, source);
            }
            for (let i = 0; i < keys.length; i++) {
                const key = keys[i];
                if (!isPrimitive(target) && !(key in target)) {
                    return false;
                }
                if (source[key] === undefined && target[key] !== undefined) {
                    return false;
                }
                if (source[key] === null && target[key] !== null) {
                    return false;
                }
                if (!isMatch(target[key], source[key])) {
                    return false;
                }
            }
            return true;
        }
        case 'function': {
            if (Object.keys(source).length > 0) {
                return isMatch(target, { ...source });
            }
            return false;
        }
        default: {
            if (!isObject(target)) {
                return eq(target, source);
            }
            return !source;
        }
    }
}
function isMapMatch(target, source) {
    if (source.size === 0) {
        return true;
    }
    if (!(target instanceof Map)) {
        return false;
    }
    for (const [key, value] of source.entries()) {
        if (!isMatch(target.get(key), value)) {
            return false;
        }
    }
    return true;
}
function isArrayMatch(target, source) {
    if (source.length === 0) {
        return true;
    }
    if (!Array.isArray(target)) {
        return false;
    }
    const countedIndex = new Set();
    for (let i = 0; i < source.length; i++) {
        const sourceItem = source[i];
        const index = target.findIndex((targetItem, index) => {
            return isMatch(targetItem, sourceItem) && !countedIndex.has(index);
        });
        if (index === -1) {
            return false;
        }
        countedIndex.add(index);
    }
    return true;
}
function isSetMatch(target, source) {
    if (source.size === 0) {
        return true;
    }
    if (!(target instanceof Set)) {
        return false;
    }
    return isArrayMatch([...target], [...source]);
}

export { isArrayMatch, isMapMatch, isMatch, isSetMatch };
