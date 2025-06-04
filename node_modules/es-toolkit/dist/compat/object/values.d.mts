/**
 * Creates an array of the own enumerable property values of `object`.
 *
 * @param {Record<PropertyKey, T> | null | undefined} object The object to query.
 * @returns {T[]} Returns an array of property values.
 * @example
 * const object = { a: 1, b: 2 };
 * values(object); // => [1, 2]
 */
declare function values<T>(object: Record<PropertyKey, T> | null | undefined): T[];
/**
 * Creates an array of the values of an array or array-like object.
 *
 * @param {ArrayLike<T>} arr The array or array-like object to query.
 * @returns {T[]} Returns an array of values.
 * @example
 * const array = ['a', 'b'];
 * values(arrayLike); // => ['a', 'b']
 */
declare function values<T>(arr: ArrayLike<T>): T[];
/**
 * Creates an array of the own enumerable property values of `object`.
 *
 * @param {T | null | undefined} object The object to query.
 * @returns {Array<T[keyof T]>} Returns an array of property values.
 * @example
 * const obj = { x: 1, y: 2, z: 3 };
 * values(obj); // => [1, 2, 3]
 */
declare function values<T extends object>(object: T | null | undefined): Array<T[keyof T]>;

export { values };
