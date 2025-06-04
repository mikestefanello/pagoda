/**
 * Removes trailing whitespace or specified characters from a string.
 *
 * @param {string} str - The string from which trailing characters will be trimmed.
 * @param {string | string[]} chars - The character(s) to remove from the end of the string. Defaults to `" "`.
 * @returns {string} - The resulting string after the specified trailing character has been removed.
 *
 * @example
 * const trimmedStr1 = trimEnd('hello---', '-') // returns 'hello'
 * const trimmedStr2 = trimEnd('123000', '0') // returns '123'
 * const trimmedStr3 = trimEnd('abcabcabc', 'c') // returns 'abcabcab'
 * const trimmedStr4 = trimEnd('trimmedxxx', 'x') // returns 'trimmed'
 */
declare function trimEnd(str: string, chars?: string | string[], guard?: unknown): string;

export { trimEnd };
