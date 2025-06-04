/**
 * Updates the value at the specified path of the given object using an updater function and a customizer.
 * If any part of the path does not exist, it will be created.
 *
 * @template T - The type of the object.
 * @param {T} obj - The object to modify.
 * @param {PropertyKey | PropertyKey[]} path - The path of the property to update.
 * @param {(value: unknown) => unknown} updater - The function to produce the updated value.
 * @param {(value: unknown) => unknown} customizer - The function to customize the update process.
 * @returns {T} - The modified object.
 */
declare function updateWith<T extends object | null | undefined>(obj: T, path: PropertyKey | readonly PropertyKey[], updater: (value: unknown) => unknown, customizer: (value: unknown) => unknown): T;

export { updateWith };
