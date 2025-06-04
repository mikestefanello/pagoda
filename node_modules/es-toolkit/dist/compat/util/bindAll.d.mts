/**
 * Binds methods of an object to the object itself, overwriting the existing method.
 * Method names may be specified as individual arguments or as arrays of method names.
 *
 * @param {Object} object - The object to bind methods to.
 * @param {...(string|string[]|number|IArguments)} [methodNames] - The method names to bind, specified individually or in arrays.
 * @returns {Object} - Returns the object.
 *
 * @example
 * const view = {
 *   'label': 'docs',
 *   'click': function() {
 *     console.log('clicked ' + this.label);
 *   }
 * };
 *
 * bindAll(view, ['click']);
 * jQuery(element).on('click', view.click);
 * // => Logs 'clicked docs' when clicked.
 *
 * @example
 * // Using individual method names
 * bindAll(view, 'click');
 * // => Same as above
 */
declare function bindAll<T>(object: T, ...methodNames: Array<string | string[]>): T;

export { bindAll };
