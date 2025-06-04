/**
 * Checks if a given path exists within an object.
 *
 * You can provide the path as a single property key, an array of property keys,
 * or a string representing a deep path.
 *
 * If the path is an index and the object is an array or an arguments object, the function will verify
 * if the index is valid and within the bounds of the array or arguments object, even if the array or
 * arguments object is sparse (i.e., not all indexes are defined).
 *
 * @param {object} object - The object to query.
 * @param {PropertyKey | PropertyKey[]} path - The path to check. This can be a single property key,
 *        an array of property keys, or a string representing a deep path.
 * @returns {boolean} Returns `true` if the path exists in the object, `false` otherwise.
 *
 * @example
 *
 * const obj = { a: { b: { c: 3 } } };
 *
 * has(obj, 'a'); // true
 * has(obj, ['a', 'b']); // true
 * has(obj, ['a', 'b', 'c']); // true
 * has(obj, 'a.b.c'); // true
 * has(obj, 'a.b.d'); // false
 * has(obj, ['a', 'b', 'c', 'd']); // false
 * has([], 0); // false
 * has([1, 2, 3], 2); // true
 * has([1, 2, 3], 5); // false
 */
declare function has(object: unknown, path: PropertyKey | readonly PropertyKey[]): boolean;

export { has };
