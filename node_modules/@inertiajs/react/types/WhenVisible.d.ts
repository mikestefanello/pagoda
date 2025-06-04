import { ReloadOptions } from '@inertiajs/core';
import { ReactElement } from 'react';
interface WhenVisibleProps {
    children: ReactElement | number | string;
    fallback: ReactElement | number | string;
    data?: string | string[];
    params?: ReloadOptions;
    buffer?: number;
    as?: string;
    always?: boolean;
}
declare const WhenVisible: {
    ({ children, data, params, buffer, as, always, fallback }: WhenVisibleProps): string | number | ReactElement<unknown, string | import("react").JSXElementConstructor<any>>;
    displayName: string;
};
export default WhenVisible;
