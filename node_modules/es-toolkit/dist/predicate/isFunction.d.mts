/**
 * Checks if `value` is a function.
 *
 * @param {unknown} value The value to check.
 * @returns {boolean} Returns `true` if `value` is a function, else `false`.
 *
 * @example
 * isFunction(Array.prototype.slice); // true
 * isFunction(async function () {}); // true
 * isFunction(function* () {}); // true
 * isFunction(Proxy); // true
 * isFunction(Int8Array); // true
 */
declare function isFunction(value: unknown): value is (...args: unknown[]) => unknown;

export { isFunction };
