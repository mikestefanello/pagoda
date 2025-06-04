/**
 * Creates an array of elements split into two groups, the first of which contains elements
 * `predicate` returns truthy for, the second of which contains elements
 * `predicate` returns falsy for. The predicate is invoked with one argument: (value).
 *
 * @template T
 * @template U
 * @param {ArrayLike<T> | T | null | undefined} source - The array or object to iterate over.
 * @param {(value: T) => value is U} predicate - The function invoked per iteration.
 * @returns {[U[], Array<Exclude<T, U>>]} - Returns the array of grouped elements.
 *
 * @example
 * partition([1, 2, 3, 4], n => n % 2 === 0);
 * // => [[2, 4], [1, 3]]
 *
 * partition([1, 2, 3, 4], 'a');
 * // => [[2, 4], [1, 3]]
 *
 * partition([1, 2, 3, 4], ['a', 2]);
 */
declare function partition<T, U extends T>(source: ArrayLike<T> | null | undefined, predicate: (value: T) => value is U): [U[], Array<Exclude<T, U>>];
/**
 * Creates an array of elements split into two groups, the first of which contains elements
 * `predicate` returns truthy for, the second of which contains elements
 * `predicate` returns falsy for. The predicate is invoked with one argument: (value).
 *
 * @template T
 * @param {ArrayLike<T> | T | null | undefined} source - The array or object to iterate over.
 * @param {((item: T, index: number, arr: any) => unknown) | Partial<T> | [keyof T, unknown] | PropertyKey} [predicate=identity] - The function invoked per iteration.
 * @returns {[T[], T[]]} - Returns the array of grouped elements.
 *
 * @example
 * partition([1, 2, 3, 4], n => n % 2 === 0);
 * // => [[2, 4], [1, 3]]
 *
 * partition([1, 2, 3, 4], 'a');
 * // => [[2, 4], [1, 3]]
 *
 * partition([1, 2, 3, 4], ['a', 2]);
 * // => [[2], [1, 3, 4]]
 */
declare function partition<T>(source: ArrayLike<T> | null | undefined, predicate?: ((value: T) => unknown) | PropertyKey | [PropertyKey, any] | Partial<T>): [T[], T[]];
/**
 * Creates an array of elements split into two groups, the first of which contains elements
 * `predicate` returns truthy for, the second of which contains elements
 * `predicate` returns falsy for. The predicate is invoked with one argument: (value).
 *
 * @template T
 * @param {T | null | undefined} object - The object to iterate over.
 * @param {((item: T[keyof T], key: keyof T) => unknown) | Partial<T[keyof T]> | [keyof T, unknown] | PropertyKey} [predicate=identity] - The function invoked per iteration.
 * @returns {[T[keyof T][], T[keyof T][]]} - Returns the array of grouped elements.
 *
 * @example
 * partition({ a: 1, b: 2, c: 3 }, n => n % 2 === 0);
 * // => [[2], [1, 3]]
 *
 * partition({ a: 1, b: 2, c: 3 }, 'a');
 * // => [[1], [2, 3]]
 *
 * partition({ a: 1, b: 2, c: 3 }, ['a', 1]);
 * // => [[1], [2, 3]]
 */
declare function partition<T extends object>(object: T | null | undefined, predicate?: ((value: T[keyof T]) => unknown) | PropertyKey | [PropertyKey, any] | Partial<T[keyof T]>): [Array<T[keyof T]>, Array<T[keyof T]>];

export { partition };
