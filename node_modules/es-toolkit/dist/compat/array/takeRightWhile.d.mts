/**
 * Creates a slice of array.
 *
 * If the array is `null` or `undefined`, returns an empty array.
 *
 * @template T
 * @param {ArrayLike<T> | null | undefined} array - The array to process.
 * @returns {T[]} - A slice of the array or an empty array if `array` is `null` or `undefined`.
 *
 * @example
 * const items = [1, 2, 3];
 * const result = takeRightWhile(items);
 * console.log(result); // [1, 2, 3]
 *
 * const result2 = takeRightWhile(null);
 * console.log(result2); // []
 */
declare function takeRightWhile<T>(array: ArrayLike<T> | null | undefined): T[];
/**
 * Creates a slice of array with elements taken from the end until the predicate function returns falsey.
 *
 * If the array is `null` or `undefined`, returns an empty array.
 *
 * @template T
 * @param {ArrayLike<T> | null | undefined} array - The array to process.
 * @param {(item: T, index: number, array: T[]) => unknown} predicate - A function invoked per iteration. Returns a truthy value to continue taking elements.
 * @returns {T[]} - A slice of the array with elements taken from the end or an empty array if `array` is `null` or `undefined`.
 *
 * @example
 * const items = [1, 2, 3, 4, 5];
 * const result = takeRightWhile(items, (item) => item > 3);
 * console.log(result); // [4, 5]
 */
declare function takeRightWhile<T>(array: ArrayLike<T> | null | undefined, predicate: (item: T, index: number, array: T[]) => unknown): T[];
/**
 * Creates a slice of array with elements taken from the end until the element does not match the given object.
 *
 * If the array is `null` or `undefined`, returns an empty array.
 *
 * @template T
 * @param {ArrayLike<T> | null | undefined} array - The array to process.
 * @param {Partial<T>} matches - A partial object that specifies the properties to match.
 * @returns {T[]} - A slice of the array with elements taken from the end or an empty array if `array` is `null` or `undefined`.
 *
 * @example
 * const items = [{ id: 10 }, { id: 20 }, { id: 30 }];
 * const result = takeRightWhile(items, { id: 30 });
 * console.log(result); // [{ id: 30 }]
 */
declare function takeRightWhile<T>(array: ArrayLike<T> | null | undefined, matches: Partial<T>): T[];
/**
 * Creates a slice of array with elements taken from the end until the element does not match the given property key and value.
 *
 * If the array is `null` or `undefined`, returns an empty array.
 *
 * @template T
 * @param {ArrayLike<T> | null | undefined} array - The array to process.
 * @param {[keyof T, unknown]} matchesProperty - An array where the first element is the property key and the second element is the value to match.
 * @returns {T[]} - A slice of the array with elements taken from the end or an empty array if `array` is `null` or `undefined`.
 *
 * @example
 * const items = [{ name: 'Alice' }, { name: 'Bob' }, { name: 'Alice' }];
 * const result = takeRightWhile(items, ['name', 'Alice']);
 * console.log(result); // [{ name: 'Alice' }]
 */
declare function takeRightWhile<T>(array: ArrayLike<T> | null | undefined, matchesProperty: [keyof T, unknown]): T[];
/**
 * Creates a slice of array with elements taken from the end until the element does not have a truthy value for the given property key.
 *
 * If the array is `null` or `undefined`, returns an empty array.
 *
 * @template T
 * @param {ArrayLike<T> | null | undefined} array - The array to process.
 * @param {PropertyKey} property - A property key. Elements are included if they have a truthy value for this key.
 * @returns {T[]} - A slice of the array with elements taken from the end or an empty array if `array` is `null` or `undefined`.
 *
 * @example
 * const items = [{ valid: false }, { valid: true }, { valid: true }];
 * const result = takeRightWhile(items, 'valid');
 * console.log(result); // [{ valid: true }, { valid: true }]
 */
declare function takeRightWhile<T>(array: ArrayLike<T> | null | undefined, property: PropertyKey): T[];

export { takeRightWhile };
