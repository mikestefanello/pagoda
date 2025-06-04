type Iteratee<T> = PropertyKey | Partial<T> | ((value: T) => unknown);
/**
 * Computes the difference between an array and another array using an iteratee function.
 *
 * @template T1, T2
 * @param {ArrayLike<T1> | null | undefined} array - The primary array from which to derive the difference.
 * @param {ArrayLike<T2>} values - The array containing elements to be excluded from the primary array.
 * @param {Iteratee<T1 | T2>} iteratee - The iteratee invoked per element.
 * @returns {T1[]} A new array containing the elements that are present in the primary array but not in the values array.
 *
 * @example
 * const result = differenceBy([2.1, 1.2], [2.3, 3.4], Math.floor);
 * // result will be [1.2]
 *
 * @example
 * const result = differenceBy([{ x: 2 }, { x: 1 }], [{ x: 1 }], 'x');
 * // result will be [{ x: 2 }]
 */
declare function differenceBy<T1, T2>(array: ArrayLike<T1> | null | undefined, values: ArrayLike<T2>, iteratee: Iteratee<T1 | T2>): T1[];
/**
 * Computes the difference between an array and two other arrays using an iteratee function.
 *
 * @template T1, T2, T3
 * @param {ArrayLike<T1> | null | undefined} array - The primary array from which to derive the difference.
 * @param {ArrayLike<T2>} values1 - The first array containing elements to be excluded from the primary array.
 * @param {ArrayLike<T3>} values2 - The second array containing elements to be excluded from the primary array.
 * @param {Iteratee<T1 | T2 | T3>} iteratee - The iteratee invoked per element.
 * @returns {T1[]} A new array containing the elements that are present in the primary array but not in the values arrays.
 *
 * @example
 * const result = differenceBy([2.1, 1.2], [2.3, 3.4], [1.2], Math.floor);
 * // result will be []
 */
declare function differenceBy<T1, T2, T3>(array: ArrayLike<T1> | null | undefined, values1: ArrayLike<T2>, values2: ArrayLike<T3>, iteratee: Iteratee<T1 | T2 | T3>): T1[];
/**
 * Computes the difference between an array and three other arrays using an iteratee function.
 *
 * @template T1, T2, T3, T4
 * @param {ArrayLike<T1> | null | undefined} array - The primary array from which to derive the difference.
 * @param {ArrayLike<T2>} values1 - The first array containing elements to be excluded from the primary array.
 * @param {ArrayLike<T3>} values2 - The second array containing elements to be excluded from the primary array.
 * @param {ArrayLike<T4>} values3 - The third array containing elements to be excluded from the primary array.
 * @param {Iteratee<T1 | T2 | T3 | T4>} iteratee - The iteratee invoked per element.
 * @returns {T1[]} A new array containing the elements that are present in the primary array but not in the values arrays.
 *
 * @example
 * const result = differenceBy([2.1, 1.2], [2.3, 3.4], [1.2], [2.1], Math.floor);
 * // result will be []
 */
declare function differenceBy<T1, T2, T3, T4>(array: ArrayLike<T1> | null | undefined, values1: ArrayLike<T2>, values2: ArrayLike<T3>, values3: ArrayLike<T4>, iteratee: Iteratee<T1 | T2 | T3 | T4>): T1[];
/**
 * Computes the difference between an array and four other arrays using an iteratee function.
 *
 * @template T1, T2, T3, T4, T5
 * @param {ArrayLike<T1> | null | undefined} array - The primary array from which to derive the difference.
 * @param {ArrayLike<T2>} values1 - The first array containing elements to be excluded from the primary array.
 * @param {ArrayLike<T3>} values2 - The second array containing elements to be excluded from the primary array.
 * @param {ArrayLike<T4>} values3 - The third array containing elements to be excluded from the primary array.
 * @param {ArrayLike<T5>} values4 - The fourth array containing elements to be excluded from the primary array.
 * @param {Iteratee<T1 | T2 | T3 | T4 | T5>} iteratee - The iteratee invoked per element.
 * @returns {T1[]} A new array containing the elements that are present in the primary array but not in the values arrays.
 *
 * @example
 * const result = differenceBy([2.1, 1.2], [2.3, 3.4], [1.2], [2.1], [3.4], Math.floor);
 * // result will be []
 */
declare function differenceBy<T1, T2, T3, T4, T5>(array: ArrayLike<T1> | null | undefined, values1: ArrayLike<T2>, values2: ArrayLike<T3>, values3: ArrayLike<T4>, values4: ArrayLike<T5>, iteratee: Iteratee<T1 | T2 | T3 | T4 | T5>): T1[];
/**
 * Computes the difference between an array and multiple arrays using an iteratee function.
 *
 * @template T
 * @param {ArrayLike<T> | null | undefined} array - The primary array from which to derive the difference.
 * @param {...Array<ArrayLike<T>>} values - Multiple arrays containing elements to be excluded from the primary array.
 * @returns {T[]} A new array containing the elements that are present in the primary array but not in the values arrays.
 *
 * @example
 * const result = differenceBy([2.1, 1.2], [2.3, 3.4], [1.2], [2.1], [3.4], Math.floor);
 * // result will be []
 */
declare function differenceBy<T>(array: ArrayLike<T> | null | undefined, ...values: Array<ArrayLike<T>>): T[];

export { differenceBy };
