/**
 * Returns a `identity` function when `value` is `null` or `undefined`.
 *
 * @param {null} [value] - The value to convert to an iteratee.
 * @returns {<T>(value: T) => T} - Returns a `identity` function.
 *
 * @example
 * const func = iteratee();
 * [{ a: 1 }, { a: 2 }, { a: 3 }].map(func) // => [{ a: 1 }, { a: 2 }, { a: 3 }]
 */
declare function iteratee(value?: null): <T>(value: T) => T;
/**
 * Returns a given `func` function when `value` is a `function`.
 *
 * @template {(...args: any[]) => unknown} F - The function type.
 * @param {F} func - The function to return.
 * @returns {F} - Returns the given function.
 *
 * @example
 * const func = iteratee((object) => object.a);
 * [{ a: 1 }, { a: 2 }, { a: 3 }].map(func) // => [1, 2, 3]
 */
declare function iteratee<F extends (...args: any[]) => unknown>(func: F): F;
/**
 * Creates a function that invokes `value` with the arguments of the created function.
 *
 * The created function returns the property value for a given element.
 *
 * @param {symbol | number | string | object | null} value - The value to convert to an iteratee.
 * @returns {(...args: any[]) => any} - Returns the new iteratee function.
 *
 * @example
 * const func = iteratee('a');
 * [{ a: 1 }, { a: 2 }, { a: 3 }].map(func) // => [1, 2, 3]
 *
 * const func = iteratee({ a: 1 });
 * [{ a: 1 }, { a: 2 }, { a: 3 }].find(func) // => { a: 1 }
 *
 * const func = iteratee(['a', 1]);
 * [{ a: 1 }, { a: 2 }, { a: 3 }].find(func) // => { a: 1 }
 */
declare function iteratee(value?: symbol | number | string | object | null): (...args: any[]) => any;

export { iteratee };
