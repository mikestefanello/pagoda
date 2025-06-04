/**
 * Creates a function that accepts arguments of `func` and either invokes `func` returning its result, if at least `arity` number of arguments have been provided, or returns a function that accepts the remaining `func` arguments, and so on.
 * The arity of `func` may be specified if `func.length` is not sufficient.
 *
 * The `curry.placeholder` value, which defaults to a `symbol`, may be used as a placeholder for partially applied arguments.
 *
 * Note: This method doesn't set the `length` property of curried functions.
 *
 * @param {(...args: any[]) => any} func - The function to curry.
 * @param {number=func.length} arity - The arity of func.
 * @param {unknown} guard - Enables use as an iteratee for methods like `Array#map`.
 * @returns {((...args: any[]) => any) & { placeholder: typeof curry.placeholder }} - Returns the new curried function.
 *
 * @example
 * const abc = function(a, b, c) {
 *   return Array.from(arguments);
 * };
 *
 * let curried = curry(abc);
 *
 * curried(1)(2)(3);
 * // => [1, 2, 3]
 *
 * curried(1, 2)(3);
 * // => [1, 2, 3]
 *
 * curried(1, 2, 3);
 * // => [1, 2, 3]
 *
 * // Curried with placeholders.
 * curried(1)(curry.placeholder, 3)(2);
 * // => [1, 2, 3]
 *
 * // Curried with arity.
 * curried = curry(abc, 2);
 *
 * curried(1)(2);
 * // => [1, 2]
 */
declare function curry(func: (...args: any[]) => any, arity?: number, guard?: unknown): ((...args: any[]) => any) & {
    placeholder: typeof curry.placeholder;
};
declare namespace curry {
    var placeholder: typeof curryPlaceholder;
}
declare const curryPlaceholder: unique symbol;

export { curry };
