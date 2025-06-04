import { property } from './property.mjs';
import { identity } from '../../function/identity.mjs';
import { mapKeys as mapKeys$1 } from '../../object/mapKeys.mjs';

function mapKeys(object, getNewKey) {
    getNewKey = getNewKey ?? identity;
    switch (typeof getNewKey) {
        case 'string':
        case 'symbol':
        case 'number':
        case 'object': {
            return mapKeys$1(object, property(getNewKey));
        }
        case 'function': {
            return mapKeys$1(object, getNewKey);
        }
    }
}

export { mapKeys };
