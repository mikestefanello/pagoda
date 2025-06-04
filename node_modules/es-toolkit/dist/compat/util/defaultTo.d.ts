/**
 * Returns the default value for `null`, `undefined`, and `NaN`.
 *
 * @param {T | null | undefined} value - The value to check.
 * @param {T} [defaultValue] - The default value to return if the first value is null, undefined, or NaN.
 * @returns {T} Returns either the first value or the default value.
 *
 * @example
 * defaultTo(null, 'default') // returns 'default'
 * defaultTo(undefined, 42) // returns 42
 * defaultTo(NaN, 0) // returns 0
 * defaultTo('actual', 'default') // returns 'actual'
 * defaultTo(123, 0) // returns 123
 */
declare function defaultTo<T>(value: T | null | undefined, defaultValue?: T): T;

export { defaultTo };
