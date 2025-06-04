type Criterion<T> = ((item: T) => unknown) | PropertyKey | PropertyKey[] | null | undefined;
/**
 * Sorts an array of objects based on multiple properties and their corresponding order directions.
 *
 * This function takes an array of objects, an array of criteria to sort by, and an array of order directions.
 * It returns the sorted array, ordering by each key according to its corresponding direction ('asc' for ascending or 'desc' for descending).
 * If values for a key are equal, it moves to the next key to determine the order.
 *
 * @template T - The type of elements in the array.
 * @param {ArrayLike<T> | object | null | undefined} collection - The array of objects to be sorted.
 * @param {Criterion<T> | Array<Criterion<T>>} criteria - An array of criteria (property names or property paths or custom key functions) to sort by.
 * @param {unknown | unknown[]} orders - An array of order directions ('asc' for ascending or 'desc' for descending).
 * @param {unknown} [guard] Enables use as an iteratee for methods like `_.reduce`.
 * @returns {T[]} - The sorted array.
 *
 * @example
 * // Sort an array of objects by 'user' in ascending order and 'age' in descending order.
 * const users = [
 *   { user: 'fred', age: 48 },
 *   { user: 'barney', age: 34 },
 *   { user: 'fred', age: 40 },
 *   { user: 'barney', age: 36 },
 * ];
 * const result = orderBy(users, ['user', (item) => item.age], ['asc', 'desc']);
 * // result will be:
 * // [
 * //   { user: 'barney', age: 36 },
 * //   { user: 'barney', age: 34 },
 * //   { user: 'fred', age: 48 },
 * //   { user: 'fred', age: 40 },
 * // ]
 */
declare function orderBy<T = any>(collection: ArrayLike<T> | object | null | undefined, criteria?: Criterion<T> | Array<Criterion<T>>, orders?: unknown | unknown[], guard?: unknown): T[];

export { type Criterion, orderBy };
