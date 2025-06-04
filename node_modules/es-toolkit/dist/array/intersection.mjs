function intersection(firstArr, secondArr) {
    const secondSet = new Set(secondArr);
    return firstArr.filter(item => {
        return secondSet.has(item);
    });
}

export { intersection };
