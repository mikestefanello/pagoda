export default function debounce<F extends (...params: any[]) => ReturnType<F>>(fn: F, delay: number): F;
