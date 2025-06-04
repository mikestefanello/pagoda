/**
 * Returns a sample element array of a specified `size`.
 *
 * This function takes an collection and a number, and returns an array containing the sampled elements using Floyd's algorithm.
 *
 * {@link https://www.nowherenearithaca.com/2013/05/robert-floyds-tiny-and-beautiful.html Floyd's algorithm}
 *
 * @template T - The type of elements in the collection.
 * @param {Record<string, T> | Record<number, T> | null | undefined} collection - The collection to sample from.
 * @param {number} size - The size of sample.
 * @returns {T[]} A new array with sample size applied.
 *
 * @example
 * const result = sampleSize([1, 2, 3], 2)
 * // result will be an array containing two of the elements from the collection.
 * // [1, 2] or [1, 3] or [2, 3]
 */
declare function sampleSize<T>(collection: Record<string, T> | Record<number, T> | null | undefined, size?: number): T[];
/**
 * Returns a sample element array of a specified `size`.
 *
 * This function takes an collection and a number, and returns an array containing the sampled elements using Floyd's algorithm.
 *
 * {@link https://www.nowherenearithaca.com/2013/05/robert-floyds-tiny-and-beautiful.html Floyd's algorithm}
 *
 * @template T - The type of the collection.
 * @param {T | null | undefined} collection - The collection to sample from.
 * @param {number} size - The size of sample.
 * @returns {T[]} A new array with sample size applied.
 *
 * @example
 * const result = sampleSize({ a: 1, b: 2, c: 3 }, 2)
 * // result will be an array containing two of the values from the collection.
 * // [1, 2] or [1, 3] or [2, 3]
 */
declare function sampleSize<T extends object>(collection: T | null | undefined, size?: number): Array<T[keyof T]>;

export { sampleSize };
