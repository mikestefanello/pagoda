import { AxiosResponse } from 'axios';
import { RequestParams } from './requestParams';
import { ActiveVisit, ErrorBag, Errors, Page } from './types';
export declare class Response {
    protected requestParams: RequestParams;
    protected response: AxiosResponse;
    protected originatingPage: Page;
    constructor(requestParams: RequestParams, response: AxiosResponse, originatingPage: Page);
    static create(params: RequestParams, response: AxiosResponse, originatingPage: Page): Response;
    handlePrefetch(): Promise<void>;
    handle(): Promise<void>;
    process(): Promise<boolean | void>;
    mergeParams(params: ActiveVisit): void;
    protected handleNonInertiaResponse(): Promise<boolean | void>;
    protected isInertiaResponse(): boolean;
    protected hasStatus(status: number): boolean;
    protected getHeader(header: string): string;
    protected hasHeader(header: string): boolean;
    protected isLocationVisit(): boolean;
    /**
     * @link https://inertiajs.com/redirects#external-redirects
     */
    protected locationVisit(url: URL): boolean | void;
    protected setPage(): Promise<void>;
    protected getDataFromResponse(response: any): any;
    protected shouldSetPage(pageResponse: Page): boolean;
    protected pageUrl(pageResponse: Page): string;
    protected mergeProps(pageResponse: Page): void;
    protected setRememberedState(pageResponse: Page): Promise<void>;
    protected getScopedErrors(errors: Errors & ErrorBag): Errors;
}
