import { clamp as clamp$1 } from '../../math/clamp.mjs';

function clamp(value, bound1, bound2) {
    if (Number.isNaN(bound1)) {
        bound1 = 0;
    }
    if (Number.isNaN(bound2)) {
        bound2 = 0;
    }
    return clamp$1(value, bound1, bound2);
}

export { clamp };
