import { QueryClient as ReactQueryClient } from 'react-query';
import { create } from 'apisauce';

const THIRTY_SECONDS = 30 * 1000;

export const QueryClient = new ReactQueryClient({
    defaultOptions: {
        queries: {
            staleTime: THIRTY_SECONDS,
        },
    },
});

export const Apisauce = create({
    baseURL: 'http://127.0.0.1:8000/api',
    xsrfCookieName: '_csrf',
    xsrfHeaderName: 'X-CSRF-Token',
});

export * from './shows/queries';
export * from './shows/dtos';
