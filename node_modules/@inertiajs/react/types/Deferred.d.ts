import { ReactElement } from 'react';
interface DeferredProps {
    children: ReactElement | number | string;
    fallback: ReactElement | number | string;
    data: string | string[];
}
declare const Deferred: {
    ({ children, data, fallback }: DeferredProps): string | number | ReactElement<unknown, string | import("react").JSXElementConstructor<any>>;
    displayName: string;
};
export default Deferred;
