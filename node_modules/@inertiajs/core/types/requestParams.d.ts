import { AxiosRequestConfig } from 'axios';
import { Response } from './response';
import { ActiveVisit, InternalActiveVisit, Page, PreserveStateOption, VisitCallbacks } from './types';
export declare class RequestParams {
    protected callbacks: {
        name: keyof VisitCallbacks;
        args: any[];
    }[];
    protected params: InternalActiveVisit;
    constructor(params: InternalActiveVisit);
    static create(params: ActiveVisit): RequestParams;
    data(): import("./types").RequestPayload | null;
    queryParams(): import("./types").RequestPayload;
    isPartial(): boolean;
    onCancelToken(cb: VoidFunction): void;
    markAsFinished(): void;
    markAsCancelled({ cancelled, interrupted }: {
        cancelled?: boolean | undefined;
        interrupted?: boolean | undefined;
    }): void;
    wasCancelledAtAll(): boolean;
    onFinish(): void;
    onStart(): void;
    onPrefetching(): void;
    onPrefetchResponse(response: Response): void;
    all(): InternalActiveVisit;
    headers(): AxiosRequestConfig['headers'];
    setPreserveOptions(page: Page): void;
    runCallbacks(): void;
    merge(toMerge: Partial<ActiveVisit>): void;
    protected wrapCallback(params: ActiveVisit, name: keyof VisitCallbacks): (...args: any[]) => void;
    protected recordCallback(name: keyof VisitCallbacks, args: any[]): void;
    protected resolvePreserveOption(value: PreserveStateOption, page: Page): boolean;
}
