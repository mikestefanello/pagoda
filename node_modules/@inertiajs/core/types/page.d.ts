import { Component, Page, PageEvent, PageHandler, PageResolver, PreserveStateOption, RouterInitParams, VisitOptions } from './types';
declare class CurrentPage {
    protected page: Page;
    protected swapComponent: PageHandler;
    protected resolveComponent: PageResolver;
    protected componentId: {};
    protected listeners: {
        event: PageEvent;
        callback: VoidFunction;
    }[];
    protected isFirstPageLoad: boolean;
    protected cleared: boolean;
    init({ initialPage, swapComponent, resolveComponent }: RouterInitParams): this;
    set(page: Page, { replace, preserveScroll, preserveState, }?: Partial<Pick<VisitOptions, 'replace' | 'preserveScroll' | 'preserveState'>>): Promise<void>;
    setQuietly(page: Page, { preserveState, }?: {
        preserveState?: PreserveStateOption;
    }): Promise<unknown>;
    clear(): void;
    isCleared(): boolean;
    get(): Page;
    merge(data: Partial<Page>): void;
    setUrlHash(hash: string): void;
    remember(data: Page['rememberedState']): void;
    swap({ component, page, preserveState, }: {
        component: Component;
        page: Page;
        preserveState: PreserveStateOption;
    }): Promise<unknown>;
    resolve(component: string): Promise<Component>;
    isTheSame(page: Page): boolean;
    on(event: PageEvent, callback: VoidFunction): VoidFunction;
    fireEventsFor(event: PageEvent): void;
}
export declare const page: CurrentPage;
export {};
