'use strict';

function random(minimum, maximum) {
    if (maximum == null) {
        maximum = minimum;
        minimum = 0;
    }
    if (minimum >= maximum) {
        throw new Error('Invalid input: The maximum value must be greater than the minimum value.');
    }
    return Math.random() * (maximum - minimum) + minimum;
}

function randomInt(minimum, maximum) {
    return Math.floor(random(minimum, maximum));
}

exports.random = random;
exports.randomInt = randomInt;
