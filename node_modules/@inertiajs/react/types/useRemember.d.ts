import { Dispatch, SetStateAction } from 'react';
export default function useRemember<State>(initialState: State, key?: string): [State, Dispatch<SetStateAction<State>>];
