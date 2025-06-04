/**
 * Reduces an array to a single value using an iteratee function, starting from the right.
 *
 * The `reduceRight()` function goes through each element in an array from right to left and applies a special function (called a "reducer") to them, one by one.
 * This function takes the result of the previous step and the current element to perform a calculation.
 * After going through all the elements, the function gives you one final result.
 *
 * When the `reduceRight()` function starts, there's no previous result to use.
 * If you provide an initial value, it starts with that.
 * If not, it uses the last element of the array and begins with the second to last element for the calculation.
 *
 * @param {readonly T[]} collection - The collection to iterate over.
 * @param {(accumulator: U, value: T, index: number, collection: readonly T[]) => U} iteratee - The function invoked per iteration.
 * @param {U} initialValue - The initial value.
 * @returns {U} - Returns the accumulated value.
 *
 * @example
 * const arrayLike = [1, 2, 3];
 * reduceRight(arrayLike, (acc, value) => acc && value % 2 === 0, true); // => false
 */
declare function reduceRight<T, U>(collection: readonly T[], iteratee: (accumulator: U, value: T, index: number, collection: readonly T[]) => U, initialValue: U): U;
/**
 * Reduces an array to a single value using an iteratee function, starting from the right.
 *
 * The `reduceRight()` function goes through each element in an array from right to left and applies a special function (called a "reducer") to them, one by one.
 * This function takes the result of the previous step and the current element to perform a calculation.
 * After going through all the elements, the function gives you one final result.
 *
 * When the `reduceRight()` function starts, there's no previous result to use.
 * If you provide an initial value, it starts with that.
 * If not, it uses the last element of the array and begins with the second to last element for the calculation.
 *
 * @param {readonly T[]} collection - The collection to iterate over.
 * @param {(accumulator: T, value: T, index: number, collection: readonly T[]) => T} iteratee - The function invoked per iteration.
 * @returns {T} - Returns the accumulated value.
 *
 * @example
 * const arrayLike = [1, 2, 3];
 * reduceRight(arrayLike, (acc, value) => acc + value); // => 6
 */
declare function reduceRight<T>(collection: readonly T[], iteratee: (accumulator: T, value: T, index: number, collection: readonly T[]) => T): T;
/**
 * Reduces an array to a single value using an iteratee function, starting from the right.
 *
 * The `reduceRight()` function goes through each element in an array from right to left and applies a special function (called a "reducer") to them, one by one.
 * This function takes the result of the previous step and the current element to perform a calculation.
 * After going through all the elements, the function gives you one final result.
 *
 * When the `reduceRight()` function starts, there's no previous result to use.
 * If you provide an initial value, it starts with that.
 * If not, it uses the last element of the array and begins with the second to last element for the calculation.
 *
 * @param {ArrayLike<T>} collection - The collection to iterate over.
 * @param {(accumulator: U, value: T, index: number, collection: ArrayLike<T>) => U} iteratee - The function invoked per iteration.
 * @param {U} initialValue - The initial value.
 * @returns {U} - Returns the accumulated value.
 *
 * @example
 * const arrayLike = {0: 1, 1: 2, 2: 3, length: 3};
 * reduceRight(arrayLike, (acc, value) => acc + value % 2 === 0, true); // => false
 */
declare function reduceRight<T, U>(collection: ArrayLike<T>, iteratee: (accumulator: U, value: T, index: number, collection: ArrayLike<T>) => U, initialValue: U): U;
/**
 * Reduces an array to a single value using an iteratee function, starting from the right.
 *
 * The `reduceRight()` function goes through each element in an array from right to left and applies a special function (called a "reducer") to them, one by one.
 * This function takes the result of the previous step and the current element to perform a calculation.
 * After going through all the elements, the function gives you one final result.
 *
 * When the `reduceRight()` function starts, there's no previous result to use.
 * If you provide an initial value, it starts with that.
 * If not, it uses the last element of the array and begins with the second to last element for the calculation.
 *
 * @param {ArrayLike<T>} collection - The collection to iterate over.
 * @param {(accumulator: U, value: T, index: number, collection: ArrayLike<T>) => U} iteratee - The function invoked per iteration.
 * @returns {T} - Returns the accumulated value.
 *
 * @example
 * const arrayLike = {0: 1, 1: 2, 2: 3, length: 3};
 * reduceRight(arrayLike, (acc, value) => acc + value); // => 6
 */
declare function reduceRight<T>(collection: ArrayLike<T>, iteratee: (accumulator: T, value: T, index: number, collection: ArrayLike<T>) => T): T;
/**
 * Reduces an object to a single value using an iteratee function, starting from the right.
 *
 * @param {T} collection - The object to iterate over.
 * @param {(accumulator: U, value: T[keyof T], key: string, collection: T) => U} iteratee - The function invoked per iteration.
 * @param {U} initialValue - The initial value.
 * @returns {U} - Returns the accumulated value.
 *
 * @example
 * const obj = { a: 1, b: 2, c: 3 };
 * reduceRight(obj, (acc, value) => acc + value % 2 === 0, true); // => false
 */
declare function reduceRight<T extends object, U>(collection: T, iteratee: (accumulator: U, value: T[keyof T], key: keyof T, collection: T) => U, initialValue: U): U;
/**
 * Reduces an object to a single value using an iteratee function, starting from the right.
 *
 * @param {T} collection - The object to iterate over.
 * @param {(accumulator: T[keyof T], value: T[keyof T], key: keyof T, collection: T) => U} iteratee - The function invoked per iteration.
 * @returns {T[keyof T]} - Returns the accumulated value.
 *
 * @example
 * const obj = { a: 1, b: 2, c: 3 };
 * reduceRight(obj, (acc, value) => acc + value); // => 6
 */
declare function reduceRight<T extends object>(collection: T, iteratee: (accumulator: T[keyof T], value: T[keyof T], key: keyof T, collection: T) => T[keyof T]): T[keyof T];

export { reduceRight };
