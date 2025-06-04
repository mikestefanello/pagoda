/**
 * Merges the properties of one or more source objects into the target object.
 *
 * This function performs a deep merge, recursively merging nested objects and arrays.
 * If a property in the source object is an array or object and the corresponding property in the target object is also an array or object, they will be merged.
 * If a property in the source object is `undefined`, it will not overwrite a defined property in the target object.
 *
 * You can provide a custom `merge` function to control how properties are merged. The `merge` function is called for each property that is being merged and receives the following arguments:
 *
 * - `targetValue`: The current value of the property in the target object.
 * - `sourceValue`: The value of the property in the source object.
 * - `key`: The key of the property being merged.
 * - `target`: The target object.
 * - `source`: The source object.
 * - `stack`: A `Map` used to keep track of objects that have already been processed to handle circular references.
 *
 * The `merge` function should return the value to be set in the target object. If it returns `undefined`, a default deep merge will be applied for arrays and objects.
 *
 * The function can handle multiple source objects and will merge them all into the target object.
 *
 * @param {T} target - The target object into which the source object properties will be merged. This object is modified in place.
 * @param {S} source - The first source object whose properties will be merged into the target object.
 * @param {(targetValue: any, sourceValue: any, key: string, target: T, source: S, stack: Map<any, any>) => any} merge - The function to customize merging properties.
 * @returns {T & S} The updated target object with properties from the source object(s) merged in.
 *
 * @template T - Type of the target object.
 * @template S - Type of the first source object.
 *
 * @example
 * const target = { a: 1, b: 2 };
 * const source = { b: 3, c: 4 };
 *
 * mergeWith(target, source, (targetValue, sourceValue) => {
 *   if (typeof targetValue === 'number' && typeof sourceValue === 'number') {
 *     return targetValue + sourceValue;
 *   }
 * });
 * // Returns { a: 1, b: 5, c: 4 }
 * @example
 * const target = { a: [1], b: [2] };
 * const source = { a: [3], b: [4] };
 *
 * const result = mergeWith(target, source, (objValue, srcValue) => {
 *   if (Array.isArray(objValue)) {
 *     return objValue.concat(srcValue);
 *   }
 * });
 *
 * expect(result).toEqual({ a: [1, 3], b: [2, 4] });
 */
declare function mergeWith<T, S>(target: T, source: S, merge: (targetValue: any, sourceValue: any, key: string, target: T, source: S, stack: Map<any, any>) => any): T & S;
/**
 * Merges the properties of one or more source objects into the target object.
 *
 * This function performs a deep merge, recursively merging nested objects and arrays.
 * If a property in the source object is an array or object and the corresponding property in the target object is also an array or object, they will be merged.
 * If a property in the source object is `undefined`, it will not overwrite a defined property in the target object.
 *
 * You can provide a custom `merge` function to control how properties are merged. The `merge` function is called for each property that is being merged and receives the following arguments:
 *
 * - `targetValue`: The current value of the property in the target object.
 * - `sourceValue`: The value of the property in the source object.
 * - `key`: The key of the property being merged.
 * - `target`: The target object.
 * - `source`: The source object.
 * - `stack`: A `Map` used to keep track of objects that have already been processed to handle circular references.
 *
 * The `merge` function should return the value to be set in the target object. If it returns `undefined`, a default deep merge will be applied for arrays and objects.
 *
 * The function can handle multiple source objects and will merge them all into the target object.
 *
 * @param {O} object - The target object into which the source object properties will be merged. This object is modified in place.
 * @param {S1} source1 - The first source object to be merged into the target object.
 * @param {S2} source2 - The second source object to be merged into the target object.
 * @param {(targetValue: any, sourceValue: any, key: string, target: any, source: any, stack: Map<any, any>) => any} merge - The function to customize merging properties.
 * @returns {O & S1 & S2} The updated target object with properties from the source objects merged in.
 *
 * @template O - Type of the target object.
 * @template S1 - Type of the first source object.
 * @template S2 - Type of the second source object.
 *
 * @example
 * const target = { a: 1, b: 2 };
 * const source = { b: 3, c: 4 };
 *
 * mergeWith(target, source, (targetValue, sourceValue) => {
 *   if (typeof targetValue === 'number' && typeof sourceValue === 'number') {
 *     return targetValue + sourceValue;
 *   }
 * });
 * // Returns { a: 1, b: 5, c: 4 }
 * @example
 * const target = { a: [1], b: [2] };
 * const source = { a: [3], b: [4] };
 *
 * const result = mergeWith(target, source, (objValue, srcValue) => {
 *   if (Array.isArray(objValue)) {
 *     return objValue.concat(srcValue);
 *   }
 * });
 *
 * expect(result).toEqual({ a: [1, 3], b: [2, 4] });
 */
