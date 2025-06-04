import { cloneDeepWith as cloneDeepWith$1, copyProperties } from '../../object/cloneDeepWith.mjs';
import { argumentsTag, booleanTag, stringTag, numberTag } from '../_internal/tags.mjs';

function cloneDeepWith(obj, cloneValue) {
    return cloneDeepWith$1(obj, (value, key, object, stack) => {
        const cloned = cloneValue?.(value, key, object, stack);
        if (cloned != null) {
            return cloned;
        }
        if (typeof obj !== 'object') {
            return undefined;
        }
        switch (Object.prototype.toString.call(obj)) {
            case numberTag:
            case stringTag:
            case booleanTag: {
                const result = new obj.constructor(obj?.valueOf());
                copyProperties(result, obj);
                return result;
            }
            case argumentsTag: {
                const result = {};
                copyProperties(result, obj);
                result.length = obj.length;
                result[Symbol.iterator] = obj[Symbol.iterator];
                return result;
            }
            default: {
                return undefined;
            }
        }
    });
}

export { cloneDeepWith };
