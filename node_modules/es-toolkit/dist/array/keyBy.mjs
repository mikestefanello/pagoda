function keyBy(arr, getKeyFromItem) {
    const result = {};
    for (let i = 0; i < arr.length; i++) {
        const item = arr[i];
        const key = getKeyFromItem(item);
        result[key] = item;
    }
    return result;
}

export { keyBy };
