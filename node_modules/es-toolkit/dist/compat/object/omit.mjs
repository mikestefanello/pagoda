import { unset } from './unset.mjs';
import { cloneDeep } from '../../object/cloneDeep.mjs';

function omit(obj, ...keysArr) {
    if (obj == null) {
        return {};
    }
    const result = cloneDeep(obj);
    for (let i = 0; i < keysArr.length; i++) {
        let keys = keysArr[i];
        switch (typeof keys) {
            case 'object': {
                if (!Array.isArray(keys)) {
                    keys = Array.from(keys);
                }
                for (let j = 0; j < keys.length; j++) {
                    const key = keys[j];
                    unset(result, key);
                }
                break;
            }
            case 'string':
            case 'symbol':
            case 'number': {
                unset(result, keys);
                break;
            }
        }
    }
    return result;
}

export { omit };
