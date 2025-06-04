import { Page, PageProps, PageResolver } from '@inertiajs/core';
import { ComponentType, FunctionComponent, Key, ReactElement, ReactNode } from 'react';
import { renderToString } from 'react-dom/server';
type ReactInstance = ReactElement;
type ReactComponent = ReactNode;
type HeadManagerOnUpdate = (elements: string[]) => void;
type HeadManagerTitleCallback = (title: string) => string;
type AppType<SharedProps extends PageProps = PageProps> = FunctionComponent<{
    children?: (props: {
        Component: ComponentType;
        key: Key;
        props: Page<SharedProps>['props'];
    }) => ReactNode;
} & SetupOptions<unknown, SharedProps>['props']>;
export type SetupOptions<ElementType, SharedProps extends PageProps> = {
    el: ElementType;
    App: AppType;
    props: {
        initialPage: Page<SharedProps>;
        initialComponent: ReactComponent;
        resolveComponent: PageResolver;
        titleCallback?: HeadManagerTitleCallback;
        onHeadUpdate?: HeadManagerOnUpdate;
    };
};
type BaseInertiaAppOptions = {
    title?: HeadManagerTitleCallback;
    resolve: PageResolver;
};
type CreateInertiaAppSetupReturnType = ReactInstance | void;
type InertiaAppOptionsForCSR<SharedProps extends PageProps> = BaseInertiaAppOptions & {
    id?: string;
    page?: Page | string;
    render?: undefined;
    progress?: false | {
        delay?: number;
        color?: string;
        includeCSS?: boolean;
        showSpinner?: boolean;
    };
    setup(options: SetupOptions<HTMLElement, SharedProps>): CreateInertiaAppSetupReturnType;
};
type CreateInertiaAppSSRContent = {
    head: string[];
    body: string;
};
type InertiaAppOptionsForSSR<SharedProps extends PageProps> = BaseInertiaAppOptions & {
    id?: undefined;
    page: Page | string;
    render: typeof renderToString;
    progress?: undefined;
    setup(options: SetupOptions<null, SharedProps>): ReactInstance;
};
export default function createInertiaApp<SharedProps extends PageProps = PageProps>(options: InertiaAppOptionsForCSR<SharedProps>): Promise<CreateInertiaAppSetupReturnType>;
export default function createInertiaApp<SharedProps extends PageProps = PageProps>(options: InertiaAppOptionsForSSR<SharedProps>): Promise<CreateInertiaAppSSRContent>;
export {};
