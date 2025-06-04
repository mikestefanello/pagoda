import { GlobalEvent, GlobalEventNames, GlobalEventResult, InternalEvent } from './types';
declare class EventHandler {
    protected internalListeners: {
        event: InternalEvent;
        listener: VoidFunction;
    }[];
    init(): void;
    onGlobalEvent<TEventName extends GlobalEventNames>(type: TEventName, callback: (event: GlobalEvent<TEventName>) => GlobalEventResult<TEventName>): VoidFunction;
    on(event: InternalEvent, callback: VoidFunction): VoidFunction;
    onMissingHistoryItem(): void;
    fireInternalEvent(event: InternalEvent): void;
    protected registerListener(type: string, listener: EventListener): VoidFunction;
    protected handlePopstateEvent(event: PopStateEvent): void;
}
export declare const eventHandler: EventHandler;
export {};
