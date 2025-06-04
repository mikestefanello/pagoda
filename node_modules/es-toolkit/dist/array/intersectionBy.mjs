function intersectionBy(firstArr, secondArr, mapper) {
    const mappedSecondSet = new Set(secondArr.map(mapper));
    return firstArr.filter(item => mappedSecondSet.has(mapper(item)));
}

export { intersectionBy };
