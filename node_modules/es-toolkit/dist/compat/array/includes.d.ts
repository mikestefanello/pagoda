/**
 * Checks if an item is included in an array.
 *
 * @param {T[]} arr - The array to search in.
 * @param {T} item - The item to search for.
 * @param {number} [fromIndex=0] - The index to start searching from. If negative, it is treated as an offset from the end of the array.
 * @returns {boolean} `true` if the item is found in the array, `false` otherwise.
 *
 * @example
 * includes([1, 2, 3], 2); // true
 * includes([1, 2, 3], 4); // false
 * includes([1, 2, 3], 3, -1); // true
 */
declare function includes<T>(arr: readonly T[], item: T, fromIndex?: number): boolean;
/**
 * Checks if a value is included in an object.
 *
 * @param {T} obj - The object to search in.
 * @param {T[keyof T]} value - The value to search for.
 * @param {number} [fromIndex=0] - The index to start searching from. If negative, it is treated as an offset from the end of the array.
 * @returns {boolean} `true` if the value is found in the object, `false` otherwise.
 *
 * @example
 * includes({ a: 1, b: 'a', c: NaN }, 1); // true
 * includes({ a: 1, b: 'a', c: NaN }, 'a'); // true
 * includes({ a: 1, b: 'a', c: NaN }, NaN); // true
 * includes({ [Symbol('sym1')]: 1 }, 1); // false
 */
declare function includes<T extends Record<string, any>>(obj: T, value: T[keyof T], fromIndex?: number): boolean;
/**
 * Checks if a substring is included in a string.
 *
 * @param {string} str - The string to search in.
 * @param {string} substr - The substring to search for.
 * @param {number} [fromIndex=0] - The index to start searching from. If negative, it is treated as an offset from the end of the string.
 * @returns {boolean} `true` if the substring is found in the string, `false` otherwise.
 *
 * @example
 * includes('hello world', 'world'); // true
 * includes('hello world', 'test'); // false
 * includes('hello world', 'o', 5); // true
 */
declare function includes(str: string, substr: string, fromIndex?: number): boolean;
/**
 * Checks if a specified value exists within a given source, which can be an array, an object, or a string.
 *
 * The comparison uses SameValueZero to check for inclusion.
 *
 * @param {T[] | Record<string, any> | string} source - The source to search in. It can be an array, an object, or a string.
 * @param {T} [target] - The value to search for in the source.
 * @param {number} [fromIndex=0] - The index to start searching from. If negative, it is treated as an offset from the end of the source.
 * @returns {boolean} `true` if the value is found in the source, `false` otherwise.
 *
 * @example
 * includes([1, 2, 3], 2); // true
 * includes({ a: 1, b: 'a', c: NaN }, 'a'); // true
 * includes('hello world', 'world'); // true
 * includes('hello world', 'test'); // false
 */
declare function includes<T>(source: readonly T[] | Record<string, any> | string, target?: T, fromIndex?: number): boolean;

export { includes };
