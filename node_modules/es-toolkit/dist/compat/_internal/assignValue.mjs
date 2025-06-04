import { eq } from '../util/eq.mjs';

const assignValue = (object, key, value) => {
    const objValue = object[key];
    if (!(Object.hasOwn(object, key) && eq(objValue, value)) || (value === undefined && !(key in object))) {
        object[key] = value;
    }
};

export { assignValue };
