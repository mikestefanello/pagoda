import { ScrollRegion } from './types';
export declare class Scroll {
    static save(): void;
    protected static regions(): NodeListOf<Element>;
    static reset(): void;
    static restore(scrollRegions: ScrollRegion[]): void;
    static restoreDocument(): void;
    static onScroll(event: Event): void;
    static onWindowScroll(): void;
}
