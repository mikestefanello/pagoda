export declare class SessionStorage {
    static locationVisitKey: string;
    static set(key: string, value: any): void;
    static get(key: string): any;
    static merge(key: string, value: any): void;
    static remove(key: string): void;
    static removeNested(key: string, nestedKey: string): void;
    static exists(key: string): boolean;
    static clear(): void;
}
