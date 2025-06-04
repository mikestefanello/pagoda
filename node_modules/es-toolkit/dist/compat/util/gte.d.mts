/**
 * Checks if value is greater than or equal to other.
 *
 * @param {unknown} value The value to compare.
 * @param {unknown} other The other value to compare.
 * @returns {boolean} Returns `true` if value is greater than or equal to other, else `false`.
 *
 * @example
 * gte(3, 1); // => true
 * gte(3, 3); // => true
 * gte(1, 3); // => false
 */
declare function gte(value: unknown, other: unknown): boolean;

export { gte };
