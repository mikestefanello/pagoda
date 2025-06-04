/**
 * Removes and returns elements from an array using a provided comparison function to determine which elements to remove.
 *
 * @template T
 * @param {T[]} array - The array to modify.
 * @param {T[]} values - The values to remove from the array.
 * @param {(a: T, b: T) => boolean} comparator - The function to compare elements of `array` with elements of `values`. Should return `true` if the two elements are considered equal.
 * @returns {T[]} - The array with specified values removed.
 *
 * @example
 * const array = [{ x: 1, y: 2 }, { x: 3, y: 4 }, { x: 5, y: 6 }];
 * const valuesToRemove = [{ x: 3, y: 4 }];
 *
 * const result = pullAllWith(array, valuesToRemove, isEqual);
 *
 * console.log(result); // [{ x: 1, y: 2 }, { x: 5, y: 6 }]
 */
declare function pullAllWith<T>(array: T[], values?: T[], comparator?: (a: T, b: T) => boolean): T[];
declare function pullAllWith<T>(array: ArrayLike<T>, values?: ArrayLike<T>, comparator?: (a: T, b: T) => boolean): ArrayLike<T>;
/**
 * Removes and returns elements from an array using a provided comparison function to determine which elements to remove.
 *
 * @template T1
 * @template T2
 * @param {T1[]} array - The array to modify.
 * @param {ArrayLike<T2>} values - The values to remove from the array.
 * @param {(a: T1, b: T2) => boolean} comparator - The function to compare elements of `array` with elements of `values`. Should return `true` if the two elements are considered equal.
 * @returns {T1[]} - The array with specified values removed.
 *
 * @example
 * const array = [{ x: 1, y: 2 }, { x: 3, y: 4 }, { x: 5, y: 6 }];
 * const valuesToRemove = [{ x: 3, y: 4 }];
 *
 * const result = pullAllWith(array, valuesToRemove, isEqual);
 *
 * console.log(result); // [{ x: 1, y: 2 }, { x: 5, y: 6 }]
 */
declare function pullAllWith<T1, T2>(array: T1[], values: ArrayLike<T2>, comparator: (a: T1, b: T2) => boolean): T1[];
/**
 * Removes and returns elements from an array using a provided comparison function to determine which elements to remove.
 *
 * @template T1
 * @template T2
 * @param {T1[]} array - The array to modify.
 * @param {ArrayLike<T2>} values - The values to remove from the array.
 * @param {(a: T1, b: T2) => boolean} comparator - The function to compare elements of `array` with elements of `values`. Should return `true` if the two elements are considered equal.
 * @returns {ArrayLike<T1>} - The array with specified values removed.
 *
 * @example
 * const array = [{ x: 1, y: 2 }, { x: 3, y: 4 }, { x: 5, y: 6 }];
 * const valuesToRemove = [{ x: 3, y: 4 }];
 *
 * const result = pullAllWith(array, valuesToRemove, isEqual);
 *
 * console.log(result); // [{ x: 1, y: 2 }, { x: 5, y: 6 }]
 */
declare function pullAllWith<T1, T2>(array: ArrayLike<T1>, values: ArrayLike<T2>, comparator: (a: T1, b: T2) => boolean): ArrayLike<T1>;

export { pullAllWith };
