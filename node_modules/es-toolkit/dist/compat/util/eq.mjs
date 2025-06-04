function eq(value, other) {
    return value === other || (Number.isNaN(value) && Number.isNaN(other));
}

export { eq };
