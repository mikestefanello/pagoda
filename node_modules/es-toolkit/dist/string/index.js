'use strict';

Object.defineProperty(exports, Symbol.toStringTag, { value: 'Module' });

const snakeCase = require('../_chunk/snakeCase-6cG1f4.js');
const upperFirst = require('../_chunk/upperFirst-Cx78bs.js');

function constantCase(str) {
    const words = snakeCase.words(str);
    return words.map(word => word.toUpperCase()).join('_');
}

function pascalCase(str) {
    const words = snakeCase.words(str);
    return words.map(word => snakeCase.capitalize(word)).join('');
}

function reverseString(value) {
    return [...value].reverse().join('');
}

function startCase(str) {
    const words = snakeCase.words(str.trim());
    let result = '';
    for (let i = 0; i < words.length; i++) {
        const word = words[i];
        if (result) {
            result += ' ';
        }
        result += word[0].toUpperCase() + word.slice(1).toLowerCase();
    }
    return result;
}

exports.camelCase = snakeCase.camelCase;
exports.capitalize = snakeCase.capitalize;
exports.snakeCase = snakeCase.snakeCase;
exports.words = snakeCase.words;
exports.deburr = upperFirst.deburr;
exports.escape = upperFirst.escape;
exports.escapeRegExp = upperFirst.escapeRegExp;
exports.kebabCase = upperFirst.kebabCase;
exports.lowerCase = upperFirst.lowerCase;
exports.lowerFirst = upperFirst.lowerFirst;
exports.pad = upperFirst.pad;
exports.trim = upperFirst.trim;
exports.trimEnd = upperFirst.trimEnd;
exports.trimStart = upperFirst.trimStart;
exports.unescape = upperFirst.unescape;
exports.upperCase = upperFirst.upperCase;
exports.upperFirst = upperFirst.upperFirst;
exports.constantCase = constantCase;
exports.pascalCase = pascalCase;
exports.reverseString = reverseString;
exports.startCase = startCase;
