import { assignValue } from '../_internal/assignValue.mjs';
import { isIndex } from '../_internal/isIndex.mjs';
import { isKey } from '../_internal/isKey.mjs';
import { toKey } from '../_internal/toKey.mjs';
import { isObject } from '../predicate/isObject.mjs';
import { toPath } from '../util/toPath.mjs';

function updateWith(obj, path, updater, customizer) {
    if (obj == null && !isObject(obj)) {
        return obj;
    }
    const resolvedPath = isKey(path, obj)
        ? [path]
        : Array.isArray(path)
            ? path
            : typeof path === 'string'
                ? toPath(path)
                : [path];
    let current = obj;
    for (let i = 0; i < resolvedPath.length && current != null; i++) {
        const key = toKey(resolvedPath[i]);
        let newValue;
        if (i === resolvedPath.length - 1) {
            newValue = updater(current[key]);
        }
        else {
            const objValue = current[key];
            const customizerResult = customizer(objValue);
            newValue =
                customizerResult !== undefined
                    ? customizerResult
                    : isObject(objValue)
                        ? objValue
                        : isIndex(resolvedPath[i + 1])
                            ? []
                            : {};
        }
        assignValue(current, key, newValue);
        current = current[key];
    }
    return obj;
}

export { updateWith };
