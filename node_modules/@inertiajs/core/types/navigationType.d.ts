declare class NavigationType {
    protected type: NavigationTimingType;
    constructor();
    protected resolveType(): NavigationTimingType;
    get(): NavigationTimingType;
    isBackForward(): boolean;
    isReload(): boolean;
}
export declare const navigationType: NavigationType;
export {};
