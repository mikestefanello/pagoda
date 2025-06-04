/**
 * Removes all specified values from an array using an iteratee function.
 *
 * This function changes `arr` in place.
 * If you want to remove values without modifying the original array, use `differenceBy`.
 *
 * @template T
 * @param {T[]} arr - The array to modify.
 * @param {ArrayLike<T>} valuesToRemove - The values to remove from the array.
 * @param {(value: T) => unknown} getValue - The iteratee invoked per element.
 * @returns {T[]} The modified array with the specified values removed.
 *
 * @example
 * const items = [{ value: 1 }, { value: 2 }, { value: 3 }, { value: 1 }];
 * const result = pullAllBy(items, [{ value: 1 }, { value: 3 }], obj => obj.value);
 * console.log(result); // [{ value: 2 }]
 */
declare function pullAllBy<T>(arr: T[], valuesToRemove: ArrayLike<T>, getValue: (value: T) => unknown): T[];
/**
 * Removes all specified values from an array using an iteratee function.
 *
 * This function changes `arr` in place.
 * If you want to remove values without modifying the original array, use `differenceBy`.
 *
 * @template T
 * @param {T[]} arr - The array to modify.
 * @param {ArrayLike<T>} valuesToRemove - The values to remove from the array.
 * @param {Partial<T>} getValue - The partial object to match against each element.
 * @returns {T[]} The modified array with the specified values removed.
 */
declare function pullAllBy<T>(arr: T[], valuesToRemove: ArrayLike<T>, getValue: Partial<T>): T[];
/**
 * Removes all specified values from an array using an iteratee function.
 *
 * This function changes `arr` in place.
 * If you want to remove values without modifying the original array, use `differenceBy`.
 *
 * @template T
 * @param {T[]} arr - The array to modify.
 * @param {ArrayLike<T>} valuesToRemove - The values to remove from the array.
 * @param {[keyof T, unknown]} getValue - The property-value pair to match against each element.
 * @returns {T[]} The modified array with the specified values removed.
 */
declare function pullAllBy<T>(arr: T[], valuesToRemove: ArrayLike<T>, getValue: [keyof T, unknown]): T[];
/**
 * Removes all specified values from an array using an iteratee function.
 *
 * This function changes `arr` in place.
 * If you want to remove values without modifying the original array, use `differenceBy`.
 *
 * @template T
 * @param {T[]} arr - The array to modify.
 * @param {ArrayLike<T>} valuesToRemove - The values to remove from the array.
 * @param {keyof T} getValue - The key of the property to match against each element.
 * @returns {T[]} The modified array with the specified values removed.
 *
 * @example
 * const items = [{ value: 1 }, { value: 2 }, { value: 3 }, { value: 1 }];
 * const result = pullAllBy(items, [{ value: 1 }, { value: 3 }], 'value');
 * console.log(result); // [{ value: 2 }]
 */
declare function pullAllBy<T>(arr: T[], valuesToRemove: ArrayLike<T>, getValue: keyof T): T[];

export { pullAllBy };
