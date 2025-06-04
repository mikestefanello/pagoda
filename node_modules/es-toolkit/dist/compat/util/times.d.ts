/**
 * Invokes the getValue function n times, returning an array of the results.
 *
 * @template R The return type of the getValue function.
 * @param {number} n - The number of times to invoke getValue.
 * @param {(index: number) => R} getValue - The function to invoke for each index.
 * @returns {R[]} An array containing the results of invoking getValue n times.
 * @example
 * times(3, (i) => i * 2); // => [0, 2, 4]
 * times(2, () => 'es-toolkit'); // => ['es-toolkit', 'es-toolkit']
 */
declare function times<R = number>(n?: number, getValue?: (index: number) => R): R[];

export { times };
