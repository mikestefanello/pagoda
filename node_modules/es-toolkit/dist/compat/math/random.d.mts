/**
 * Generate a random number within 0 and 1.
 *
 * @returns {number} A random number between 0 (inclusive) and 1 (exclusive). The number can be an integer or a decimal.
 * @throws {Error} Throws an error if `maximum` is not greater than `0`.
 *
 * @example
 * const result = random(); // Returns a random number between 0 and 1.
 */
declare function random(floating?: boolean): number;
/**
 * Generate a random number within 0 and 1.
 *
 * @returns {number} A random number between 0 (inclusive) and 1 (exclusive). The number can be an integer or a decimal.
 * @throws {Error} Throws an error if `maximum` is not greater than `0`.
 *
 * @example
 * const result = random(); // Returns a random number between 0 and 1.
 */
declare function random(min: number, index: string | number, guard: object): number;
/**
 * Generate a random number within the given range.
 *
 * If only one argument is provided, a number between `0` and the given number is returned.
 *
 * @param {number} maximum - The upper bound (exclusive).
 * @returns {number} A random number between 0 (inclusive) and maximum (exclusive). The number can be an integer or a decimal.
 * @throws {Error} Throws an error if `maximum` is not greater than `0`.
 *
 * @example
 * const result1 = random(5); // Returns a random number between 0 and 5.
 * const result2 = random(0); // Returns a random number between 0 and 0 (which is 0).
 */
declare function random(maximum: number, floating?: boolean): number;
/**
 * Generate a random number within the given range.
 *
 * @param {number} minimum - The lower bound (inclusive).
 * @param {number} maximum - The upper bound (exclusive).
 * @returns {number} A random number between minimum (inclusive) and maximum (exclusive). The number can be an integer or a decimal.
 * @throws {Error} Throws an error if `maximum` is not greater than `minimum`.
 *
 * @example
 * const result1 = random(0, 5); // Returns a random number between 0 and 5.
 * const result2 = random(5, 0); // If the minimum is greater than the maximum, an error is thrown.
 * const result3 = random(5, 5); // If the minimum is equal to the maximum, an error is thrown.
 */
declare function random(minimum: number, maximum: number, floating?: boolean): number;

export { random };
