export default class Queue<T> {
    protected items: (() => T)[];
    protected processingPromise: Promise<void> | null;
    add(item: () => T): Promise<void>;
    process(): Promise<void>;
    protected processNext(): Promise<void>;
}
