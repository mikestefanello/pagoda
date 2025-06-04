import { property } from './property.mjs';
import { identity } from '../../function/identity.mjs';
import { mapValues as mapValues$1 } from '../../object/mapValues.mjs';

function mapValues(object, getNewValue) {
    getNewValue = getNewValue ?? identity;
    switch (typeof getNewValue) {
        case 'string':
        case 'symbol':
        case 'number':
        case 'object': {
            return mapValues$1(object, property(getNewValue));
        }
        case 'function': {
            return mapValues$1(object, getNewValue);
        }
    }
}

export { mapValues };
