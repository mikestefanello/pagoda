import { VisitOptions } from '@inertiajs/core';
export default function usePrefetch(options?: VisitOptions): {
    lastUpdatedAt: number | null;
    isPrefetching: boolean;
    isPrefetched: boolean;
    flush: () => void;
};