declare function mergeWith<O, S1, S2>(object: O, source1: S1, source2: S2, merge: (targetValue: any, sourceValue: any, key: string, target: any, source: any, stack: Map<any, any>) => any): O & S1 & S2;
/**
 * Merges the properties of one or more source objects into the target object.
 *
 * This function performs a deep merge, recursively merging nested objects and arrays.
 * If a property in the source object is an array or object and the corresponding property in the target object is also an array or object, they will be merged.
 * If a property in the source object is `undefined`, it will not overwrite a defined property in the target object.
 *
 * You can provide a custom `merge` function to control how properties are merged. The `merge` function is called for each property that is being merged and receives the following arguments:
 *
 * - `targetValue`: The current value of the property in the target object.
 * - `sourceValue`: The value of the property in the source object.
 * - `key`: The key of the property being merged.
 * - `target`: The target object.
 * - `source`: The source object.
 * - `stack`: A `Map` used to keep track of objects that have already been processed to handle circular references.
 *
 * The `merge` function should return the value to be set in the target object. If it returns `undefined`, a default deep merge will be applied for arrays and objects.
 *
 * The function can handle multiple source objects and will merge them all into the target object.
 *
 * @param {O} object - The target object into which the source object properties will be merged. This object is modified in place.
 * @param {S1} source1 - The first source object whose properties will be merged into the target object.
 * @param {S2} source2 - The second source object whose properties will be merged into the target object.
 * @param {S3} source3 - The third source object whose properties will be merged into the target object.
 * @param {(targetValue: any, sourceValue: any, key: string, target: any, source: any, stack: Map<any, any>) => any} merge - The function to customize merging properties.
 * @returns {O & S1 & S2 & S3} The updated target object with properties from the source object(s) merged in.
 *
 * @template O - Type of the target object.
 * @template S1 - Type of the first source object.
 * @template S2 - Type of the second source object.
 * @template S3 - Type of the third source object.
 *
 * @example
 * const target = { a: 1, b: 2 };
 * const source = { b: 3, c: 4 };
 *
 * mergeWith(target, source, (targetValue, sourceValue) => {
 *   if (typeof targetValue === 'number' && typeof sourceValue === 'number') {
 *     return targetValue + sourceValue;
 *   }
 * });
 * // Returns { a: 1, b: 5, c: 4 }
 * @example
 * const target = { a: [1], b: [2] };
 * const source = { a: [3], b: [4] };
 *
 * const result = mergeWith(target, source, (objValue, srcValue) => {
 *   if (Array.isArray(objValue)) {
 *     return objValue.concat(srcValue);
 *   }
 * });
 *
 * expect(result).toEqual({ a: [1, 3], b: [2, 4] });
 */
