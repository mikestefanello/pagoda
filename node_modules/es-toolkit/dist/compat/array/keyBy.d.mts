/**
 * Maps each element of an array based on a provided key-generating function.
 *
 * This function takes an array and a function that generates a key from each element. It returns
 * an object where the keys are the generated keys and the values are the corresponding elements.
 * If there are multiple elements generating the same key, the last element among them is used
 * as the value.
 *
 * @param {ArrayLike<T> | null | undefined} collection - The collection to iterate over.
 * @param {Function | PropertyKey | Array | Object} [iteratee] - The iteratee to transform keys.
 *   - If a function is provided, it's invoked for each element in the collection.
 *   - If a property name (string) is provided, that property of each element is used as the key.
 *   - If a property-value pair (array) is provided, elements with matching property values are used.
 *   - If a partial object is provided, elements with matching properties are used.
 *   - If omitted, the identity function is used.
 * @returns {Object} Returns the composed aggregate object.
 *
 * @example
 * // Using an array of objects
 * keyBy([{ id: 'a' }, { id: 'b' }], 'id');
 * // => { a: { id: 'a' }, b: { id: 'b' } }
 *
 * @example
 * // Using a function iteratee
 * keyBy(['a', 'b', 'c'], val => val.toUpperCase());
 * // => { A: 'a', B: 'b', C: 'c' }
 */
declare function keyBy<T>(collection: ArrayLike<T> | null | undefined, iteratee?: ((value: T) => unknown) | PropertyKey | [keyof T, unknown] | Partial<T> | null): Record<string, T>;
/**
 * Maps each value of an object based on a provided key-generating function.
 *
 * This function takes an object and a function that generates a key from each value. It returns
 * an object where the keys are the generated keys and the values are the corresponding values.
 * If there are multiple values generating the same key, the last value among them is used
 * as the value.
 *
 * @example
 * // Using an object
 * keyBy({ a: { id: 1 }, b: { id: 2 } }, 'id');
 * // => { '1': { id: 1 }, '2': { id: 2 } }
 *
 * @param collection - The object to iterate over.
 * @param iteratee - The function or shorthand to generate the key.
 * @returns An object composed of keys generated from the iteratee.
 */
declare function keyBy<T extends object>(collection: T | null | undefined, iteratee?: ((value: T[keyof T]) => unknown) | PropertyKey | [keyof T[keyof T], unknown] | Partial<T[keyof T]> | null): Record<string, T[keyof T]>;

export { keyBy };
