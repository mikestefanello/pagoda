'use strict';

function capitalize(str) {
    return (str.charAt(0).toUpperCase() + str.slice(1).toLowerCase());
}

const CASE_SPLIT_PATTERN = /\p{Lu}?\p{Ll}+|[0-9]+|\p{Lu}+(?!\p{Ll})|\p{Emoji_Presentation}|\p{Extended_Pictographic}|\p{L}+/gu;
function words(str) {
    return Array.from(str.match(CASE_SPLIT_PATTERN) ?? []);
}

function camelCase(str) {
    const words$1 = words(str);
    if (words$1.length === 0) {
        return '';
    }
    const [first, ...rest] = words$1;
    return `${first.toLowerCase()}${rest.map(word => capitalize(word)).join('')}`;
}

function snakeCase(str) {
    const words$1 = words(str);
    return words$1.map(word => word.toLowerCase()).join('_');
}

exports.camelCase = camelCase;
exports.capitalize = capitalize;
exports.snakeCase = snakeCase;
exports.words = words;
