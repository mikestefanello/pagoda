import { FormDataConvertible, Method, RequestPayload, VisitOptions } from './types';
export declare function hrefToUrl(href: string | URL): URL;
export declare const transformUrlAndData: (href: string | URL, data: RequestPayload, method: Method, forceFormData: VisitOptions['forceFormData'], queryStringArrayFormat: VisitOptions['queryStringArrayFormat']) => [URL, RequestPayload];
export declare function mergeDataIntoQueryString(method: Method, href: URL | string, data: Record<string, FormDataConvertible>, qsArrayFormat?: 'indices' | 'brackets'): [string, Record<string, FormDataConvertible>];
export declare function urlWithoutHash(url: URL | Location): URL;
export declare const setHashIfSameUrl: (originUrl: URL | Location, destinationUrl: URL | Location) => void;
export declare const isSameUrlWithoutHash: (url1: URL | Location, url2: URL | Location) => boolean;