declare function mergeWith<O, S1, S2, S3>(object: O, source1: S1, source2: S2, source3: S3, merge: (targetValue: any, sourceValue: any, key: string, target: any, source: any, stack: Map<any, any>) => any): O & S1 & S2 & S3;
/**
 * Merges the properties of one or more source objects into the target object.
 *
 * This function performs a deep merge, recursively merging nested objects and arrays.
 * If a property in the source object is an array or object and the corresponding property in the target object is also an array or object, they will be merged.
 * If a property in the source object is `undefined`, it will not overwrite a defined property in the target object.
 *
 * You can provide a custom `merge` function to control how properties are merged. The `merge` function is called for each property that is being merged and receives the following arguments:
 *
 * - `targetValue`: The current value of the property in the target object.
 * - `sourceValue`: The value of the property in the source object.
 * - `key`: The key of the property being merged.
 * - `target`: The target object.
 * - `source`: The source object.
 * - `stack`: A `Map` used to keep track of objects that have already been processed to handle circular references.
 *
 * The `merge` function should return the value to be set in the target object. If it returns `undefined`, a default deep merge will be applied for arrays and objects.
 *
 * The function can handle multiple source objects and will merge them all into the target object.
 *
 * @param {O} object - The target object into which the source object properties will be merged. This object is modified in place.
 * @param {S1} source1 - The first source object whose properties will be merged into the target object.
 * @param {S2} source2 - The second source object whose properties will be merged into the target object.
 * @param {S3} source3 - The third source object whose properties will be merged into the target object.
 * @param {S4} source4 - The fourth source object whose properties will be merged into the target object.
 * @param {(targetValue: any, sourceValue: any, key: string, target: any, source: any, stack: Map<any, any>) => any} merge - The function to customize merging properties.
 * @returns {O & S1 & S2 & S3 & S4} The updated target object with properties from the source object(s) merged in.
 *
 * @template O - Type of the target object.
 * @template S1 - Type of the first source object.
 * @template S2 - Type of the second source object.
 * @template S3 - Type of the third source object.
 * @template S4 - Type of the fourth source object.
 *
 * @example
 * const target = { a: 1, b: 2 };
 * const source = { b: 3, c: 4 };
 *
 * mergeWith(target, source, (targetValue, sourceValue) => {
 *   if (typeof targetValue === 'number' && typeof sourceValue === 'number') {
 *     return targetValue + sourceValue;
 *   }
 * });
 * // Returns { a: 1, b: 5, c: 4 }
 * @example
 * const target = { a: [1], b: [2] };
 * const source = { a: [3], b: [4] };
 *
 * const result = mergeWith(target, source, (objValue, srcValue) => {
 *   if (Array.isArray(objValue)) {
 *     return objValue.concat(srcValue);
 *   }
 * });
 *
 * expect(result).toEqual({ a: [1, 3], b: [2, 4] });
 */
declare function mergeWith<O, S1, S2, S3, S4>(object: O, source1: S1, source2: S2, source3: S3, source4: S4, merge: (targetValue: any, sourceValue: any, key: string, target: any, source: any, stack: Map<any, any>) => any): O & S1 & S2 & S3;
/**
 * Merges the properties of one or more source objects into the target object.
 *
 * This function performs a deep merge, recursively merging nested objects and arrays.
 * If a property in the source object is an array or object and the corresponding property in the target object is also an array or object, they will be merged.
 * If a property in the source object is `undefined`, it will not overwrite a defined property in the target object.
 *
 * You can provide a custom `merge` function to control how properties are merged. The `merge` function is called for each property that is being merged and receives the following arguments:
 *
 * - `targetValue`: The current value of the property in the target object.
 * - `sourceValue`: The value of the property in the source object.
 * - `key`: The key of the property being merged.
 * - `target`: The target object.
 * - `source`: The source object.
 * - `stack`: A `Map` used to keep track of objects that have already been processed to handle circular references.
 *
 * The `merge` function should return the value to be set in the target object. If it returns `undefined`, a default deep merge will be applied for arrays and objects.
 *
 * The function can handle multiple source objects and will merge them all into the target object.
 *
 * @param {any} object - The target object into which the source object properties will be merged. This object is modified in place.
 * @param {any[]} sources - The source objects whose properties will be merged into the target object.
 * @param {...any} otherArgs - Additional source objects to merge into the target object, including the custom `merge` function.
 * @returns {any} The updated target object with properties from the source object(s) merged in.
 *
 * @example
 * const target = { a: 1, b: 2 };
 * const source = { b: 3, c: 4 };
 *
 * mergeWith(target, source, (targetValue, sourceValue) => {
 *   if (typeof targetValue === 'number' && typeof sourceValue === 'number') {
 *     return targetValue + sourceValue;
 *   }
 * });
 * // Returns { a: 1, b: 5, c: 4 }
 * @example
 * const target = { a: [1], b: [2] };
 * const source = { a: [3], b: [4] };
 *
 * const result = mergeWith(target, source, (objValue, srcValue) => {
 *   if (Array.isArray(objValue)) {
 *     return objValue.concat(srcValue);
 *   }
 * });
 *
 * expect(result).toEqual({ a: [1, 3], b: [2, 4] });
 */
declare function mergeWith(object: any, ...otherArgs: any[]): any;

export { mergeWith };
