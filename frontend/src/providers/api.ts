import { QueryClient as ReactQueryClient } from 'react-query';
import { ApiResponse, create } from 'apisauce';

const THIRTY_SECONDS = 30 * 1000;

export const QueryClient = new ReactQueryClient({
    defaultOptions: {
        queries: {
            staleTime: THIRTY_SECONDS,
        },
    },
});

export const ApiClient = create({
    baseURL: import.meta.env.VITE_BACKEND_URL + '/api',
    xsrfCookieName: '_csrf',
    xsrfHeaderName: 'X-CSRF-Token',
});

ApiClient.addResponseTransform((response: ApiResponse<any>) => {
    if (!response.ok) throw response;
});
