import { FormDataConvertible } from './types';
export declare const isFormData: (value: any) => value is FormData;
export declare function objectToFormData(source: Record<string, FormDataConvertible>, form?: FormData, parentKey?: string | null): FormData;
