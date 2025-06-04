export declare const encryptHistory: (data: any) => Promise<ArrayBuffer>;
export declare const historySessionStorageKeys: {
    key: string;
    iv: string;
};
export declare const decryptHistory: (data: any) => Promise<any>;
