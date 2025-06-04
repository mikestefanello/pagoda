/**
 * Maps each element of an array based on a provided key-generating function.
 *
 * This function takes an array and a function that generates a key from each element. It returns
 * an object where the keys are the generated keys and the values are the corresponding elements.
 * If there are multiple elements generating the same key, the last element among them is used
 * as the value.
 *
 * @template T - The type of elements in the array.
 * @template K - The type of keys.
 * @param {T[]} arr - The array of elements to be mapped.
 * @param {(item: T) => K} getKeyFromItem - A function that generates a key from an element.
 * @returns {Record<K, T>} An object where keys are mapped to each element of an array.
 *
 * @example
 * const array = [
 *   { category: 'fruit', name: 'apple' },
 *   { category: 'fruit', name: 'banana' },
 *   { category: 'vegetable', name: 'carrot' }
 * ];
 * const result = keyBy(array, item => item.category);
 * // result will be:
 * // {
 * //   fruit: { category: 'fruit', name: 'banana' },
 * //   vegetable: { category: 'vegetable', name: 'carrot' }
 * // }
 */
declare function groupBy<T, K extends PropertyKey>(source: ArrayLike<T> | null | undefined, getKeyFromItem?: ((item: T, index: number, arr: any) => unknown) | Partial<T> | [keyof T, unknown] | PropertyKey | null): Record<K, T[]>;
/**
 * Groups the elements of an object based on a provided key-generating function.
 *
 * This function takes an object and a function that generates a key from each value. It returns
 * an object where the keys are the generated keys and the values are arrays of elements that share
 * the same key.
 *
 * @template T - The type of values in the object.
 * @template K - The type of keys.
 * @param {Record<any, T> | null | undefined} source - The object to group.
 * @param {Function | PropertyKey | Array | Object} [getKeyFromItem] - The iteratee to transform keys.
 *   - If a function is provided, it's invoked for each element in the collection.
 *   - If a property name (string) is provided, that property of each element is used as the key.
 *   - If a property-value pair (array) is provided, elements with matching property values are used.
 *   - If a partial object is provided, elements with matching properties are used.
 * @returns {Record<K, T>} An object where each key is associated with an array of elements that
 * share that key.
 *
 * @example
 * // Using an object
 * const obj = { a: { category: 'fruit' }, b: { category: 'vegetable' }, c: { category: 'fruit' } };
 * const result = groupBy(obj, 'category');
 * // => { fruit: [{ category: 'fruit' }, { category: 'fruit' }], vegetable: [{ category: 'vegetable' }] }
 */
declare function groupBy<T, K extends PropertyKey>(source: Record<any, T> | null | undefined, getKeyFromItem?: ((item: T, index: number, arr: any) => unknown) | Partial<T> | [keyof T, unknown] | PropertyKey | null): Record<K, T[]>;

export { groupBy };
