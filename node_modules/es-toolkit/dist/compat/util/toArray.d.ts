/**
 * Converts a value to an array.
 *
 * @param {unknown} value - The value to convert.
 * @returns {any[]} Returns the converted array.
 *
 * @example
 * toArray({ 'a': 1, 'b': 2 }) // => returns [1,2]
 * toArray('abc') // => returns ['a', 'b', 'c']
 * toArray(1) // => returns []
 * toArray(null) // => returns []
 */
declare function toArray(value?: unknown): any[];

export { toArray };
