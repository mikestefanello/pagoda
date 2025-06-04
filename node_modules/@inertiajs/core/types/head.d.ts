export default function createHeadManager(isServer: boolean, titleCallback: (title: string) => string, onUpdate: (elements: string[]) => void): {
    forceUpdate: () => void;
    createProvider: () => {
        update: (elements: string[]) => void;
        disconnect: () => void;
    };
};
