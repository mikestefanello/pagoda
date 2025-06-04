import { RequestStream } from './requestStream';
import { ActiveVisit, ClientSideVisitOptions, GlobalEvent, GlobalEventNames, GlobalEventResult, InFlightPrefetch, Page, PendingVisit, PendingVisitOptions, PollOptions, PrefetchedResponse, PrefetchOptions, ReloadOptions, RequestPayload, RouterInitParams, Visit, VisitCallbacks, VisitHelperOptions, VisitOptions } from './types';
export declare class Router {
    protected syncRequestStream: RequestStream;
    protected asyncRequestStream: RequestStream;
    init({ initialPage, resolveComponent, swapComponent }: RouterInitParams): void;
    get<T extends RequestPayload = RequestPayload>(url: URL | string, data?: T, options?: VisitHelperOptions<T>): void;
    post<T extends RequestPayload = RequestPayload>(url: URL | string, data?: T, options?: VisitHelperOptions<T>): void;
    put<T extends RequestPayload = RequestPayload>(url: URL | string, data?: T, options?: VisitHelperOptions<T>): void;
    patch<T extends RequestPayload = RequestPayload>(url: URL | string, data?: T, options?: VisitHelperOptions<T>): void;
    delete<T extends RequestPayload = RequestPayload>(url: URL | string, options?: Omit<VisitOptions<T>, 'method'>): void;
    reload<T extends RequestPayload = RequestPayload>(options?: ReloadOptions<T>): void;
    remember(data: unknown, key?: string): void;
    restore(key?: string): unknown;
    on<TEventName extends GlobalEventNames>(type: TEventName, callback: (event: GlobalEvent<TEventName>) => GlobalEventResult<TEventName>): VoidFunction;
    cancel(): void;
    cancelAll(): void;
    poll(interval: number, requestOptions?: ReloadOptions, options?: PollOptions): {
        stop: VoidFunction;
        start: VoidFunction;
    };
    visit<T extends RequestPayload = RequestPayload>(href: string | URL, options?: VisitOptions<T>): void;
    getCached(href: string | URL, options?: VisitOptions): InFlightPrefetch | PrefetchedResponse | null;
    flush(href: string | URL, options?: VisitOptions): void;
    flushAll(): void;
    getPrefetching(href: string | URL, options?: VisitOptions): InFlightPrefetch | PrefetchedResponse | null;
    prefetch(href: string | URL, options: Partial<Visit<RequestPayload> & VisitCallbacks> | undefined, { cacheFor }: PrefetchOptions): void;
    clearHistory(): void;
    decryptHistory(): Promise<Page>;
    replace(params: ClientSideVisitOptions): void;
    push(params: ClientSideVisitOptions): void;
    protected clientVisit(params: ClientSideVisitOptions, { replace }?: {
        replace?: boolean;
    }): void;
    protected getPrefetchParams(href: string | URL, options: VisitOptions): ActiveVisit;
    protected getPendingVisit(href: string | URL, options: VisitOptions, pendingVisitOptions?: Partial<PendingVisitOptions>): PendingVisit;
    protected getVisitEvents(options: VisitOptions): VisitCallbacks;
    protected loadDeferredProps(): void;
}
