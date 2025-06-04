/**
 * Removes elements from an array at specified indices and returns the removed elements.
 *
 * @template T
 * @param {T[]} array - The array from which elements will be removed.
 * @param {Array<number | readonly number[] | string | readonly string[]>} indicesToRemove - An array of indices specifying the positions of elements to remove.
 * @returns {T[]} An array containing the elements that were removed from the original array.
 *
 * @example
 * const numbers = [10, 20, 30, 40, 50];
 * const removed = pullAt(numbers, [1, 3, 4]);
 * console.log(removed); // [20, 40, 50]
 * console.log(numbers); // [10, 30]
 */
declare function pullAt<T>(array: T[], ...indicesToRemove: Array<number | readonly number[] | string | readonly string[]>): T[];
/**
 * Removes elements from an array at specified indices and returns the removed elements.
 *
 * @template T
 * @param {ArrayLike<T>} array - The array from which elements will be removed.
 * @param {Array<number | readonly number[] | string | readonly string[]>} indicesToRemove - An array of indices specifying the positions of elements to remove.
 * @returns {ArrayLike<T>} An array containing the elements that were removed from the original array.
 *
 * @example
 * const numbers = [10, 20, 30, 40, 50];
 * const removed = pullAt(numbers, [1, 3, 4]);
 * console.log(removed); // [20, 40, 50]
 * console.log(numbers); // [10, 30]
 */
declare function pullAt<T>(array: ArrayLike<T>, ...indicesToRemove: Array<number | readonly number[] | string | readonly string[]>): ArrayLike<T>;

export { pullAt };
