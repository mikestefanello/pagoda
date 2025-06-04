/**
 * Returns a random element from an array.
 *
 * @template T
 * @param {T[]} array - The array to sample from.
 * @returns {T | undefined} A random element from the array, or `undefined` if the array is empty.
 *
 * @example
 * const array = [1, 2, 3];
 * const result = sample(array);
 * console.log(result); // Output: 1, 2, or 3 (randomly selected)
 */
declare function sample<T>(array: readonly T[]): T | undefined;
/**
 * Returns a random character from a string.
 *
 * @param {string} str - The string to sample from.
 * @returns {string | undefined} A random character from the string, or `undefined` if the string is empty.
 *
 * @example
 * const str = "hello";
 * const result = sample(str);
 * console.log(result); // Output: 'h', 'e', 'l', 'l', or 'o' (randomly selected)
 */
declare function sample(str: string): string | undefined;
/**
 * Returns a random element from an array.
 *
 * @template T
 * @param {ArrayLike<T>} array - The array-like object to sample from.
 * @returns {T | undefined} A random element from the array, or `undefined` if the array is empty.
 *
 * @example
 * const arrayLike: ArrayLike<string> = { 0: 'a', 1: 'b', 2: 'c', length: 3 };
 * const result = sample(arrayLike);
 * console.log(result); // Output: 'a', 'b', or 'c' (randomly selected)
 */
declare function sample<T>(array: ArrayLike<T>): T | undefined;
/**
 * Returns a random value from an object.
 *
 * @template T - The type of values in the object.
 * @param {Record<string, T>} obj - The object to sample from.
 * @returns {T | undefined} A random value from the object, or `undefined` if the object is empty.
 *
 * @example
 * const obj = { a: 1, b: 2, c: 3 };
 * const result = sample(obj);
 * console.log(result); // Output: 1, 2, or 3 (randomly selected)
 */
declare function sample<T>(obj: Record<string, T>): T | undefined;
/**
 * Returns a random element from an array-like object or a regular object.
 *
 * This function takes an array-like object (such as an array or string) or a regular object,
 * and returns a randomly selected element or value. If the collection is empty or invalid, it returns `undefined`.
 *
 * @template T - The type of elements in the collection.
 * @param {ArrayLike<T> | Record<string, T>} collection - The collection to sample from.
 * @returns {T | string | undefined} A random element from the collection, or `undefined` if the collection is empty or invalid.
 *
 * @example
 * // Array example
 * const array = [1, 2, 3];
 * const result = sample(array);
 * console.log(result); // Output: 1, 2, or 3 (randomly selected)
 *
 * // String example
 * const str = 'abc';
 * const result2 = sample(str);
 * console.log(result2); // Output: 'a', 'b', or 'c' (randomly selected)
 *
 * // Object example
 * const obj = { a: 1, b: 2, c: 3 };
 * const result3 = sample(obj);
 * console.log(result3); // Output: 1, 2, or 3 (randomly selected)
 */
declare function sample<T>(collection: ArrayLike<T> | Record<string, T>): T | string | undefined;

export { sample };
