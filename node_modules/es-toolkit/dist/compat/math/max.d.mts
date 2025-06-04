/**
 * Finds the element in an array that has the maximum value.
 *
 * @template T - The type of elements in the array.
 * @param {[T, ...T[]]} items - The array of elements to search.
 * @returns {T} - The element with the maximum value.
 * @example
 * // Returns 9
 * max([3, 1, 4, 1, 5, 9]);
 *
 * @example
 * // Returns 8
 * max([0, -3, 2, 8, 7]);
 */
declare function max<T>(items: readonly [T, ...T[]]): T;
/**
 * Finds the element in an array that has the maximum value.
 * Returns undefined when no arguments are provided.
 *
 * @returns {undefined} - Returns `undefined` when the function is called with no arguments.
 */
declare function max(): undefined;
/**
 * Finds the element in an array that has the maximum value.
 *
 * @template T - The type of elements in the array.
 * @param {ArrayLike<T> | null | undefined} [items] - The array of elements to search.
 * @returns {T | undefined} - The element with the maximum value, or `undefined` if the array is empty, `null`, or `undefined`.
 */
declare function max<T>(items?: ArrayLike<T> | null | undefined): T | undefined;

export { max };
