import { FormDataConvertible, LinkPrefetchOption, Method, PendingVisit, PreserveStateOption, Progress } from '@inertiajs/core';
interface BaseInertiaLinkProps {
    as?: string;
    data?: Record<string, FormDataConvertible>;
    href: string | {
        url: string;
        method: Method;
    };
    method?: Method;
    headers?: Record<string, string>;
    onClick?: (event: React.MouseEvent<Element>) => void;
    preserveScroll?: PreserveStateOption;
    preserveState?: PreserveStateOption;
    replace?: boolean;
    only?: string[];
    except?: string[];
    onCancelToken?: (cancelToken: import('axios').CancelTokenSource) => void;
    onBefore?: () => void;
    onStart?: (event: PendingVisit) => void;
    onProgress?: (progress: Progress) => void;
    onFinish?: (event: PendingVisit) => void;
    onCancel?: () => void;
    onSuccess?: () => void;
    onError?: () => void;
    queryStringArrayFormat?: 'indices' | 'brackets';
    async?: boolean;
    cacheFor?: number | string;
    prefetch?: boolean | LinkPrefetchOption | LinkPrefetchOption[];
}
export type InertiaLinkProps = BaseInertiaLinkProps & Omit<React.HTMLAttributes<HTMLElement>, keyof BaseInertiaLinkProps> & Omit<React.AllHTMLAttributes<HTMLElement>, keyof BaseInertiaLinkProps>;
declare const Link: import("react").ForwardRefExoticComponent<BaseInertiaLinkProps & Omit<import("react").HTMLAttributes<HTMLElement>, keyof BaseInertiaLinkProps> & Omit<import("react").AllHTMLAttributes<HTMLElement>, keyof BaseInertiaLinkProps> & import("react").RefAttributes<unknown>>;
export default Link;
