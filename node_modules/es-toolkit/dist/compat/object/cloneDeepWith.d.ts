/**
 * Creates a deep clone of the given object using a customizer function.
 *
 * @template T - The type of the object.
 * @param {T} obj - The object to clone.
 * @param {Function} [cloneValue] - A function to customize the cloning process.
 * @returns {T} - A deep clone of the given object.
 *
 * @example
 * // Clone a primitive value
 * const num = 29;
 * const clonedNum = cloneDeepWith(num);
 * console.log(clonedNum); // 29
 * console.log(clonedNum === num); // true
 *
 * @example
 * // Clone an object with a customizer
 * const obj = { a: 1, b: 2 };
 * const clonedObj = cloneDeepWith(obj, (value) => {
 *   if (typeof value === 'number') {
 *     return value * 2; // Double the number
 *   }
 * });
 * console.log(clonedObj); // { a: 2, b: 4 }
 * console.log(clonedObj === obj); // false
 *
 * @example
 * // Clone an array with a customizer
 * const arr = [1, 2, 3];
 * const clonedArr = cloneDeepWith(arr, (value) => {
 *   return value + 1; // Increment each value
 * });
 * console.log(clonedArr); // [2, 3, 4]
 * console.log(clonedArr === arr); // false
 */
declare function cloneDeepWith<T>(obj: T, cloneValue?: (value: any, key: PropertyKey | undefined, object: T | undefined, stack: Map<any, any>) => any): T;

export { cloneDeepWith };
