/**
 * Converts `value` to a finite number.
 *
 * @param {unknown} value - The value to convert.
 * @returns {number} Returns the number.
 *
 * @example
 * toNumber(3.2); // => 3.2
 * toNumber(Number.MIN_VALUE); // => 5e-324
 * toNumber(Infinity); // => 1.7976931348623157e+308
 * toNumber('3.2'); // => 3.2
 * toNumber(Symbol.iterator); // => 0
 * toNumber(NaN); // => 0
 */
declare function toFinite(value?: unknown): number;

export { toFinite };
