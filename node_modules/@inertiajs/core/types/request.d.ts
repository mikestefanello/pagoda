import { AxiosProgressEvent, AxiosRequestConfig } from 'axios';
import { RequestParams } from './requestParams';
import { Response } from './response';
import { ActiveVisit, Page } from './types';
export declare class Request {
    protected page: Page;
    protected response: Response;
    protected cancelToken: AbortController;
    protected requestParams: RequestParams;
    protected requestHasFinished: boolean;
    constructor(params: ActiveVisit, page: Page);
    static create(params: ActiveVisit, page: Page): Request;
    send(): Promise<void | undefined>;
    protected finish(): void;
    protected fireFinishEvents(): void;
    cancel({ cancelled, interrupted }: {
        cancelled?: boolean;
        interrupted?: boolean;
    }): void;
    protected onProgress(progress: AxiosProgressEvent): void;
    protected getHeaders(): AxiosRequestConfig['headers'];
}
