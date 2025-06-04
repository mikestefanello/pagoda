declare function App({ children, initialPage, initialComponent, resolveComponent, titleCallback, onHeadUpdate, }: {
    children: any;
    initialPage: any;
    initialComponent: any;
    resolveComponent: any;
    titleCallback: any;
    onHeadUpdate: any;
}): import("react").FunctionComponentElement<import("react").ProviderProps<any>>;
declare namespace App {
    var displayName: string;
}
export default App;
