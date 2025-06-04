/**
 * Rejects items from a array and returns an array of elements.
 *
 * @template T
 * @param {ArrayLike<T> | null | undefined} arr - The array to iterate over.
 * @param {(item: T, index: number, arr: T[]) => unknown} doesMatch - The function invoked per iteration.
 * @returns {T[]} - Returns a new array of elements that do not satisfy the given doesMatch function.
 *
 * @example
 * reject([1, 2, 3], n => n % 2 === 0)
 * // => [1, 3]
 */
declare function reject<T>(arr: ArrayLike<T> | null | undefined, doesMatch?: (item: T, index: number, arr: readonly T[]) => unknown): T[];
/**
 * Rejects elements in a arr that match the properties of the given partial object.
 *
 * @template T
 * @param {ArrayLike<T> | null | undefined} arr - The array to iterate over.
 * @param {Partial<T>} doesMatch - A partial object that specifies the properties to match.
 * @returns {T[]} - Returns a new array of elements that do not match the given properties.
 *
 * @example
 * const arr = [{ id: 1, name: 'Alice' }, { id: 2, name: 'Bob' }];
 * reject(arr, { name: 'Bob' });
 * // => [{ id: 1, name: 'Alice' }]
 */
declare function reject<T>(arr: ArrayLike<T> | null | undefined, doesMatch: Partial<T>): T[];
/**
 * Rejects elements in a arr that match the given key-value pair.
 *
 * @template T
 * @param {ArrayLike<T> | null | undefined} arr - The array to iterate over.
 * @param {[keyof T, unknown]} doesMatchProperty - The key-value pair to match.
 * @returns {T[]} - Returns a new array of elements that do not match the given key-value pair.
 *
 * @example
 * const arr = [{ id: 1, name: 'Alice' }, { id: 2, name: 'Bob' }];
 * reject(arr, ['name', 'Alice']);
 * // => [{ id: 2, name: 'Bob' }]
 */
declare function reject<T>(arr: ArrayLike<T> | null | undefined, doesMatchProperty: [keyof T, unknown]): T[];
/**
 * Rejects the arr, returning elements that do not contain the given property name.
 *
 * @template T
 * @param {ArrayLike<T> | null | undefined} arr - The array to iterate over.
 * @param {PropertyKey} propertyToCheck - The property name to check.
 * @returns {T[]} - Returns a new array of elements that do not have the given property name.
 *
 * @example
 * const arr = [{ id: 1, name: 'Alice' }, { id: 2, name: 'Bob' }, { id: 3, age: 28 }];
 * reject(arr, 'name');
 * // => [{ id: 3, age: 28 }]
 */
declare function reject<T>(arr: ArrayLike<T> | null | undefined, propertyToCheck: PropertyKey): T[];
/**
 * Rejects items from a object and returns an array of elements that match the given predicate function.
 *
 * @template T
 * @param {T | null | undefined} object - The object to iterate over.
 * @param {(value: T[keyof T], key: keyof T, object: T) => unknown} doesMatch - The function invoked per iteration.
 * @returns {T[]} - Returns a new array of elements that do not satisfy the given predicate function.
 *
 * @example
 * const obj = { item1: { a: 0 }, item2: { a: 1 }, item3: { a: 0 } }
 * reject(obj, value => value.a)
 * // => [{ a: 0 }, { a: 0 }]
 *
 * const obj = { a: 1, b: 2, c: 3 };
 * reject(obj, value => value > 2)
 * // => [1, 2]
 */
declare function reject<T extends Record<string, unknown>>(object: T | null | undefined, doesMatch: (value: T[keyof T], key: keyof T, object: T) => unknown): T[];
/**
 * Rejects elements in a object that match the properties of the given partial object.
 *
 * @template T
 * @param {T | null | undefined} object - The object to iterate over.
 * @param {Partial<T[keyof T]>} doesMatch - The partial object to match
 * @returns {T[]} - Returns a new array of elements that do not match the given properties.
 *
 * @example
 * const obj = { a: { id: 1, name: 'Alice' }, b: { id: 2, name: 'Bob' } };
 * reject(obj, { name: 'Bob' });
 * // => [{ id: 1, name: 'Alice' }]
 */
declare function reject<T extends Record<string, unknown>>(object: T | null | undefined, doesMatch: Partial<T[keyof T]>): T[];
/**
 * Rejects elements in a arr that match the given key-value pair.
 *
 * @template T
 * @param {T | null | undefined} object - The object to iterate over.
 * @param {[keyof T[keyof T], unknown]} doesMatchProperty - The key-value pair to match.
 * @returns {T[]} - Returns a new array of elements that do not match the given key-value pair.
 *
 * @example
 * const obj = { alice: { id: 1, name: 'Alice' }, bob: { id: 2, name: 'Bob' } };
 * reject(obj, ['name', 'Alice']);
 * // => [{ id: 2, name: 'Bob' }]
 */
declare function reject<T extends Record<string, unknown>>(object: T | null | undefined, doesMatchProperty: [keyof T[keyof T], unknown]): T[];
/**
 * Rejects the object, returning elements that do not contain the given property name.
 *
 * @template T
 * @param {T | null | undefined} object - The object to iterate over.
 * @param {PropertyKey} propertyToCheck - The property name to check.
 * @returns {T[]} - Returns a new array of elements that do not have the given property name.
 *
 * @example
 * const obj = { a: { id: 1, name: 'Alice' }, b: { id: 2, name: 'Bob' }, c: { id: 3, age: 28 } };
 * reject(obj, 'name');
 * // => [{ id: 3, age: 28 }]
 */
declare function reject<T extends Record<string, unknown>>(object: T | null | undefined, propertyToCheck: PropertyKey): T[];

export { reject };
