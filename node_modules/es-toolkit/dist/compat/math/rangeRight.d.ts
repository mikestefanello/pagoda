/**
 * Returns an array of numbers from `end` (exclusive) to `0` (inclusive), decrementing by `1`.
 *
 * @param {number} end - The end number of the range (exclusive).
 * @returns {number[]} An array of numbers from `end` (exclusive) to `0` (inclusive) with a step of `1`.
 *
 * @example
 * // Returns [3, 2, 1, 0]
 * rangeRight(4);
 */
declare function rangeRight(end: number): number[];
/**
 * Returns an array of numbers from `end` (exclusive) to `start` (inclusive), decrementing by `1`.
 *
 * @param {number} start - The starting number of the range (inclusive).
 * @param {number} end - The end number of the range (exclusive).
 * @returns {number[]} An array of numbers from `end` (exclusive) to `start` (inclusive) with a step of `1`.
 *
 * @example
 * // Returns [3, 2, 1]
 * rangeRight(1, 4);
 */
declare function rangeRight(start: number, end: number): number[];
/**
 * Returns an array of numbers from `end` (exclusive) to `start` (inclusive), decrementing by `step`.
 *
 * @param {number} start - The starting number of the range (inclusive).
 * @param {number} end - The end number of the range (exclusive).
 * @param {number} step - The step value for the range.
 * @returns {number[]} An array of numbers from `end` (exclusive) to `start` (inclusive) with the specified `step`.
 *
 * @example
 * // Returns [15, 10, 5, 0]
 * rangeRight(0, 20, 5);
 */
declare function rangeRight(start: number, end: number, step: number): number[];
/**
 * Enables use as an iteratee for methods like `_.map`.
 *
 * @param {number} end - The current iteratee value.
 * @param {PropertyKey} index - The iteration index.
 * @param {object} guard - The iteratee object.
 * @returns {number[]} An array of numbers from `start` (inclusive) to `end` (exclusive) with the specified `step`.
 */
declare function rangeRight(end: number, index: PropertyKey, guard: object): number[];

export { rangeRight };
