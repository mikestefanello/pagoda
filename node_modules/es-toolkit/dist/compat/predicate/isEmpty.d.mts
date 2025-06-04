/**
 * Checks if a given value is empty.
 *
 * @returns {true} Always returns true when no arguments are provided.
 *
 * @example
 * isEmpty(); // true
 */
declare function isEmpty(): true;
/**
 * Checks if a given string is empty.
 *
 * @param {string} value - The string to check.
 * @returns {boolean} `true` if the string is empty, `false` otherwise.
 *
 * @example
 * isEmpty(""); // true
 * isEmpty("hello"); // false
 */
declare function isEmpty(value: string): value is '';
/**
 * Checks if a given Map is empty.
 *
 * @param {Map<any, any>} value - The Map to check.
 * @returns {boolean} `true` if the Map is empty, `false` otherwise.
 *
 * @example
 * isEmpty(new Map()); // true
 * isEmpty(new Map([["key", "value"]])); // false
 */
declare function isEmpty(value: Map<any, any>): boolean;
/**
 * Checks if a given Set is empty.
 *
 * @param {Set<any>} value - The Set to check.
 * @returns {boolean} `true` if the Set is empty, `false` otherwise.
 *
 * @example
 * isEmpty(new Set()); // true
 * isEmpty(new Set([1, 2, 3])); // false
 */
declare function isEmpty(value: Set<any>): boolean;
/**
 * Checks if a given array is empty.
 *
 * @param {any[]} value - The array to check.
 * @returns {boolean} `true` if the array is empty, `false` otherwise.
 *
 * @example
 * isEmpty([]); // true
 * isEmpty([1, 2, 3]); // false
 */
declare function isEmpty(value: any[]): value is [];
/**
 * Checks if a given object is empty.
 *
 * @param {T | null | undefined} value - The object to check.
 * @returns {boolean} `true` if the object is empty, `false` otherwise.
 *
 * @example
 * isEmpty({}); // true
 * isEmpty({ a: 1 }); // false
 */
declare function isEmpty<T extends Record<any, any>>(value: T | null | undefined): value is Record<keyof T, never> | null | undefined;
/**
 * Checks if a given value is empty.
 *
 * @param {unknown} value - The value to check.
 * @returns {boolean} `true` if the value is empty, `false` otherwise.
 *
 * @example
 * isEmpty(null); // true
 * isEmpty(undefined); // true
 * isEmpty(42); // true
 */
declare function isEmpty(value: unknown): boolean;

export { isEmpty };
