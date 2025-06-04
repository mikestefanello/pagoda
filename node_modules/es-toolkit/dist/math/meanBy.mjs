import { mean } from './mean.mjs';

function meanBy(items, getValue) {
    const nums = items.map(x => getValue(x));
    return mean(nums);
}

export { meanBy };
