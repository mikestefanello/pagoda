/**
 * Creates an array of string keyed-value pairs from an object.
 *
 * @param {Record<string | number, T>} object The object to query.
 * @returns {Array<[key: string, value: T]>} Returns the array of property pairs.
 * @example
 * const object = { a: 1, b: 2 };
 * toPairs(object); // [['a', 1], ['b', 2]]
 */
declare function toPairs<T>(object: Record<string | number, T>): Array<[key: string, value: T]>;
/**
 * Creates an array of value pairs from a set.
 *
 * @param {Set<T>} set The set to query.
 * @returns {Array<[key: T, value: T]>} Returns the array of value pairs.
 * @example
 * const set = new Set([1, 2]);
 * toPairs(set); // [[1, 1], [2, 2]]
 */
declare function toPairs<T>(set: Set<T>): Array<[key: T, value: T]>;
/**
 * Creates an array of key-value pairs from a map.
 *
 * @param {Map<K, V>} map The map to query.
 * @returns {Array<[key: K, value: V]>} Returns the array of key-value pairs.
 * @example
 * const map = new Map();
 * map.set('a', 1);
 * map.set('b', 2);
 * toPairs(map); // [['a', 1], ['b', 2]]
 */
declare function toPairs<K, V>(map: Map<K, V>): Array<[key: K, value: V]>;

export { toPairs };
