/**
 * Checks if there is an element in an array that is truthy.
 *
 * @template T
 * @param {ArrayLike<T> | null | undefined} arr The array to iterate over.
 * @returns {boolean} Returns `true` if any element is truthy, else `false`.
 *
 * @example
 * some([1, 2, 3, 4]);
 * // => true
 */
declare function some<T>(arr: ArrayLike<T> | null | undefined): boolean;
/**
 * Checks if there is an element in an array that matches the given predicate function.
 *
 * @template T
 * @param {ArrayLike<T> | null | undefined} arr The array to iterate over.
 * @param {(item: T, index: number, arr: readonly T[]) => unknown} predicate The function invoked per iteration.
 * @returns {boolean} Returns `true` if any element passes the predicate check, else `false`.
 *
 * @example
 * some([1, 2, 3, 4], n => n % 2 === 0);
 * // => true
 */
declare function some<T>(arr: ArrayLike<T> | null | undefined, predicate: (item: T, index: number, arr: readonly T[]) => unknown): boolean;
/**
 * Checks if there is an element in an array that matches the given key-value pair.
 *
 * @template T
 * @param {ArrayLike<T> | null | undefined} arr The array to iterate over.
 * @param {[keyof T, unknown]} predicate The key-value pair to match.
 * @returns {boolean} Returns `true` if any element passes the predicate check, else `false`.
 *
 * @example
 * some([{ a: 1 }, { a: 2 }, { a: 3 }], ['a', 2]);
 * // => true
 */
declare function some<T>(arr: ArrayLike<T> | null | undefined, predicate: [keyof T, unknown]): boolean;
/**
 * Checks if there is an element in an array that has a truthy value for the given property name.
 *
 * @template T
 * @param {ArrayLike<T> | null | undefined} arr The array to iterate over.
 * @param {PropertyKey} propertyToCheck The property name to check.
 * @returns {boolean} Returns `true` if any element has a truthy value for the property, else `false`.
 *
 * @example
 * some([{ a: 1 }, { a: 2 }, { a: 3 }], 'a');
 * // => true
 */
declare function some<T>(arr: ArrayLike<T> | null | undefined, propertyToCheck: PropertyKey): boolean;
/**
 * Checks if there is an element in an array that matches the given partial object.
 *
 * @template T
 * @param {ArrayLike<T> | null | undefined} arr The array to iterate over.
 * @param {Partial<T>} doesMatch The partial object to match.
 * @returns {boolean} Returns `true` if any element matches the partial object, else `false`.
 *
 * @example
 * some([{ a: 1 }, { a: 2 }, { a: 3 }], { a: 2 });
 * // => true
 */
declare function some<T>(arr: ArrayLike<T> | null | undefined, doesMatch: Partial<T>): boolean;
/**
 * Checks if there is an element in an object that matches the given predicate function.
 *
 * @template T
 * @param {T | null | undefined} object The object to iterate over.
 * @returns {boolean} Returns `true` if any element is truthy, else `false`.
 *
 * @example
 * some({ a: 1, b: 2, c: 3 });
 * // => true
 */
declare function some<T extends Record<string, unknown>>(object: T | null | undefined): boolean;
/**
 * Checks if there is an element in an object that matches the given predicate function.
 *
 * @template T
 * @param {T | null | undefined} object The object to iterate over.
 * @param {(value: T[keyof T], key: keyof T, object: T) => unknown} doesMatch A function that takes an value, its key, and the object, and returns a truthy value if the item matches the criteria.
 * @returns {boolean} Returns `true` if any element passes the predicate check, else `false`.
 *
 * @example
 * some({ a: 1, b: 2, c: 3 }, n => n % 2 === 0);
 * // => true
 */
declare function some<T extends Record<string, unknown>>(object: T | null | undefined, doesMatch: (value: T[keyof T], key: keyof T, object: T) => unknown): boolean;
/**
 * Checks if there is an element in an object that matches the given partial value.
 *
 * @template T
 * @param {T | null | undefined} object The object to iterate over.
 * @param {Partial<T[keyof T]>} doesMatch A partial value to match against the values of the object.
 * @returns {boolean} Returns `true` if any element passes the predicate check, else `false`.
 *
 * @example
 * some({ a: { id: 1, name: 'Alice' }, b: { id: 2, name: 'Bob' } }, { name: 'Bob' });
 * // => true
 */
declare function some<T extends Record<string, unknown>>(object: T | null | undefined, doesMatch: Partial<T[keyof T]>): boolean;
/**
 * Checks if there is an element in an object that matches a property with a specific value.
 *
 * @template T
 * @param {T | null | undefined} object The object to iterate over.
 * @param {[keyof T, unknown]} doesMatchProperty An array where the first element is the property key and the second element is the value to match.
 * @returns {boolean} Returns `true` if any element passes the predicate check, else `false`.
 *
 * @example
 * const obj = { alice: { id: 1, name: 'Alice' }, bob: { id: 2, name: 'Bob' } };
 * const result = some(obj, ['name', 'Alice']);
 * // => true
 */
declare function some<T extends Record<string, unknown>>(object: T | null | undefined, doesMatchProperty: [keyof T[keyof T], unknown]): boolean;
/**
 * Checks if there is an element in an object that has a specific property, where the property name is provided as a PropertyKey.
 *
 * @template T
 * @param {T | null | undefined} object The object to iterate over.
 * @param {PropertyKey} propertyToCheck The property name to check.
 * @returns {boolean} Returns `true` if any element passes the predicate check, else `false`.
 *
 * @example
 * const obj = { alice: { id: 1, name: 'Alice' }, bob: { id: 2, name: 'Bob' } };
 * const result = some(obj, 'name');
 * // => true
 */
declare function some<T extends Record<string, unknown>>(object: T | null | undefined, propertyToCheck: PropertyKey): boolean;

export { some };
