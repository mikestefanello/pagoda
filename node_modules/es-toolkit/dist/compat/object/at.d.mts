/**
 * Returns an array of values corresponding to `paths` of `object`.
 *
 * @template T - The type of the object.
 * @param {T} object - The object to iterate over.
 * @param {...(PropertyKey | PropertyKey[] | ArrayLike<PropertyKey>)} [paths] - The property paths to pick.
 * @returns {Array<unknown>} - Returns the picked values.
 *
 * @example
 * ```js
 * const object = { 'a': [{ 'b': { 'c': 3 } }, 4] };
 *
 * at(object, ['a[0].b.c', 'a[1]']);
 * // => [3, 4]
 * ```
 */
declare function at<T>(object: T, ...paths: Array<PropertyKey | PropertyKey[] | ArrayLike<PropertyKey>>): unknown[];

export { at };
