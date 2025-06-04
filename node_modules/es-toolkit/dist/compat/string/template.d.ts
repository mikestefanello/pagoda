import { escape } from './escape.js';

declare const templateSettings: {
    escape: RegExp;
    evaluate: RegExp;
    interpolate: RegExp;
    variable: string;
    imports: {
        _: {
            escape: typeof escape;
            template: typeof template;
        };
    };
};
interface TemplateOptions {
    escape?: RegExp;
    evaluate?: RegExp;
    interpolate?: RegExp;
    variable?: string;
    imports?: Record<string, unknown>;
    sourceURL?: string;
}
/**
 * Compiles a template string into a function that can interpolate data properties.
 *
 * This function allows you to create a template with custom delimiters for escaping,
 * evaluating, and interpolating values. It can also handle custom variable names and
 * imported functions.
 *
 * @param {string} string - The template string.
 * @param {TemplateOptions} [options] - The options object.
 * @param {RegExp} [options.escape] - The regular expression for "escape" delimiter.
 * @param {RegExp} [options.evaluate] - The regular expression for "evaluate" delimiter.
 * @param {RegExp} [options.interpolate] - The regular expression for "interpolate" delimiter.
 * @param {string} [options.variable] - The data object variable name.
 * @param {Record<string, unknown>} [options.imports] - The object of imported functions.
 * @param {string} [options.sourceURL] - The source URL of the template.
 * @param {unknown} [guard] - The guard to detect if the function is called with `options`.
 * @returns {(data?: object) => string} Returns the compiled template function.
 *
 * @example
 * // Use the "escape" delimiter to escape data properties.
 * const compiled = template('<%- value %>');
 * compiled({ value: '<div>' }); // returns '&lt;div&gt;'
 *
 * @example
 * // Use the "interpolate" delimiter to interpolate data properties.
 * const compiled = template('<%= value %>');
 * compiled({ value: 'Hello, World!' }); // returns 'Hello, World!'
 *
 * @example
 * // Use the "evaluate" delimiter to evaluate JavaScript code.
 * const compiled = template('<% if (value) { %>Yes<% } else { %>No<% } %>');
 * compiled({ value: true }); // returns 'Yes'
 *
 * @example
 * // Use the "variable" option to specify the data object variable name.
 * const compiled = template('<%= data.value %>', { variable: 'data' });
 * compiled({ value: 'Hello, World!' }); // returns 'Hello, World!'
 *
 * @example
 * // Use the "imports" option to import functions.
 * const compiled = template('<%= _.toUpper(value) %>', { imports: { _: { toUpper } } });
 * compiled({ value: 'hello, world!' }); // returns 'HELLO, WORLD!'
 *
 * @example
 * // Use the custom "escape" delimiter.
 * const compiled = template('<@ value @>', { escape: /<@([\s\S]+?)@>/g });
 * compiled({ value: '<div>' }); // returns '&lt;div&gt;'
 *
 * @example
 * // Use the custom "evaluate" delimiter.
 * const compiled = template('<# if (value) { #>Yes<# } else { #>No<# } #>', { evaluate: /<#([\s\S]+?)#>/g });
 * compiled({ value: true }); // returns 'Yes'
 *
 * @example
 * // Use the custom "interpolate" delimiter.
 * const compiled = template('<$ value $>', { interpolate: /<\$([\s\S]+?)\$>/g });
 * compiled({ value: 'Hello, World!' }); // returns 'Hello, World!'
 *
 * @example
 * // Use the "sourceURL" option to specify the source URL of the template.
 * const compiled = template('hello <%= user %>!', { sourceURL: 'template.js' });
 */
declare function template(string: string, options?: TemplateOptions, guard?: unknown): ((data?: object) => string) & {
    source: string;
};

export { template, templateSettings };
