export declare class InitialVisit {
    static handle(): void;
    protected static clearRememberedStateOnReload(): void;
    protected static handleBackForward(): boolean;
    /**
     * @link https://inertiajs.com/redirects#external-redirects
     */
    protected static handleLocation(): boolean;
    protected static handleDefault(): void;
}
