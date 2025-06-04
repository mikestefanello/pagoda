/**
 * Returns the intersection of multiple arrays after applying the iteratee function to their elements.
 *
 * This function takes multiple arrays and an iteratee function (or property key) to
 * compare the elements after transforming them. It returns a new array containing the elements from
 * the first array that are present in all subsequent arrays after applying the iteratee to each element.
 *
 * @template T1, T2
 * @param {ArrayLike<T1> | null | undefined} array - The first array to compare.
 * @param {ArrayLike<T2>} values - The second array to compare.
 * @param {(value: T1 | T2) => unknown | string} iteratee - The iteratee invoked on each element
 *  for comparison. It can also be a property key to compare based on that property.
 * @returns {T1[]} A new array containing the elements from the first array that are present
 *  in all subsequent arrays after applying the iteratee.
 *
 * @example
 * const array1 = [{ x: 1 }, { x: 2 }, { x: 3 }];
 * const array2 = [{ x: 2 }, { x: 3 }, { x: 4 }];
 * const result = intersectionBy(array1, array2, 'x');
 * // result will be [{ x: 2 }, { x: 3 }] since these elements have the same `x` property.
 *
 * const array1 = [{ x: 1 }, { x: 2 }, { x: 3 }];
 * const array2 = [{ x: 2 }, { x: 3 }, { x: 4 }];
 * const result = intersectionBy(array1, array2, value => value.x);
 * // result will be [{ x: 2 }, { x: 3 }] since these elements have the same `x` property.
 */
declare function intersectionBy<T1, T2>(array: ArrayLike<T1> | null | undefined, values: ArrayLike<T2>, iteratee: ((value: T1 | T2) => unknown) | string): T1[];
/**
 * Returns the intersection of multiple arrays after applying the iteratee function to their elements.
 *
 * This function takes multiple arrays and an iteratee function (or property key) to
 * compare the elements after transforming them. It returns a new array containing the elements from
 * the first array that are present in all subsequent arrays after applying the iteratee to each element.
 *
 * @template T1, T2, T3
 * @param {ArrayLike<T1> | null | undefined} array - The first array to compare.
 * @param {ArrayLike<T2>} values1 - The second array to compare.
 * @param {ArrayLike<T3>} values2 - The third array to compare.
 * @param {(value: T1 | T2 | T3) => unknown | string} iteratee - The iteratee invoked on each element
 *  for comparison. It can also be a property key to compare based on that property.
 * @returns {T1[]} A new array containing the elements from the first array that are present
 *  in all subsequent arrays after applying the iteratee.
 *
 * @example
 * const array1 = [{ x: 1 }, { x: 2 }, { x: 3 }];
 * const array2 = [{ x: 2 }, { x: 3 }];
 * const array3 = [{ x: 3 }];
 * const result = intersectionBy(array1, array2, array3, 'x');
 * // result will be [{ x: 3 }] since this element has the same `x` property in all arrays.
 *
 * const array1 = [{ x: 1 }, { x: 2 }, { x: 3 }];
 * const array2 = [{ x: 2 }, { x: 3 }];
 * const array3 = [{ x: 3 }];
 * const result = intersectionBy(array1, array2, array3, value => value.x);
 * // result will be [{ x: 3 }] since this element has the same `x` property in all arrays.
 */
declare function intersectionBy<T1, T2, T3>(array: ArrayLike<T1> | null | undefined, values1: ArrayLike<T2>, values2: ArrayLike<T3>, iteratee: ((value: T1 | T2 | T3) => unknown) | string): T1[];
/**
 * Returns the intersection of multiple arrays after applying the iteratee function to their elements.
 *
 * This function takes multiple arrays and an iteratee function (or property key) to
 * compare the elements after transforming them. It returns a new array containing the elements from
 * the first array that are present in all subsequent arrays after applying the iteratee to each element.
 *
 * @template T1, T2, T3, T4
 * @param {ArrayLike<T1> | null | undefined} array - The first array to compare.
 * @param {ArrayLike<T2>} values1 - The second array to compare.
 * @param {ArrayLike<T3>} values2 - The third array to compare.
 * @param {...(ArrayLike<T4> | ((value: T1 | T2 | T3 | T4) => unknown) | string)} values - Additional arrays to compare, or the iteratee function.
 * @returns {T1[]} A new array containing the elements from the first array that are present
 *  in all subsequent arrays after applying the iteratee.
 *
 * @example
 * const array1 = [{ x: 1 }, { x: 2 }, { x: 3 }];
 * const array2 = [{ x: 2 }, { x: 3 }];
 * const array3 = [{ x: 3 }];
 * const array4 = [{ x: 3 }, { x: 4 }];
 * const result = intersectionBy(array1, array2, array3, array4, 'x');
 * // result will be [{ x: 3 }] since this element has the same `x` property in all arrays.
 *
 * const array1 = [{ x: 1 }, { x: 2 }, { x: 3 }];
 * const array2 = [{ x: 2 }, { x: 3 }];
 * const array3 = [{ x: 3 }];
 * const array4 = [{ x: 3 }, { x: 4 }];
 * const result = intersectionBy(array1, array2, array3, array4, value => value.x);
 * // result will be [{ x: 3 }] since this element has the same `x` property in all arrays.
 */
declare function intersectionBy<T1, T2, T3, T4>(array: ArrayLike<T1> | null | undefined, values1: ArrayLike<T2>, values2: ArrayLike<T3>, ...values: Array<ArrayLike<T4> | ((value: T1 | T2 | T3 | T4) => unknown) | string>): T1[];
/**
 * Returns the intersection of multiple arrays after applying the iteratee function to their elements.
 *
 * This function takes multiple arrays and an iteratee function (or property key) to
 * compare the elements after transforming them. It returns a new array containing the elements from
 * the first array that are present in all subsequent arrays after applying the iteratee to each element.
 *
 * @template T
 * @param {ArrayLike<T> | null | undefined} [array] - The first array to compare.
 * @param {...ArrayLike<T>} values - Additional arrays to compare.
 * @returns {T[]} A new array containing the elements from the first array that are present
 *  in all subsequent arrays after applying the iteratee.
 *
 * @example
 * const array1 = [1, 2, 3];
 * const array2 = [2, 3];
 * const array3 = [3];
 * const result = intersectionBy(array1, array2, array3);
 * // result will be [3] since these all elements have the same value 3.
 */
declare function intersectionBy<T>(array?: ArrayLike<T> | null | undefined, ...values: Array<ArrayLike<T>>): T[];

export { intersectionBy };
