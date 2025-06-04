import { InertiaAppResponse, Page } from './types';
type AppCallback = (page: Page) => InertiaAppResponse;
type ServerOptions = {
    port?: number;
    cluster?: boolean;
};
type Port = number;
declare const _default: (render: AppCallback, options?: Port | ServerOptions) => void;
export default _default;
