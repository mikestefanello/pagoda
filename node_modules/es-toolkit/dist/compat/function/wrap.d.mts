/**
 * Creates a new function that wraps the given function `func`.
 * In this process, you can apply additional logic defined in the `wrapper` function before and after the execution of the original function.
 *
 * @param {F} func - A function to wrap.
 * @param {(value: F, ...args: Parameters<F>) => ReturnType<F>} wrapper - A wrapper function that receives the original function and its arguments.
 *
 * @example
 * const greet = (name: string) => `Hi, ${name}`;
 * const wrapped = wrap(greet, (value, name) => `[LOG] ${value(name)}`);
 * wrapped('Bob'); // => "[LOG] Hi, Bob"
 */
declare function wrap<F extends (...args: unknown[]) => unknown>(func: F, wrapper: (value: F, ...args: Parameters<F>) => ReturnType<F>): F;
/**
 * Creates a new function that passes the original value `value` as the first argument to the `wrapper`.
 * This allows you to decorate or extend the behavior of the original value in a flexible way.
 *
 * @param {T} value - A non-function value to wrap.
 * @param {(value: T, ...args: A) => R} wrapper - A wrapper function that receives the value and arguments.
 *
 * @example
 * const wrapped = wrap('value', v => `<p>${v}</p>`);
 * wrapped(); // => "<p>value</p>"
 */
declare function wrap<T, A extends unknown[], R>(value: T, wrapper: (value: T, ...args: A) => R): (...args: A) => R;

export { wrap };
