import { FormDataConvertible, FormDataKeys, FormDataValues, Method, Progress, VisitOptions } from '@inertiajs/core';
type SetDataByObject<TForm> = (data: TForm) => void;
type SetDataByMethod<TForm> = (data: (previousData: TForm) => TForm) => void;
type SetDataByKeyValuePair<TForm extends Record<any, any>> = <K extends FormDataKeys<TForm>>(key: K, value: FormDataValues<TForm, K>) => void;
type FormDataType = Record<string, FormDataConvertible>;
type FormOptions = Omit<VisitOptions, 'data'>;
export interface InertiaFormProps<TForm extends FormDataType> {
    data: TForm;
    isDirty: boolean;
    errors: Partial<Record<FormDataKeys<TForm>, string>>;
    hasErrors: boolean;
    processing: boolean;
    progress: Progress | null;
    wasSuccessful: boolean;
    recentlySuccessful: boolean;
    setData: SetDataByObject<TForm> & SetDataByMethod<TForm> & SetDataByKeyValuePair<TForm>;
    transform: (callback: (data: TForm) => object) => void;
    setDefaults(): void;
    setDefaults(field: FormDataKeys<TForm>, value: FormDataConvertible): void;
    setDefaults(fields: Partial<TForm>): void;
    reset: (...fields: FormDataKeys<TForm>[]) => void;
    clearErrors: (...fields: FormDataKeys<TForm>[]) => void;
    setError(field: FormDataKeys<TForm>, value: string): void;
    setError(errors: Record<FormDataKeys<TForm>, string>): void;
    submit: (...args: [Method, string, FormOptions?] | [{
        url: string;
        method: Method;
    }, FormOptions?]) => void;
    get: (url: string, options?: FormOptions) => void;
    patch: (url: string, options?: FormOptions) => void;
    post: (url: string, options?: FormOptions) => void;
    put: (url: string, options?: FormOptions) => void;
    delete: (url: string, options?: FormOptions) => void;
    cancel: () => void;
}
export default function useForm<TForm extends FormDataType>(initialValues?: TForm): InertiaFormProps<TForm>;
export default function useForm<TForm extends FormDataType>(rememberKey: string, initialValues?: TForm): InertiaFormProps<TForm>;
export {};
