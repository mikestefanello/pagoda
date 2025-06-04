/**
 * Returns a new array containing the first `count` elements from the input array `arr`.
 * If `count` is greater than the length of `arr`, the entire array is returned.
 *
 * @template T - Type of elements in the input array.
 *
 * @param {ArrayLike<T> | null | undefined} arr - The array to take elements from.
 * @param {number} [count=1] - The number of elements to take.
 * @param {unknown} [guard] - Enables use as an iteratee for methods like `_.map`.
 * @returns {T[]} A new array containing the first `count` elements from `arr`.
 *
 * @example
 * // Returns [1, 2, 3]
 * take([1, 2, 3, 4, 5], 3);
 *
 * @example
 * // Returns ['a', 'b']
 * take(['a', 'b', 'c'], 2);
 *
 * @example
 * // Returns [1, 2, 3]
 * take([1, 2, 3], 5);
 */
declare function take<T>(arr: ArrayLike<T> | null | undefined, count?: number, guard?: unknown): T[];

export { take };
